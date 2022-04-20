package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"encoding/json"
)

type Coin struct {
	Data []struct {
		CirculatingSupply float64       `json:"circulating_supply"`
		CmcRank           int64       `json:"cmc_rank"`
		DateAdded         string      `json:"date_added"`
		ID                int64       `json:"id"`
		LastUpdated       string      `json:"last_updated"`
		MaxSupply         int64       `json:"max_supply"`
		Name              string      `json:"name"`
		NumMarketPairs    int64       `json:"num_market_pairs"`
		Platform          interface{} `json:"platform"`
		Quote             struct {
			Usd struct {
				FullyDilutedMarketCap float64 `json:"fully_diluted_market_cap"`
				LastUpdated           string  `json:"last_updated"`
				MarketCap             float64 `json:"market_cap"`
				MarketCapDominance    float64 `json:"market_cap_dominance"`
				PercentChange1h       float64 `json:"percent_change_1h"`
				PercentChange24h      float64 `json:"percent_change_24h"`
				PercentChange30d      float64 `json:"percent_change_30d"`
				PercentChange60d      float64 `json:"percent_change_60d"`
				PercentChange7d       float64 `json:"percent_change_7d"`
				PercentChange90d      float64 `json:"percent_change_90d"`
				Price                 float64 `json:"price"`
				Volume24h             float64 `json:"volume_24h"`
			} `json:"USD"`
		} `json:"quote"`
		Slug        string   `json:"slug"`
		Symbol      string   `json:"symbol"`
		Tags        []string `json:"tags"`
		TotalSupply float64    `json:"total_supply"`
	} `json:"data"`
	Status struct {
		CreditCount  int64       `json:"credit_count"`
		Elapsed      int64       `json:"elapsed"`
		ErrorCode    int64       `json:"error_code"`
		ErrorMessage interface{} `json:"error_message"`
		Notice       interface{} `json:"notice"`
		Timestamp    string      `json:"timestamp"`
		TotalCount   int64       `json:"total_count"`
	} `json:"status"`
}

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "2")
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "b8bc8fb1-2401-4c24-8d83-5f43483be731")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}
	fmt.Println(resp.Status)
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBody))

	var c Coin
	if err := json.Unmarshal(respBody, &c); err != nil {
		panic(err)
	}

	for i, d := range c.Data {
		fmt.Println(i, d.ID)
	}

  
}
