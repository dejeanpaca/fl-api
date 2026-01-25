package td

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func FetchStocks(client http.Client) {
	tdUrlStocks := fmt.Sprintf(UrlStocks, ApiKey)

	GetAndWriteURL(client, tdUrlStocks, "store/stocks.json")
}

func FetchEODForSymbol(client http.Client, symbol string) error {
	tdUrlEOD := fmt.Sprintf(UrlEOD, ApiKey)
	reqUrl := fmt.Sprintf("%s&symbol=%s", tdUrlEOD, symbol)
	targetFile := fmt.Sprintf("store/%s.json", symbol)
	err := GetAndWriteURL(client, reqUrl, targetFile)
	return err
}

func FetchApiUsage(client http.Client) error {
	tdUrlUsage := fmt.Sprintf("https://api.twelvedata.com/api_usage?apikey=%s", ApiKey)
	return GetAndWriteURL(client, tdUrlUsage, "store/usage.json")
}

func GetAndWriteURL(client http.Client, url string, targetFile string) error {
	log.Println("url:", url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("req status: ", res.Status, " ", res.StatusCode)

	if res.Body != nil {
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		err = os.WriteFile(targetFile, body, 0666)
		if err != nil {
			return err
		}
	} else {
		return errors.New("response body is nil")
	}

	return nil
}
