// Implementation of CoinMarketCap API
// Copyright (c) 2019 Nikita Chisnikov <chisnikov@gmail.com>
// Distributed under the MIT/X11 software license

package cmcproapi

import (
	"encoding/json"
	"errors"
	"time"
)

type jsonTime time.Time

type Response struct {
	Data   json.RawMessage `json:"data"`
	Status ResponseStatus  `json:"status"`
}

type ResponseStatus struct {
	Timestamp    jsonTime `json:"timestamp"`
	ErrorCode    int      `json:"error_code"`
	ErrorMessage string   `json:"error_message"`
	Elapsed      int      `json:"elapsed"`
	CreditCount  int      `json:"credit_count"`
}

func (jt *jsonTime) UnmarshalJSON(data []byte) error {
	var str string
	var t time.Time
	var err error
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	if str == "null" || str == `""` || str == "" {
		*(*time.Time)(jt) = time.Time{}
		return nil
	}
	if t, err = time.Parse(time.RFC3339, str); err != nil {
		return err
	}
	*(*time.Time)(jt) = t
	return nil
}

func (jt *jsonTime) MarshalJSON() ([]byte, error) {
	return json.Marshal((*time.Time)(jt).Format(time.RFC3339))
}

// handleStatus checks the status of response.
func (resp *Response) handleStatus() (err error) {
	if resp.Status.ErrorCode != 0 {
		err = errors.New(resp.Status.ErrorMessage)
	}
	return
}
