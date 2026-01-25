package snp

import (
	"encoding/json"
	"log"
	"os"

	"github.com/fl-api/broque/stocks-sync/internal/types"
)

var (
	SNPSymbols []types.SNPConstituent
)

// loads snp constituents from a json list
func LoadSNP() error {
	// Let's first read the `config.json` file
	content, err := os.ReadFile("data/s&p500-constituents.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
		return err
	}

	// Now let's unmarshall the data into `payload`
	var payload []types.SNPConstituent
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
		return err
	}

	SNPSymbols = payload
	log.Println("Loaded ", len(SNPSymbols), " S&P constituents")
	return nil
}
