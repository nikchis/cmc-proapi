// Implementation of CoinMarketCap API
// Copyright (c) 2019 Nikita Chisnikov <chisnikov@gmail.com>
// Distributed under the MIT/X11 software license

package cmcproapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// do returns response from HTTP request or returns error if time exceeded.
func (c *Client) do(request *http.Request) (*http.Response, error) {
	type reqres struct {
		response *http.Response
		err      error
	}
	timer := time.NewTimer(c.httpTimeout)
	done := make(chan reqres, 1)
	go func() {
		r, err := c.httpClient.Do(request)
		done <- reqres{r, err}
	}()
	select {
	case r := <-done:
		return r.response, r.err
	case <-timer.C:
		return nil, errors.New(ltMsgRequestExceeded)
	}
}

// call prepares and process HTTP request to endpoint.
func (c *Client) call(endpoint string, query url.Values) (result []byte, err error) {
	var req *http.Request
	var resp *http.Response
	rawurl := fmt.Sprintf("%s/%s/%s", c.apiDomain, c.apiVersion, endpoint)
	if req, err = http.NewRequest("GET", rawurl, nil); err != nil {
		return
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Add(ltCmcProApiKeyX, c.apiKey)
	req.URL.RawQuery = query.Encode()
	if resp, err = c.do(req); err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusUnauthorized &&
		resp.StatusCode != http.StatusPaymentRequired && resp.StatusCode != http.StatusForbidden &&
		resp.StatusCode != http.StatusTooManyRequests {
		err = errors.New(resp.Status)
		return
	}
	defer resp.Body.Close()
	if result, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	return
}

// handleRequest returns raw JSON data after succesfull request to API and handling response status.
func (c *Client) handleRequest(args ...interface{}) (json.RawMessage, error) {
	var query *url.Values
	var endpoint string
	var rawresponse []byte
	var response Response
	var err error
	if args == nil {
		err = errors.New(ltMsgEmptyArgs)
		return nil, err
	}
	for _, arg := range args {
		switch val := arg.(type) {
		case string:
			if endpoint == "" {
				endpoint = val
			}
		case *url.Values:
			query = val
		default:
			err = errors.New(ltMsgUnsupArgType)
			return nil, err
		}
	}
	if query == nil {
		query = &url.Values{}
	}
	if rawresponse, err = c.call(endpoint, *query); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(rawresponse, &response); err != nil {
		return nil, err
	}
	if err = response.handleStatus(); err != nil {
		return nil, err
	}
	return response.Data, nil
}
