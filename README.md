cmc-proapi [![License](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/nikchis/cmc-proapi/master/LICENSE) [![Build Status](https://travis-ci.org/nikchis/cmc-proapi.svg?branch=master)](https://travis-ci.org/nikchis/cmc-proapi) [![Go Report Card](https://goreportcard.com/badge/github.com/nikchis/cmc-proapi?)](https://goreportcard.com/report/github.com/nikchis/cmc-proapi) [![GoDoc](https://godoc.org/github.com/nikchis/cmc-proapi?status.svg)](https://godoc.org/github.com/nikchis/cmc-proapi)
==========

cmc-proapi is an implementation of the CoinMarketCap API in Golang.

## Import
	import "github.com/nikchis/cmc-proapi"
	
## Usage

~~~ go
package main

import (
	"fmt"
	"github.com/nikchis/cmc-proapi"
)

const (
	ApiKey = "YOUR_API_KEY"
)

func main() {
	// cmcproapi client
	cmc := cmcproapi.New(ApiKey)

	// get info about bitcoin
	info, err := cmc.GetCurrencyInfoBySymbol("BTC")
	if err != nil {
		fmt.Printf("%+v\n", info)
	}
}
~~~
