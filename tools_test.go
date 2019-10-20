// Implementation of CoinMarketCap API
// Copyright (c) 2019 Nikita Chisnikov <chisnikov@gmail.com>
// Distributed under the MIT/X11 software license

package cmcproapi

import (
	"testing"
)

func TestGetPriceConversion(t *testing.T) {
	cmc, err := NewTest()
	_, err = cmc.GetPriceConversionBySymbol("2", "LTC", "USD")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
