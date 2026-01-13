package types

import "time"

type SNPConstituent struct {
	Symbol               string    `json:"Symbol"`
	Security             string    `json:"Security"`
	GICSSector           string    `json:"GICS Sector"`
	GICSSubIndustry      string    `json:"GICS SubIndustry"`
	HeadquartersLocation string    `json:"Headquarters Location"`
	DateAdded            time.Time `json:"Date Added"`
	CIK                  int       `json:"CIK"`
	Founded              int       `json:"Founded"`
}

type TDEODStock struct {
	Symbol    string    `json:"symbol"`
	Exchange  string    `json:"exchange"`
	MicCode   string    `json:"micCode"`
	Currency  string    `json:"currency"`
	Datetime  string    `json:"datetime"`
	Timestamp time.Time `json:"timestamp"`
	Close     float64   `json:"close"`
}
