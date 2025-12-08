package snp

import (
	"encoding/json"
	"log"
	"os"
)

// loads snp constituents from a json list
func loadSNP() {
	// Let's first read the `config.json` file
	content, err := os.ReadFile("data/s&p500-constituents.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshall the data into `payload`
	var payload []SNPConstituent
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
}
