package main

import (
	"fmt"
	scan "tv-screener-test/scanner"
)

func main() {
	data, err := scan.FetchIDXData()
	if err != nil {
		fmt.Println(err)
	}
	for i, it := range data {
		fmt.Printf("%d. %v\nMarket Cap: %v\nNetIncome: %v\nPrice: %v.\nP/E: %v\nEPS: %v\nDividen: %v\nFundamental Currency Code: %v\n", i, it.Ticker, it.MarketCap, it.NetIncome, it.Price, it.PERatio, it.EPS, it.CurrentDividenYield, it.FundamentalCurrencyCode)
	}
}
