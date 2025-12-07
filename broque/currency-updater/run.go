package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	eurURL := "https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/eur.min.json"
	targetFile := "/var/www/fiestalabs/html/currencies/eur.json"

	log.Println("fl > Updating currencies from: " + eurURL)

	spaceClient := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, eurURL, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	writeErr := os.WriteFile(targetFile, body, os.ModeExclusive)
	if err != nil {
		log.Fatal(writeErr)
	}	}



	log.Println("Done updating currencies")
}
