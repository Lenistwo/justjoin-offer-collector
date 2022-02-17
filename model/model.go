package model

import "time"

type Offer struct {
	Title           string           `json:"title"`
	Street          string           `json:"street"`
	City            string           `json:"city"`
	CountryCode     string           `json:"country_code"`
	AddressText     string           `json:"address_text"`
	MarkerIcon      string           `json:"marker_icon"`
	WorkplaceType   string           `json:"workplace_type"`
	CompanyName     string           `json:"company_name"`
	CompanyUrl      string           `json:"company_url"`
	CompanySize     string           `json:"company_size"`
	ExperienceLevel string           `json:"experience_level"`
	Latitude        string           `json:"latitude"`
	Longitude       string           `json:"longitude"`
	PublishedAt     time.Time        `json:"published_at"`
	RemoteInterview bool             `json:"remote_interview"`
	Id              string           `json:"id"`
	EmploymentTypes []EmploymentType `json:"employment_types"`
	CompanyLogoUrl  string           `json:"company_logo_url"`
	Skills          []Skill          `json:"skills"`
	Remote          bool             `json:"remote"`
}

type Skill struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
}

type EmploymentType struct {
	Type   string `json:"type"`
	Salary Salary `json:"salary"`
}

type Salary struct {
	From     int    `json:"from"`
	To       int    `json:"to"`
	Currency string `json:"currency"`
}

type ExchangeRates struct {
	Table         string `json:"table"`
	No            string `json:"no"`
	EffectiveDate string `json:"effectiveDate"`
	Rates         []Rate `json:"rates"`
}

type Rate struct {
	Currency string  `json:"currency"`
	Code     string  `json:"code"`
	Mid      float64 `json:"mid"`
}
