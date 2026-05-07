// Testing goes here
// examples/basic/main.go

package main

import (
	"fmt"
	"slices"

	"github.com/VYDev37/go-tvscanner-api/pkg/scanner"
	"github.com/VYDev37/go-tvscanner-api/pkg/utils"
)

func main() {
	// Fetch the top 1000 stocks in the Indonesian Market
	data, err := scanner.FetchStockData(scanner.FetcherOptions{
		Market: "indonesia",
		Limit:  1000,
	})
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}

	// Example: Filter specific tickers
	target := []string{"ADRO", "BBCA", "BBRI"}
	filteredData := utils.Filter(data, func(it scanner.TVAsset) bool {
		return slices.Contains(target, it.Ticker)
	})

	fmt.Println("\nFetching data:")
	fmt.Println("=====================================================")

	for i, it := range filteredData {
		fmt.Printf("\n[%d] %s \n", i+1, it.Ticker)
		fmt.Printf(" ├─ Sector    : %s\n", it.Sector)
		fmt.Printf(" ├─ Price     : %s %s\n", it.FundamentalCurrencyCode, it.Price)
		fmt.Printf(" ├─ P/E Ratio : %s\n", it.PERatio)
		fmt.Printf(" ├─ PBV Ratio : %s\n", it.PBV)
		fmt.Printf(" └─ ROE       : %s\n", it.ROE)
	}
}
