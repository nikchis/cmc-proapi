// Implementation of CoinMarketCap API
// Copyright (c) 2019 Nikita Chisnikov <chisnikov@gmail.com>
// Distributed under the MIT/X11 software license

package cmcproapi

const (
	ltCmcProApiKey  = "CMC_PRO_API_KEY"
	ltCmcProApiKeyX = "X-CMC_PRO_API_KEY"

	ltMsgEmptyArgs       = "empty arguments"
	ltMsgRequestExceeded = "request exceeded the timelimit"
	ltMsgUnsupArgType    = "unsupported argument type"

	ltUriCurrencyMap            = "cryptocurrency/map"
	ltUriCurrencyInfo           = "cryptocurrency/info"
	ltUriCurrencyListingsLatest = "cryptocurrency/listings/latest"
	ltUriCurrencyQuotesLatest   = "cryptocurrency/quotes/latest"
	ltUriGlobalQuotesLatest     = "global-metrics/quotes/latest"
	ltUriToolsPriceConversion   = "tools/price-conversion"
)
