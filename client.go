// Implementation of CoinMarketCap API
// Copyright (c) 2019 Nikita Chisnikov <chisnikov@gmail.com>
// Distributed under the MIT/X11 software license

package cmcproapi

import (
	"errors"
	"net/http"
	"os"
	"time"
)

const (
	ApiDomain         = "https://pro-api.coinmarketcap.com"
	ApiVersion        = "v1"
	ApiRequestTimeout = 30
)

type Client struct {
	apiKey      string
	apiDomain   string
	apiVersion  string
	httpClient  *http.Client
	httpTimeout time.Duration
}

// New returns an instantiated Client struct.
func New(apiKey string) (c *Client, err error) {
	if apiKey == "" {
		apiKey = os.Getenv(ltCmcProApiKey)
	}
	c = &Client{
		apiKey:      apiKey,
		apiDomain:   ApiDomain,
		apiVersion:  ApiVersion,
		httpClient:  &http.Client{},
		httpTimeout: ApiRequestTimeout * time.Second,
	}
	return
}

// NewCustom returns an instantiated Client struct with custom properties.
//
// e.g. NewCustom(apiKey, apiDomain, apiVersion, &http.Client{}, time.Minute)
func NewCustom(args ...interface{}) (c *Client, err error) {
	if args == nil {
		err = errors.New(ltMsgEmptyArgs)
		return
	}
	c = &Client{}
	for _, arg := range args {
		switch val := arg.(type) {
		case string:
			if c.apiKey == "" {
				c.apiKey = val
			} else if c.apiDomain == "" {
				c.apiDomain = val
			} else if c.apiVersion == "" {
				c.apiVersion = val
			}
		case *http.Client:
			c.httpClient = val
		case time.Duration:
			c.httpTimeout = val
		default:
			err = errors.New(ltMsgUnsupArgType)
			return
		}
	}
	return
}
