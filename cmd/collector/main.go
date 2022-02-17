package main

import (
	"github.com/lenistwo/filter"
	"github.com/lenistwo/service/csv"
	"github.com/lenistwo/service/offer"
	"net/http"
)

var (
	client http.Client
)

func main() {
	offers := offer.Retrieve()
	criteria := filter.SearchCriteria{
		Keywords:        []string{"Java"},
		EmploymentType:  filter.AllTypes,
		Skills:          []string{"java"},
		ExperienceLevel: filter.Mid,
		Salary:          10000,
		Currency:        "pln",
		Remote:          true,
	}

	csv.Write(filter.FindOffers(offers, criteria))
}
