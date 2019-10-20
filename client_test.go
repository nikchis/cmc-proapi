// Implementation of CoinMarketCap API
// Copyright (c) 2019 Nikita Chisnikov <chisnikov@gmail.com>
// Distributed under the MIT/X11 software license

package cmcproapi

import (
	"net/http"
	"testing"
	"time"
)

const (
	TestApiKey     = "42d49444-36c2-47f2-97e5-6b1322657e3d"
	TestApiDomain  = "https://sandbox-api.coinmarketcap.com"
	TestApiVersion = "v1"
)

// NewTest returns an instantiated Client struct for testing purposes.
func NewTest() (c *Client, err error) {
	c = &Client{
		apiKey:      TestApiKey,
		apiDomain:   TestApiDomain,
		apiVersion:  TestApiVersion,
		httpClient:  &http.Client{},
		httpTimeout: ApiRequestTimeout * time.Second,
	}
	return
}

func TestNewCustom(t *testing.T) {
	_, err := NewCustom(
		TestApiKey,
		TestApiDomain,
		TestApiVersion,
		&http.Client{},
		ApiRequestTimeout*time.Second,
	)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
