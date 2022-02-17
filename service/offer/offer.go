package offer

import (
	"encoding/json"
	"github.com/lenistwo/model"
	"github.com/lenistwo/service/rest"
)

const ApiURL = "https://justjoin.it/api/offers"

func Retrieve() []model.Offer {
	response := rest.SendGetRequest(ApiURL)
	var offers []model.Offer
	checkError(json.Unmarshal(response, &offers))
	return offers
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
