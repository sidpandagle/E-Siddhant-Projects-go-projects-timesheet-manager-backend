package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly/v2"
	"github.com/gofiber/fiber/v2/log"
)

func SaveDataToJSON(fileName string, data interface{}) {
	// log.Info(e.Text)
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Step 3: Write JSON data to a file
	err = os.WriteFile(fileName+".json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}
}
func GetTickerData() {
	c := colly.NewCollector()
	tickers := []Ticker{}

	c.OnScraped(func(e *colly.Response) {
		log.Info(tickers)
		SaveDataToJSON("ticker", tickers)
	})

	c.OnHTML("table[style='background:transparent;'] tr", func(e *colly.HTMLElement) {
		ticker := Ticker{}
		if e.ChildText("td") != "" && e.ChildText("td + td") != "" {
			ticker.Name = e.ChildText("td + td")
			ticker.Ticker = e.ChildText("td a.external")
			ticker.Country = "INDIA"
			ticker.Alpha2Code = "in"
			tickers = append(tickers, ticker)
		}
	})

	c.Visit("https://en.wikipedia.org/wiki/List_of_companies_listed_on_the_National_Stock_Exchange_of_India")
}

func GetStockData() {
	// tickers, terr := GetTickerToJSON()
	// if terr != nil {
	// 	log.Info(terr)
	// }

	tickers := []Ticker{
		{Name: "20 Microns Limited", Ticker: "20MICRONS", Country: "INDIA", Alpha2Code: "IN"},
		{Name: "21st Century Management Services Limited", Ticker: "21STCENMGM", Country: "INDIA", Alpha2Code: "IN"},
		{Name: "3i Infotech Limited", Ticker: "3IINFOTECH", Country: "INDIA", Alpha2Code: "IN"},
		{Name: "3M India Limited", Ticker: "3MINDIA", Country: "INDIA", Alpha2Code: "IN"},
	}
	counter := 0
	c := colly.NewCollector()

	stocks := []Stock{}
	// Register the OnHTML callback before visiting URLs
	c.OnHTML("fin-streamer[data-field='marketCap']", func(e *colly.HTMLElement) {
		log.Info("OnHTML Called")
		stock := Stock{}
		stock.Ticker = "" // Set the ticker here if needed
		stock.MarketCapital = e.Text
		log.Info(e.Text)
		stocks = append(stocks, stock)
	})

	// Register OnScraped to log after scraping is done
	c.OnScraped(func(r *colly.Response) {
		log.Info("OnScraped Called")
		log.Info(counter, len(tickers), stocks, r.StatusCode)
		counter += 1
	})

	// err := c.Visit("https://finance.yahoo.com/quote/" + tickers[0].Ticker + ".NS/")
	log.Info("https://finance.yahoo.com/")
	err := c.Visit("https://finance.yahoo.com/")
	if err != nil {
		log.Error("Error visiting URL:", err)
	}

}

type Ticker struct {
	Name       string `json:"Name"`
	Ticker     string `json:"ticker"`
	Country    string `json:"country"`
	Alpha2Code string `json:"alpha2Code"`
}

type Stock struct {
	Ticker        string `json:"ticker"`
	MarketCapital string `json:"marketCapital"`
}

func GetTickerToJSON() ([]Ticker, error) {
	// Read the JSON file
	bytes, err := os.ReadFile("ticker.json")
	if err != nil {
		return nil, err // Return an error if reading the file fails
	}

	// Unmarshal JSON data into a slice of Ticker structs
	var tickers []Ticker
	err = json.Unmarshal(bytes, &tickers)
	if err != nil {
		return nil, err // Return an error if unmarshalling fails
	}

	return tickers, nil // Return the tickers slice and nil error
}
