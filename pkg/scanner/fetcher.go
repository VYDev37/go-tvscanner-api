// pkg/scanner/fetcher.go

package scanner

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// FetcherOptions defines the configuration for the data fetcher.
type FetcherOptions struct {
	Market string // The target market (e.g., "indonesia", "america")
	Offset int    // Starting index for the scan results
	Limit  int    // Maximum number of results to fetch
}

// FetchStockData retrieves fundamental data from TradingView based on the provided options.
func FetchStockData(opts FetcherOptions) ([]TVAsset, error) {
	if opts.Market == "" {
		opts.Market = "indonesia"
	}
	if opts.Limit == 0 {
		opts.Limit = 100
	}

	url := fmt.Sprintf("https://scanner.tradingview.com/%s/scan", opts.Market)

	payload := BuildPayload(opts.Offset, opts.Limit, opts.Market)
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal json payload: %w", err)
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	req.Header.Set("Referer", "https://www.tradingview.com/")
	req.Header.Set("Origin", "https://www.tradingview.com")

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		rawBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("TradingView blocked! Status: %d, %s", resp.StatusCode, string(rawBody))
	}

	var result struct {
		Data []struct {
			S string        `json:"s"`
			D []interface{} `json:"d"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	var assets []TVAsset
	expectedLen := len(DefaultColumns)

	for _, item := range result.Data {
		if len(item.D) < expectedLen {
			continue
		}

		row := make(map[string]interface{})
		for i, colName := range DefaultColumns {
			row[colName] = item.D[i]
		}

		assets = append(assets, TVAsset{
			// --- 1. Overview & Base Identity ---
			Ticker:                  fmt.Sprintf("%v", row["name"]),
			Sector:                  fmt.Sprintf("%v", row["sector"]),
			Market:                  fmt.Sprintf("%v", row["market"]),
			FundamentalCurrencyCode: fmt.Sprintf("%v", row["fundamental_currency_code"]),
			Price:                   Money(convertToFloat(row["close"])),
			Change:                  RawNum(convertToFloat(row["change"])),
			Volume:                  RawNum(convertToFloat(row["volume"])),
			RelativeVolume10D:       Ratio(convertToFloat(row["relative_volume_10d_calc"])),
			MarketCap:               Money(convertToFloat(row["market_cap_basic"])),
			FreeFloatQty:            RawNum(convertToFloat(row["float_shares_outstanding_current"])),
			FreeFloatPercentage:     Percent(convertToFloat(row["float_shares_percent_current"])),

			// --- 2. Valuation ---
			PERatio:         Ratio(convertToFloat(row["price_earnings_ttm"])),
			PEGRatio:        Ratio(convertToFloat(row["price_earnings_growth_ttm"])),
			PSRatio:         Ratio(convertToFloat(row["price_sales_current"])),
			PBV:             Ratio(convertToFloat(row["price_book_fq"])),
			PriceToCash:     Ratio(convertToFloat(row["price_to_cash_ratio"])),
			PriceToFCF:      Ratio(convertToFloat(row["price_free_cash_flow_ttm"])),
			EnterpriseValue: Money(convertToFloat(row["enterprise_value_current"])),
			EVEBITDA:        Ratio(convertToFloat(row["enterprise_value_ebitda_ttm"])),
			EVToRevenue:     Ratio(convertToFloat(row["enterprise_value_to_revenue_ttm"])),
			EVToEBIT:        Ratio(convertToFloat(row["enterprise_value_to_ebit_ttm"])),

			// --- 3. Profitability ---
			GrossMargin:     Percent(convertToFloat(row["gross_margin_ttm"])),
			OperatingMargin: Percent(convertToFloat(row["operating_margin_ttm"])),
			PreTaxMargin:    Percent(convertToFloat(row["pre_tax_margin_ttm"])),
			NetMargin:       Percent(convertToFloat(row["net_margin_ttm"])),
			FCFMargin:       Percent(convertToFloat(row["free_cash_flow_margin_ttm"])),
			ROA:             Percent(convertToFloat(row["return_on_assets_fq"])),
			ROE:             Percent(convertToFloat(row["return_on_equity_fq"])),
			ROIC:            Percent(convertToFloat(row["return_on_invested_capital_fq"])),

			// --- 4. Income Statement ---
			TotalRevenue:    Money(convertToFloat(row["total_revenue_ttm"])),
			RevenueGrowth:   Percent(convertToFloat(row["total_revenue_yoy_growth_ttm"])),
			GrossProfit:     Money(convertToFloat(row["gross_profit_ttm"])),
			OperatingIncome: Money(convertToFloat(row["oper_income_ttm"])),
			NetIncome:       Money(convertToFloat(row["net_income_ttm"])),
			EBITDA:          Money(convertToFloat(row["ebitda_ttm"])),
			EPS:             Money(convertToFloat(row["earnings_per_share_diluted_ttm"])),
			EPSGrowth:       Percent(convertToFloat(row["earnings_per_share_diluted_yoy_growth_ttm"])),

			// --- 5. Balance Sheet ---
			TotalAssets:            Money(convertToFloat(row["total_assets_fq"])),
			TotalCurrentAssets:     Money(convertToFloat(row["total_current_assets_fq"])),
			CashAndShortTermInvest: Money(convertToFloat(row["cash_n_short_term_invest_fq"])),
			TotalLiabilities:       Money(convertToFloat(row["total_liabilities_fq"])),
			TotalDebt:              Money(convertToFloat(row["total_debt_fq"])),
			NetDebt:                Money(convertToFloat(row["net_debt_fq"])),
			TotalEquity:            Money(convertToFloat(row["total_equity_fq"])),
			CurrentRatio:           Ratio(convertToFloat(row["current_ratio_fq"])),
			QuickRatio:             Ratio(convertToFloat(row["quick_ratio_fq"])),
			DERatio:                Ratio(convertToFloat(row["debt_to_equity_fq"])),
			CashToTotalDebt:        Ratio(convertToFloat(row["cash_n_short_term_invest_to_total_debt_fq"])),

			// --- 6. Cashflow ---
			OperatingCashflow:   Money(convertToFloat(row["cash_f_operating_activities_ttm"])),
			InvestingCashflow:   Money(convertToFloat(row["cash_f_investing_activities_ttm"])),
			FinancingCashflow:   Money(convertToFloat(row["cash_f_financing_activities_ttm"])),
			FreeCashflow:        Money(convertToFloat(row["free_cash_flow_ttm"])),
			CapitalExpenditures: Money(convertToFloat(row["neg_capital_expenditures_ttm"])),

			// --- 7. Dividends ---
			CurrentDividendYield: Percent(convertToFloat(row["dividends_yield_current"])),
			DividendPayoutRatio:  Percent(convertToFloat(row["dividend_payout_ratio_ttm"])),
			DPS:                  Money(convertToFloat(row["dps_common_stock_prim_issue_fy"])),
			DPSYoYGrowth:         Percent(convertToFloat(row["dps_common_stock_prim_issue_yoy_growth_fy"])),

			// --- 8. Technicals & Momentum ---
			RSI:               RawNum(convertToFloat(row["RSI"])),
			Momentum:          RawNum(convertToFloat(row["Mom"])),
			AwesomeOscillator: RawNum(convertToFloat(row["AO"])),
			CCI20:             RawNum(convertToFloat(row["CCI20"])),
			StochK:            RawNum(convertToFloat(row["Stoch.K"])),
			StochD:            RawNum(convertToFloat(row["Stoch.D"])),
			TechRating1D:      RawNum(convertToFloat(row["TechRating_1D"])),
		})
	}

	return assets, nil
}

// convertToFloat safely converts an interface value to a float64.
func convertToFloat(i interface{}) float64 {
	if i == nil {
		return 0
	}
	switch v := i.(type) {
	case float64:
		return v
	case int:
		return float64(v)
	default:
		return 0
	}
}
