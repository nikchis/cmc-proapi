// Implementation of CoinMarketCap API
// Copyright (c) 2019 Nikita Chisnikov <chisnikov@gmail.com>
// Distributed under the MIT/X11 software license

package cmcproapi

import (
	"encoding/json"
	"net/url"
)

type PriceConversion struct {
	Id          int                 `json:"id"`
	Name        string              `json:"name"`
	Symbol      string              `json:"symbol"`
	Amount      float64             `json:"amount"`
	LastUpdated jsonTime            `json:"last_updated"`
	Quote       *ConversionQuoteMap `json:"quote"`
}

type ConversionQuoteMap map[string]ConversionQuote

type ConversionQuote struct {
	Price       float64  `json:"price"`
	LastUpdated jsonTime `json:"last_updated"`
}

// GetPriceConversionById converts an amount of one cryptocurrency or fiat currency into
// one or more different currencies utilizing the latest market rate for each currency.
func (c *Client) GetPriceConversionById(
	amount, id, convert_id string) (result PriceConversion, err error) {
	var raw json.RawMessage
	q := url.Values{}
	q.Set("convert_id", convert_id)
	q.Add("id", id)
	q.Add("amount", amount)
	if raw, err = c.handleRequest(ltUriToolsPriceConversion, &q); err != nil {
		return
	}
	err = json.Unmarshal(raw, &result)
	return
}

// GetPriceConversionBySymbol acts identically to GetPriceConversionById, except that it
// uses convert and symbol instead of convert_id and id as query parameters.
func (c *Client) GetPriceConversionBySymbol(
	amount, symbol, convert string) (result PriceConversion, err error) {
	var raw json.RawMessage
	q := url.Values{}
	q.Set("convert", convert)
	q.Add("symbol", symbol)
	q.Add("amount", amount)
	if raw, err = c.handleRequest(ltUriToolsPriceConversion, &q); err != nil {
		return
	}
	err = json.Unmarshal(raw, &result)
	return
}
