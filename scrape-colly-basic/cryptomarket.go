package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	fName := "cryptocoinmarketcap.csv"
	file, err := os.Create(fName)
	defer file.Close()

	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"Name", "Symbol", "Price (USD)", "Volume (USD)", "Market capacity (USD)", "Change (1h)", "Change (24h)", "Change (7d)"})

	// Instantiate default collector
	c := colly.NewCollector()

	c.OnHTML("table tbody tr", func(e *colly.HTMLElement) {
		buffer := make([]string, 0)

		e.ForEach("td", func(i int, td *colly.HTMLElement) {
			fmt.Println(i, td.Text)
			buffer = append(buffer, td.Text)
		})
		writer.Write(buffer)
	})

	c.Visit("https://coinmarketcap.com/all/views/all/")

	log.Printf("Scraping finished, check file %q for results\n", fName)
}
