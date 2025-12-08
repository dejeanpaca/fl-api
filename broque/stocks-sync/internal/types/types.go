package types

import "time"

type SNPConstituent struct {
	symbol               string
	security             string
	gicsSector           string
	gicsSubIndustry      string
	headquartersLocation string
	dateAdded            string
	cik                  int
	founded              int
}

type TDEODStock struct {
	symbol    string
	exchange  string
	micCode   string
	currency  string
	datetime  string
	timestamp time.Time
	close     float64
}
