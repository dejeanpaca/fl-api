package types

import (
	"encoding/json"
	"strconv"
	"time"
)

type SNPConstituent struct {
	Symbol               string             `json:"Symbol"`
	Security             string             `json:"Security"`
	GICSSector           string             `json:"GICS Sector"`
	GICSSubIndustry      string             `json:"GICS SubIndustry"`
	HeadquartersLocation string             `json:"Headquarters Location"`
	DateAdded            SNPConstituentTime `json:"Date Added"`
	CIK                  int                `json:"CIK"`
	Founded              string             `json:"Founded"`
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

type SNPConstituentTime struct {
	time.Time
}

func (t *SNPConstituentTime) UnmarshalJSON(b []byte) (err error) {
	date, err := time.Parse(`"2006-01-02"`, string(b))
	if err != nil {
		t.Time = time.Now()
		return nil
	}
	t.Time = date
	return
}

func (p *SNPConstituent) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	var symbol string
	err = json.Unmarshal(*objMap["Symbol"], &symbol)
	if err != nil {
		return err
	}
	p.Symbol = symbol

	var cik int
	err = json.Unmarshal(*objMap["CIK"], &cik)
	if err != nil {
		return err
	}
	p.CIK = cik

	var foundedInt int
	err = json.Unmarshal(*objMap["Founded"], &foundedInt)
	if err != nil {
		// age  is string
		var foundedString string
		err = json.Unmarshal(*objMap["Founded"], &foundedString)
		if err != nil {
			return err
		}
		p.Founded = foundedString

	} else {
		p.Founded = strconv.Itoa(foundedInt)
	}

	p.DateAdded = SNPConstituentTime{
		Time: time.Time{},
	}

	if objMap["Date Added"] != nil {
		_ = json.Unmarshal(*objMap["Date Added"], &p.DateAdded)
	}

	var security string
	err = json.Unmarshal(*objMap["Security"], &security)
	if err != nil {
		return err
	}
	p.Security = security

	var gicsSector string
	err = json.Unmarshal(*objMap["GICS Sector"], &gicsSector)
	if err != nil {
		return err
	}
	p.GICSSector = gicsSector

	if objMap["GICS SubIndustry"] != nil {
		var gicsSubIndustry string
		err = json.Unmarshal(*objMap["GICS SubIndustry"], &gicsSubIndustry)
		if err != nil {
			return err
		}
		p.GICSSubIndustry = gicsSubIndustry
	}

	var headquartersLocation string
	err = json.Unmarshal(*objMap["Headquarters Location"], &headquartersLocation)
	if err != nil {
		return err
	}
	p.HeadquartersLocation = headquartersLocation

	return nil
}
