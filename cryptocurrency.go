// Implementation of CoinMarketCap API
// Copyright (c) 2019 Nikita Chisnikov <chisnikov@gmail.com>
// Distributed under the MIT/X11 software license

package cmcproapi

import (
	"encoding/json"
	"net/url"
)

type ListingStatus string

const (
	ListingActive    ListingStatus = "active"
	ListingInactive  ListingStatus = "inactive"
	ListingUntracked ListingStatus = "untracked"
)

type CurrencyMap struct {
	Id                  int               `json:"id"`
	Name                string            `json:"name"`
	Symbol              string            `json:"symbol"`
	Slug                string            `json:"slug"`
	IsActive            int               `json:"is_active"`
	FirstHistoricalData jsonTime          `json:"first_historical_data"`
	LastHistoricalData  jsonTime          `json:"last_historical_data"`
	Platform            *CurrencyPlatform `json:"platform"`
}

type CurrencyPlatform struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	Slug         string `json:"slug"`
	TokenAddress string `json:"token_address"`
}

type CurrencyUrls struct {
	Website      []string `json:"website"`
	TechnicalDoc []string `json:"technical_doc"`
	Explorer     []string `json:"explorer"`
	SourceCode   []string `json:"source_code"`
	MessageBoard []string `json:"message_board"`
	Chat         []string `json:"chat"`
	Announcement []string `json:"announcement"`
	Reddit       []string `json:"reddit"`
	Twitter      []string `json:"twitter"`
}

type CurrencyInfoMap map[string]CurrencyInfo

type CurrencyInfo struct {
	Id          int               `json:"id"`
	Name        string            `json:"name"`
	Symbol      string            `json:"symbol"`
	Category    string            `json:"category"`
	Slug        string            `json:"slug"`
	Logo        string            `json:"logo"`
	Description string            `json:"description"`
	DateAdded   jsonTime          `json:"date_added"`
	Notice      string            `json:"notice"`
	Tags        []string          `json:"tags"`
	Platform    *CurrencyPlatform `json:"platform"`
	Urls        *CurrencyUrls     `json:"urls"`
}

type CurrencyListing struct {
	Id                     int               `json:"id"`
	Name                   string            `json:"name"`
	Symbol                 string            `json:"symbol"`
	Slug                   string            `json:"slug"`
	CmcRank                int               `json:"cmc_rank"`
	NumMarketPairs         int               `json:"num_market_pairs,omitempty"`
	CirculatingSupply      float64           `json:"circulating_supply"`
	TotalSupply            float64           `json:"total_supply"`
	MarketCapByTotalSupply float64           `json:"market_cap_by_total_supply,omitempty"`
	MaxSupply              float64           `json:"max_supply"`
	LastUpdated            jsonTime          `json:"last_updated"`
	DateAdded              jsonTime          `json:"date_added"`
	Tags                   []string          `json:"tags"`
	Platform               *CurrencyPlatform `json:"platform"`
	Quote                  *CurrencyQuoteMap `json:"quote"`
}

type CurrencyQuoteMap map[string]CurrencyQuote

type CurrencyQuote struct {
	Price             float64  `json:"price"`
	Volume24h         float64  `json:"volume_24h"`
	Volume24hReported float64  `json:"volume_24h_reported,omitempty"`
	Volume7d          float64  `json:"volume_7d,omitempty"`
	Volume7dReported  float64  `json:"volume_7d_reported,omitempty"`
	Volume30d         float64  `json:"volume_30d,omitempty"`
	Volume30dReported float64  `json:"volume_30d_reported,omitempty"`
	MarketCap         float64  `json:"market_cap"`
	PercentChange1h   float64  `json:"percent_change_1h"`
	PercentChange24h  float64  `json:"percent_change_24h"`
	PercentChange7d   float64  `json:"percent_change_7d"`
	LastUpdated       jsonTime `json:"last_updated"`
}

// GetCurrencyMap returns a mapping of cryptocurrencies to unique CoinMarketCap ids.
//
// Per best practices it recommends to utilizing ID instead of cryptocurrency symbols
// to securely identify cryptocurrencies with other endpoints and in application logic.
func (c *Client) GetCurrencyMap(
	lstatus ListingStatus, start, limit, symbol string) (result []CurrencyMap, err error) {
	var raw json.RawMessage
	q := url.Values{}
	q.Set("listing_status", string(lstatus))
	q.Add("start", start)
	q.Add("limit", limit)
	q.Add("sort", "cmc_rank")
	if symbol != "" {
		q.Add("symbol", symbol)
	}
	if raw, err = c.handleRequest(ltUriCurrencyMap, &q); err != nil {
		return
	}
	err = json.Unmarshal(raw, &result)
	return
}

// GetCurrencyMapAllActive acts identically to GetCurrencyMap, except that it
// uses predefined values for query parameters.
func (c *Client) GetCurrencyMapAllActive() (result []CurrencyMap, err error) {
	var raw json.RawMessage
	q := url.Values{}
	q.Set("listing_status", string(ListingActive))
	q.Add("limit", "5000")
	if raw, err = c.handleRequest(ltUriCurrencyMap, &q); err != nil {
		return
	}
	err = json.Unmarshal(raw, &result)
	return
}

// GetCurrencyInfoById returns all static metadata available for one or more cryptocurrencies.
//
// This information includes details like logo, description, official website URL, social links,
// and links to a cryptocurrency's technical documentation.
func (c *Client) GetCurrencyInfoById(id string) (result CurrencyInfoMap, err error) {
	var raw json.RawMessage
	q := url.Values{}
	q.Set("id", id)
	if raw, err = c.handleRequest(ltUriCurrencyInfo, &q); err != nil {
		return
	}
	err = json.Unmarshal(raw, &result)
	return
}

// GetCurrencyInfoBySymbol acts identically to GetCurrencyInfoById, except that it
// uses symbol instead of id as query parameter.
func (c *Client) GetCurrencyInfoBySymbol(symbol string) (result CurrencyInfoMap, err error) {
	var raw json.RawMessage
	q := url.Values{}
	q.Set("symbol", symbol)
	if raw, err = c.handleRequest(ltUriCurrencyInfo, &q); err != nil {
		return
	}
	err = json.Unmarshal(raw, &result)
	return
}

// GetCurrencyListingsLatestById returns a paginated list of all active cryptocurrencies
// with latest market data. The default "market_cap" sort returns cryptocurrency
// in order of CoinMarketCap's market cap rank.
func (c *Client) GetCurrencyListingsLatestById(
	start, limit, convert_id string) (result []CurrencyListing, err error) {
	var raw json.RawMessage
	q := url.Values{}
	q.Set("start", start)
	q.Add("limit", limit)
	q.Add("convert_id", convert_id)
	if raw, err = c.handleRequest(ltUriCurrencyListingsLatest, &q); err != nil {
		return
	}
	err = json.Unmarshal(raw, &result)
	return
}

// GetCurrencyListingsLatestBySymbol acts identically to GetCurrencyListingsLatestById, except
// that it uses symbols for convert parameter instead of ids.
func (c *Client) GetCurrencyListingsLatestBySymbol(
	start, limit, convert string) (result []CurrencyListing, err error) {
	var raw json.RawMessage
	q := url.Values{}
	q.Set("start", start)
	q.Add("limit", limit)
	q.Add("convert", convert)
	if raw, err = c.handleRequest(ltUriCurrencyListingsLatest, &q); err != nil {
		return
	}
	err = json.Unmarshal(raw, &result)
	return
}

// GetCurrencyListingsLatestAll acts identically to GetCurrencyListingsLatestById, except
// that it uses predefined values for query parameters.
func (c *Client) GetCurrencyListingsLatestAll() (result []CurrencyListing, err error) {
	var raw json.RawMessage
	q := url.Values{}
	q.Set("convert", "BTC,USD")
	q.Add("start", "1")
	q.Add("limit", "5000")
	if raw, err = c.handleRequest(ltUriCurrencyListingsLatest, &q); err != nil {
		return
	}
	err = json.Unmarshal(raw, &result)
	return
}

// GetCurrencyQuotesLatestById returns the latest market quote for 1 or more cryptocurrencies.
func (c *Client) GetCurrencyQuotesLatestById(
	id, convert_id string) (result CurrencyQuoteMap, err error) {
	var raw json.RawMessage
	q := url.Values{}
	q.Set("id", id)
	q.Add("convert_id", convert_id)
	if raw, err = c.handleRequest(ltUriCurrencyQuotesLatest, &q); err != nil {
		return
	}
	err = json.Unmarshal(raw, &result)
	return
}

// GetCurrencyQuotesLatestBySymbol acts identically to GetCurrencyQuotesLatestById, except
// that it uses symbols instead of ids.
func (c *Client) GetCurrencyQuotesLatestBySymbol(
	symbol, convert string) (result CurrencyQuoteMap, err error) {
	var raw json.RawMessage
	q := url.Values{}
	q.Set("symbol", symbol)
	q.Add("convert", convert)
	if raw, err = c.handleRequest(ltUriCurrencyQuotesLatest, &q); err != nil {
		return
	}
	err = json.Unmarshal(raw, &result)
	return
}
