package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fl-api/broque/stocks-sync/internal/snp"
)

func main() {
	tdApiKey := "c80090ac9e4144d2bc584dbf2bca1b06"
	tdUrlEOD := "https://api.twelvedata.com/eod?apikey=%s&"
	tdUrlEOD = fmt.Sprintf(tdUrlEOD, tdApiKey)
	tdUrlStocks := "https://api.twelvedata.com/stocks?apikey=%s"
	tdUrlStocks = fmt.Sprintf(tdUrlStocks, tdApiKey)

	symbol := "AAPL"
	reqUrl := fmt.Sprintf("%s&symbol=%s", tdUrlEOD, symbol)

	targetFile := fmt.Sprintf("%s.json", symbol)

	/// get data from reqUrl

	client := http.Client{
		Timeout: time.Second * 10,
	}

	snp.LoadSNP()

	getAndWriteURL(client, tdUrlStocks, "stocks.json")
	getAndWriteURL(client, reqUrl, targetFile)
}

func getAndWriteURL(client http.Client, url string, targetFile string) {
	fmt.Println("url:", url)
	// implement the function to get data from url and write to targetFile
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
		return
	}

	fmt.Println("req status: ", res.Status, " ", res.StatusCode)

	if res.Body != nil {
		defer res.Body.Close()

		body, readErr := io.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
			return
		}

		err := os.WriteFile(targetFile, body, 0666)
		if err != nil {
			log.Fatal(err)
			return
		}
	} else {
		fmt.Println("No response body")
	}
}
