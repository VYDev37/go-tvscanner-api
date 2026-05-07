# Go TVScanner API 🚀

[![Go Report Card](https://goreportcard.com/badge/github.com/VYDev37/go-tvscanner-api)](https://goreportcard.com/report/github.com/VYDev37/go-tvscanner-api)
[![Go Reference](https://pkg.go.dev/badge/github.com/VYDev37/go-tvscanner-api.svg)](https://pkg.go.dev/github.com/VYDev37/go-tvscanner-api)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**Go TVScanner API** is a lightning-fast, production-ready Go wrapper for the TradingView stock screener. Fetch robust fundamental, technical, and valuation data seamlessly for the Indonesian Stock Exchange (IDX) or any supported global market. 

Built with idiomatic Go, this library focuses on performance, precise numeric data retention, and developer-friendly presentation.

## Features

- 🏎️ **Fast and Lightweight**: No heavy dependencies, standard library driven.
- 📊 **Golden Fundamental Columns**: Ready-to-use access to PE Ratio, PBV, ROE, Dividend Yield, and more.
- 🛠 **Raw Data Integrity**: Underlying values use custom types (`Money`, `Ratio`, `Percent`, `RawNum`) based on `float64` for precise calculations.
- 💬 **Human-Readable Formats**: Built-in `String()` methods automatically format values when printed (e.g. `12.50%`, `2.40x`, `1.20 T`).

## Installation

```bash
go get github.com/VYDev37/go-tvscanner-api
```

## Quick Start

```go
package main

import (
	"fmt"
	"slices"

	"github.com/VYDev37/go-tvscanner-api/pkg/scanner"
	"github.com/VYDev37/go-tvscanner-api/pkg/utils"
)

func main() {
	// Fetch the top 1000 stocks in the Indonesian market
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

	for i, it := range filteredData {
		fmt.Printf("\n[%d] %s \n", i+1, it.Ticker)
		fmt.Printf(" ├─ Sector    : %s\n", it.Sector)
		fmt.Printf(" ├─ Price     : %s %s\n", it.FundamentalCurrencyCode, it.Price)
		fmt.Printf(" ├─ P/E Ratio : %s\n", it.PERatio)
		fmt.Printf(" ├─ PBV Ratio : %s\n", it.PBV)
		fmt.Printf(" └─ ROE       : %s\n", it.ROE)
	}
}
```

## Supported "Golden Columns"

This library automatically fetches a comprehensive payload covering Valuation, Profitability, Income Statement, Balance Sheet, Cashflow, Dividends, and Technicals.

It returns the `TVAsset` struct, populated with all key metrics mapped from TradingView. The underlying types (`Money`, `Ratio`, `Percent`, `RawNum`) implement the `String()` interface for easy printing:

<details>
<summary><strong>1. Overview & Base Identity</strong></summary>

| Metric | Property | Go Type | TV Column |
|--------|----------|---------|-----------|
| Ticker / Name | `Ticker` | `string` | `name` |
| Sector | `Sector` | `string` | `sector` |
| Market | `Market` | `string` | `market` |
| Currency | `FundamentalCurrencyCode` | `string` | `fundamental_currency_code` |
| Close Price | `Price` | `scanner.Money` | `close` |
| Change | `Change` | `scanner.RawNum` | `change` |
| Volume | `Volume` | `scanner.RawNum` | `volume` |
| Relative Vol 10D | `RelativeVolume10D` | `scanner.Ratio` | `relative_volume_10d_calc` |
| Market Cap | `MarketCap` | `scanner.Money` | `market_cap_basic` |
| Free Float (Qty) | `FreeFloatQty` | `scanner.RawNum` | `float_shares_outstanding_current` |
| Free Float (%) | `FreeFloatPercentage` | `scanner.Percent` | `float_shares_percent_current` |

</details>

<details>
<summary><strong>2. Valuation</strong></summary>

| Metric | Property | Go Type | TV Column |
|--------|----------|---------|-----------|
| P/E Ratio | `PERatio` | `scanner.Ratio` | `price_earnings_ttm` |
| PEG Ratio | `PEGRatio` | `scanner.Ratio` | `price_earnings_growth_ttm` |
| P/S Ratio | `PSRatio` | `scanner.Ratio` | `price_sales_current` |
| P/B Ratio | `PBV` | `scanner.Ratio` | `price_book_fq` |
| Price to Cash | `PriceToCash` | `scanner.Ratio` | `price_to_cash_ratio` |
| Price to FCF | `PriceToFCF` | `scanner.Ratio` | `price_free_cash_flow_ttm` |
| Enterprise Value | `EnterpriseValue` | `scanner.Money` | `enterprise_value_current` |
| EV/EBITDA | `EVEBITDA` | `scanner.Ratio` | `enterprise_value_ebitda_ttm` |
| EV/Revenue | `EVToRevenue` | `scanner.Ratio` | `enterprise_value_to_revenue_ttm` |
| EV/EBIT | `EVToEBIT` | `scanner.Ratio` | `enterprise_value_to_ebit_ttm` |

</details>

<details>
<summary><strong>3. Profitability</strong></summary>

| Metric | Property | Go Type | TV Column |
|--------|----------|---------|-----------|
| Gross Margin | `GrossMargin` | `scanner.Percent` | `gross_margin_ttm` |
| Operating Margin | `OperatingMargin` | `scanner.Percent` | `operating_margin_ttm` |
| Pre-Tax Margin | `PreTaxMargin` | `scanner.Percent` | `pre_tax_margin_ttm` |
| Net Margin | `NetMargin` | `scanner.Percent` | `net_margin_ttm` |
| FCF Margin | `FCFMargin` | `scanner.Percent` | `free_cash_flow_margin_ttm` |
| ROA | `ROA` | `scanner.Percent` | `return_on_assets_fq` |
| ROE | `ROE` | `scanner.Percent` | `return_on_equity_fq` |
| ROIC | `ROIC` | `scanner.Percent` | `return_on_invested_capital_fq` |

</details>

<details>
<summary><strong>4. Income Statement</strong></summary>

| Metric | Property | Go Type | TV Column |
|--------|----------|---------|-----------|
| Total Revenue | `TotalRevenue` | `scanner.Money` | `total_revenue_ttm` |
| Revenue Growth | `RevenueGrowth` | `scanner.Percent` | `total_revenue_yoy_growth_ttm` |
| Gross Profit | `GrossProfit` | `scanner.Money` | `gross_profit_ttm` |
| Operating Income | `OperatingIncome` | `scanner.Money` | `oper_income_ttm` |
| Net Income | `NetIncome` | `scanner.Money` | `net_income_ttm` |
| EBITDA | `EBITDA` | `scanner.Money` | `ebitda_ttm` |
| EPS (Diluted) | `EPS` | `scanner.Money` | `earnings_per_share_diluted_ttm` |
| EPS Growth | `EPSGrowth` | `scanner.Percent` | `earnings_per_share_diluted_yoy_growth_ttm` |

</details>

<details>
<summary><strong>5. Balance Sheet</strong></summary>

| Metric | Property | Go Type | TV Column |
|--------|----------|---------|-----------|
| Total Assets | `TotalAssets` | `scanner.Money` | `total_assets_fq` |
| Total Current Assets | `TotalCurrentAssets` | `scanner.Money` | `total_current_assets_fq` |
| Cash & ST Invest | `CashAndShortTermInvest` | `scanner.Money` | `cash_n_short_term_invest_fq` |
| Total Liabilities | `TotalLiabilities` | `scanner.Money` | `total_liabilities_fq` |
| Total Debt | `TotalDebt` | `scanner.Money` | `total_debt_fq` |
| Net Debt | `NetDebt` | `scanner.Money` | `net_debt_fq` |
| Total Equity | `TotalEquity` | `scanner.Money` | `total_equity_fq` |
| Current Ratio | `CurrentRatio` | `scanner.Ratio` | `current_ratio_fq` |
| Quick Ratio | `QuickRatio` | `scanner.Ratio` | `quick_ratio_fq` |
| D/E Ratio | `DERatio` | `scanner.Ratio` | `debt_to_equity_fq` |
| Cash/Total Debt | `CashToTotalDebt` | `scanner.Ratio` | `cash_n_short_term_invest_to_total_debt_fq` |

</details>

<details>
<summary><strong>6. Cashflow</strong></summary>

| Metric | Property | Go Type | TV Column |
|--------|----------|---------|-----------|
| Operating Cashflow | `OperatingCashflow` | `scanner.Money` | `cash_f_operating_activities_ttm` |
| Investing Cashflow | `InvestingCashflow` | `scanner.Money` | `cash_f_investing_activities_ttm` |
| Financing Cashflow | `FinancingCashflow` | `scanner.Money` | `cash_f_financing_activities_ttm` |
| Free Cashflow | `FreeCashflow` | `scanner.Money` | `free_cash_flow_ttm` |
| Capital Expenditures | `CapitalExpenditures` | `scanner.Money` | `neg_capital_expenditures_ttm` |

</details>

<details>
<summary><strong>7. Dividends</strong></summary>

| Metric | Property | Go Type | TV Column |
|--------|----------|---------|-----------|
| Dividend Yield | `CurrentDividendYield` | `scanner.Percent` | `dividends_yield_current` |
| Payout Ratio | `DividendPayoutRatio` | `scanner.Percent` | `dividend_payout_ratio_ttm` |
| DPS | `DPS` | `scanner.Money` | `dps_common_stock_prim_issue_fy` |
| DPS Growth | `DPSYoYGrowth` | `scanner.Percent` | `dps_common_stock_prim_issue_yoy_growth_fy` |

</details>

<details>
<summary><strong>8. Technicals & Momentum</strong></summary>

| Metric | Property | Go Type | TV Column |
|--------|----------|---------|-----------|
| RSI | `RSI` | `scanner.RawNum` | `RSI` |
| Momentum | `Momentum` | `scanner.RawNum` | `Mom` |
| Awesome Oscillator | `AwesomeOscillator` | `scanner.RawNum` | `AO` |
| CCI20 | `CCI20` | `scanner.RawNum` | `CCI20` |
| Stochastic %K | `StochK` | `scanner.RawNum` | `Stoch.K` |
| Stochastic %D | `StochD` | `scanner.RawNum` | `Stoch.D` |
| Tech Rating 1D | `TechRating1D` | `scanner.RawNum` | `TechRating_1D` |

</details>

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.