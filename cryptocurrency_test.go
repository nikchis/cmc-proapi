// Implementation of CoinMarketCap API
// Copyright (c) 2019 Nikita Chisnikov <chisnikov@gmail.com>
// Distributed under the MIT/X11 software license

package cmcproapi

import (
	"testing"
)

func TestGetCurrencyMap(t *testing.T) {
	cmc, err := NewTest()
	_, err = cmc.GetCurrencyMap(ListingActive, "1", "50", "")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestGetCurrencyInfo(t *testing.T) {
	cmc, err := NewTest()
	_, err = cmc.GetCurrencyInfoBySymbol("BTC")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestGetCurrencyListingsLatest(t *testing.T) {
	cmc, err := NewTest()
	_, err = cmc.GetCurrencyListingsLatestById("1", "600", "2781")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestGetCurrencyListingsLatestAll(t *testing.T) {
	cmc, err := NewTest()
	_, err = cmc.GetCurrencyListingsLatestAll()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestGetCurrencyQuotesLatest(t *testing.T) {
	cmc, err := NewTest()
	_, err = cmc.GetCurrencyQuotesLatestBySymbol("LTC", "USD,BTC")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
