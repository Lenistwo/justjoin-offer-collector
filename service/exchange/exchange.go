package exchange

import (
	"encoding/json"
	"github.com/lenistwo/model"
	"github.com/lenistwo/service/rest"
	"strings"
)

const (
	ApiURL  = "http://api.nbp.pl/api/exchangerates/tables/a?format=json"
	PLNCode = "pln"
)

var exchangeRates model.ExchangeRates

func init() {
	retrieveExchangeRates()
}

func retrieveExchangeRates() {
	response := rest.SendGetRequest(ApiURL)
	var rates []model.ExchangeRates
	checkError(json.Unmarshal(response, &rates))
	if len(rates) < 1 {
		panic("Invalid Exchange Rate Table")
	}
	exchangeRates = rates[0]
}

func Calculate(from, to string, amount int) int {
	return int(float64(amount) * findRate(from) * findRate(to))
}

func findRate(code string) float64 {
	if strings.ToLower(code) == PLNCode {
		return 1
	}

	for _, r := range exchangeRates.Rates {
		if strings.ToLower(r.Code) == strings.ToLower(code) {
			return r.Mid
		}
	}
	return 0
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
