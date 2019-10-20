// Implementation of CoinMarketCap API
// Copyright (c) 2019 Nikita Chisnikov <chisnikov@gmail.com>
// Distributed under the MIT/X11 software license

package cmcproapi

import (
	"encoding/json"
	"net/url"
)

type GlobalMetrics struct {
	BtcDominance           float64         `json:"btc_dominance"`
	EthDominance           float64         `json:"eth_dominance"`
	ActiveCryptocurrencies int             `json:"active_cryptocurrencies"`
	TotalCryptocurrencies  int             `json:"total_cryptocurrencies"`
	ActiveMarketPairs      int             `json:"active_market_pairs"`
	ActiveExchanges        int             `json:"active_exchanges"`
	TotalExchanges         int             `json:"total_exchanges"`
	LastUpdated            jsonTime        `json:"last_updated"`
	Quote                  *GlobalQuoteMap `json:"quote"`
}

type GlobalQuoteMap map[string]GlobalQuote

type GlobalQuote struct {
	TotalMarketCap           float64  `json:"total_market_cap"`
	TotalVolume24h           float64  `json:"total_volume_24h"`
	TotalVolume24hReported   float64  `json:"total_volume_24h_reported"`
	AltcoinVolume24h         float64  `json:"altcoin_volume_24h"`
	AltcoinVolume24hReported float64  `json:"altcoin_volume_24h_reported"`
	AltcoinMarketCap         float64  `json:"altcoin_market_cap"`
	LastUpdated              jsonTime `json:"last_updated"`
}

// GetGlobalQuotesLatestById returns the latest global cryptocurrency market metrics.
//
// Use the "convert_id" to return market values in multiple fiat and cryptocurrency
// conversions in the same call.
func (c *Client) GetGlobalQuotesLatestById(convert_id string) (result GlobalMetrics, err error) {
	var raw json.RawMessage
	q := url.Values{}
	q.Set("convert_id", convert_id)
	if raw, err = c.handleRequest(ltUriGlobalQuotesLatest, &q); err != nil {
		return
	}
	err = json.Unmarshal(raw, &result)
	return
}

// GetGlobalQuotesLatestBySymbol acts identically to GetGlobalQuotesLatestById, except that it
// uses convert instead of convert_id as query parameter.
func (c *Client) GetGlobalQuotesLatestBySymbol(convert string) (result GlobalMetrics, err error) {
	var raw json.RawMessage
	q := url.Values{}
	q.Set("convert", convert)
	if raw, err = c.handleRequest(ltUriGlobalQuotesLatest, &q); err != nil {
		return
	}
	err = json.Unmarshal(raw, &result)
	return
}
