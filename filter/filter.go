package filter

import (
	"github.com/lenistwo/model"
	"github.com/lenistwo/service/exchange"
	"strings"
)

type ExperienceLevel string
type EmploymentType string

type SearchCriteria struct {
	Keywords        []string
	EmploymentType  EmploymentType
	Skills          []string
	ExperienceLevel ExperienceLevel
	Salary          int
	Currency        string
	Remote          bool
}

const (
	Senior          ExperienceLevel = "senior"
	Mid             ExperienceLevel = "mid"
	Junior          ExperienceLevel = "junior"
	AllLevels       ExperienceLevel = "all"
	AllTypes        EmploymentType  = "all"
	MandateContract EmploymentType  = "mandate_contract"
	B2B             EmploymentType  = "b2b"
	Permanent       EmploymentType  = "permanent"
)

var filters []func(offers []model.Offer, criteria SearchCriteria) []model.Offer

func init() {
	filters = []func(offers []model.Offer, criteria SearchCriteria) []model.Offer{remote, experienceLevel, employmentType, salary, skills}
}

func FindOffers(o []model.Offer, c SearchCriteria) []model.Offer {
	var filtered []model.Offer
	for _, offer := range o {
		for _, k := range c.Keywords {
			if strings.Contains(strings.ToLower(offer.Title), strings.ToLower(k)) {
				filtered = append(filtered, offer)
				break
			}
		}
	}

	for _, f := range filters {
		filtered = f(filtered, c)
	}

	return filtered
}

func remote(offers []model.Offer, criteria SearchCriteria) []model.Offer {
	var filtered []model.Offer
	for _, offer := range offers {
		if offer.Remote != criteria.Remote {
			continue
		}
		filtered = append(filtered, offer)
	}

	return filtered
}

func experienceLevel(offers []model.Offer, criteria SearchCriteria) []model.Offer {
	if criteria.ExperienceLevel == AllLevels {
		return offers
	}

	var filtered []model.Offer
	for _, offer := range offers {
		if string(criteria.ExperienceLevel) != offer.ExperienceLevel {
			continue
		}

		filtered = append(filtered, offer)
	}
	return filtered
}

func employmentType(offers []model.Offer, criteria SearchCriteria) []model.Offer {
	if criteria.EmploymentType == AllTypes {
		return offers
	}

	var filtered []model.Offer
	for _, offer := range offers {
		found := false
		for _, e := range offer.EmploymentTypes {
			if e.Type == string(criteria.EmploymentType) {
				found = true
				break
			}
		}

		if !found {
			continue
		}

		filtered = append(filtered, offer)
	}
	return filtered
}

func salary(offers []model.Offer, criteria SearchCriteria) []model.Offer {
	var filtered []model.Offer

	for _, offer := range offers {
		for _, employment := range offer.EmploymentTypes {
			if criteria.EmploymentType != AllTypes && employment.Type != string(criteria.EmploymentType) {
				continue
			}

			s := employment.Salary
			if s.To == 0 && s.From == 0 {
				continue
			}

			if strings.ToLower(criteria.Currency) != strings.ToLower(employment.Salary.Currency) {
				s.To = exchange.Calculate(s.Currency, criteria.Currency, s.To)
			}

			if criteria.Salary > s.To {
				continue
			}

			if containsId(filtered, offer.Id) {
				break
			}

			filtered = append(filtered, offer)
		}
	}

	return filtered
}

func skills(offers []model.Offer, criteria SearchCriteria) []model.Offer {
	var filtered []model.Offer
	for _, offer := range offers {
		found := false
		for _, s := range offer.Skills {
			if contains(criteria.Skills, s.Name) {
				found = true
				break
			}
		}

		if !found {
			continue
		}

		filtered = append(filtered, offer)
	}
	return filtered
}

func contains(arr []string, e string) bool {
	for _, i := range arr {
		if strings.ToLower(i) == strings.ToLower(e) {
			return true
		}
	}
	return false
}

func containsId(arr []model.Offer, id string) bool {
	for _, i := range arr {
		if strings.ToLower(i.Id) == strings.ToLower(id) {
			return true
		}
	}
	return false
}
