package util

// Constants for all supported currencies
const (
	USD  = "USD"
	EUR  = "EUR"
	THB  = "THB"
	COIN = "COIN"
)

// SupportedCurrencies is a list of all supported currencies
var SupportedCurrencies = []string{USD, EUR, THB, COIN}

// IsCurrencySupported returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	for _, c := range SupportedCurrencies {
		if c == currency {
			return true
		}
	}
	return false
}
