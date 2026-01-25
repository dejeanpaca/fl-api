package main

import (
	"log"
	"net/http"
	"time"

	"github.com/fl-api/broque/stocks-sync/internal/snp"
	"github.com/fl-api/broque/stocks-sync/internal/td"
)

func main() {
	client := http.Client{
		Timeout: time.Second * 10,
	}

	err := snp.LoadSNP()
	if err != nil {
		panic(err)
	}

	loadStocks(&client)
}

func loadStocks(client *http.Client) {
	count := 0

	for _, symbol := range snp.SNPSymbols {
		log.Println("Fetching EOD for symbol:", symbol.Symbol)
		err := td.FetchEODForSymbol(*client, symbol.Symbol)
		if err != nil {
			log.Println("Error fetching EOD for symbol:", symbol.Symbol, err)
		}

		time.Sleep(100 * time.Millisecond)
		count = count + 1
	}

	log.Println("Fetched EOD for ", count, " symbols")
}
