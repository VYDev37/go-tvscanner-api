package scanner

import "fmt"

// TVAsset represents a single financial asset retrieved from the TradingView screener.
// It contains raw float64 values for precise calculations and fundamental metrics.
type TVAsset struct {
	// --- 1. Overview & Base Identity ---
	Ticker                  string  `json:"name"`                             // Ticker symbol (e.g., "BBCA")
	Sector                  string  `json:"sector"`                           // Sector classification (e.g., "Finance")
	Market                  string  `json:"market"`                           // Exchange market (e.g., "indonesia")
	FundamentalCurrencyCode string  `json:"fundamental_currency_code"`        // The currency code (e.g., "IDR")
	Description             string  `json:"description"`                      // Company name
	Price                   Money   `json:"close"`                            // Current close price
	Change                  RawNum  `json:"change"`                           // Price change value
	Volume                  RawNum  `json:"volume"`                           // Trading volume
	RelativeVolume10D       Ratio   `json:"relative_volume_10d_calc"`         // Relative volume compared to 10-day average
	MarketCap               Money   `json:"market_cap_basic"`                 // Basic market capitalization
	FreeFloatQty            RawNum  `json:"float_shares_outstanding_current"` // Free Float (in qty)
	FreeFloatPercentage     Percent `json:"float_shares_percent_current"`     // Free Float percentage

	// --- 2. Valuation ---
	PERatio         Ratio `json:"price_earnings_ttm"`              // Price to Earnings (TTM)
	PEGRatio        Ratio `json:"price_earnings_growth_ttm"`       // PEG Ratio (TTM) - PEG < 1 means undervalued
	PSRatio         Ratio `json:"price_sales_current"`             // Price to Sales (Current)
	PBV             Ratio `json:"price_book_fq"`                   // Price to Book Value (FQ)
	PriceToCash     Ratio `json:"price_to_cash_ratio"`             // Price to Cash Ratio
	PriceToFCF      Ratio `json:"price_free_cash_flow_ttm"`        // Price to Free Cash Flow (TTM)
	EnterpriseValue Money `json:"enterprise_value_current"`        // Enterprise Value (Current)
	EVEBITDA        Ratio `json:"enterprise_value_ebitda_ttm"`     // Enterprise Value / EBITDA (TTM)
	EVToRevenue     Ratio `json:"enterprise_value_to_revenue_ttm"` // Enterprise Value to Revenue (TTM)
	EVToEBIT        Ratio `json:"enterprise_value_to_ebit_ttm"`    // Enterprise Value to EBIT (TTM)

	// --- 3. Profitability ---
	GrossMargin     Percent `json:"gross_margin_ttm"`              // Gross Margin (TTM)
	OperatingMargin Percent `json:"operating_margin_ttm"`          // Operating Margin (TTM)
	PreTaxMargin    Percent `json:"pre_tax_margin_ttm"`            // Pre-Tax Margin (TTM)
	NetMargin       Percent `json:"net_margin_ttm"`                // Net Margin (TTM)
	FCFMargin       Percent `json:"free_cash_flow_margin_ttm"`     // Free Cash Flow Margin (TTM)
	ROA             Percent `json:"return_on_assets_fq"`           // Return on Assets (FQ)
	ROE             Percent `json:"return_on_equity_fq"`           // Return on Equity (FQ)
	ROIC            Percent `json:"return_on_invested_capital_fq"` // Return on Invested Capital (FQ)

	// --- 4. Income Statement ---
	TotalRevenue    Money   `json:"total_revenue_ttm"`                         // Total Revenue (TTM)
	RevenueGrowth   Percent `json:"total_revenue_yoy_growth_ttm"`              // Total Revenue YoY Growth (TTM)
	GrossProfit     Money   `json:"gross_profit_ttm"`                          // Gross Profit (TTM)
	OperatingIncome Money   `json:"oper_income_ttm"`                           // Operating Income (TTM)
	NetIncome       Money   `json:"net_income_ttm"`                            // Net Income (TTM)
	EBITDA          Money   `json:"ebitda_ttm"`                                // EBITDA (TTM)
	EPS             Money   `json:"earnings_per_share_diluted_ttm"`            // Earnings Per Share Diluted (TTM)
	EPSGrowth       Percent `json:"earnings_per_share_diluted_yoy_growth_ttm"` // EPS Diluted YoY Growth (TTM)
	FiscalPeriod    string  `json:"fiscal_period_current"`                     // Example: "2026-Q1"
	FiscalPeriodEnd Epoch   `json:"fiscal_period_end_current"`                 // Fiscal end date (unix timestamp)

	// --- 5. Balance Sheet ---
	TotalAssets            Money `json:"total_assets_fq"`                           // Total Assets (FQ)
	TotalCurrentAssets     Money `json:"total_current_assets_fq"`                   // Total Current Assets (FQ)
	CashAndShortTermInvest Money `json:"cash_n_short_term_invest_fq"`               // Cash & Short Term Investments (FQ)
	TotalLiabilities       Money `json:"total_liabilities_fq"`                      // Total Liabilities (FQ)
	TotalDebt              Money `json:"total_debt_fq"`                             // Total Debt (FQ)
	NetDebt                Money `json:"net_debt_fq"`                               // Net Debt (FQ)
	TotalEquity            Money `json:"total_equity_fq"`                           // Total Equity (FQ)
	CurrentRatio           Ratio `json:"current_ratio_fq"`                          // Current Ratio (FQ) - Current Assets vs Current Debt
	QuickRatio             Ratio `json:"quick_ratio_fq"`                            // Quick Ratio (FQ) - Strict liquidity metric
	DERatio                Ratio `json:"debt_to_equity_fq"`                         // Debt to Equity Ratio (FQ)
	CashToTotalDebt        Ratio `json:"cash_n_short_term_invest_to_total_debt_fq"` // Cash to Total Debt Ratio (FQ)

	// --- 6. Cashflow ---
	OperatingCashflow   Money `json:"cash_f_operating_activities_ttm"` // Cash from Operating Activities (TTM)
	InvestingCashflow   Money `json:"cash_f_investing_activities_ttm"` // Cash from Investing Activities (TTM)
	FinancingCashflow   Money `json:"cash_f_financing_activities_ttm"` // Cash from Financing Activities (TTM)
	FreeCashflow        Money `json:"free_cash_flow_ttm"`              // Free Cash Flow (TTM)
	CapitalExpenditures Money `json:"neg_capital_expenditures_ttm"`    // Capital Expenditures (TTM)

	// --- 7. Dividends ---
	CurrentDividendYield Percent `json:"dividends_yield_current"`                   // Current Dividend Yield
	DividendPayoutRatio  Percent `json:"dividend_payout_ratio_ttm"`                 // Dividend Payout Ratio (TTM)
	DPS                  Money   `json:"dps_common_stock_prim_issue_fy"`            // Dividends Per Share (FY)
	DPSYoYGrowth         Percent `json:"dps_common_stock_prim_issue_yoy_growth_fy"` // Dividends Per Share YoY Growth (FY)

	// --- 8. Technicals & Momentum ---
	RSI               RawNum `json:"RSI"`           // Relative Strength Index (14)
	Momentum          RawNum `json:"Mom"`           // Momentum indicator
	AwesomeOscillator RawNum `json:"AO"`            // Awesome Oscillator
	CCI20             RawNum `json:"CCI20"`         // Commodity Channel Index (20)
	StochK            RawNum `json:"Stoch.K"`       // Stochastic %K
	StochD            RawNum `json:"Stoch.D"`       // Stochastic %D
	TechRating1D      RawNum `json:"TechRating_1D"` // AI Tech Rating 1D (-1 Strong Sell to 1 Strong Buy)
}

type Percent float64
type Ratio float64
type Money float64
type RawNum float64
type Epoch int64

func (p Percent) String() string { return fmt.Sprintf("%.2f%%", p) }
func (r Ratio) String() string   { return fmt.Sprintf("%.2fx", r) }
func (m Money) String() string {
	if m >= 1_000_000_000_000 {
		return fmt.Sprintf("%.2f T", float64(m)/1_000_000_000_000)
	} else if m >= 1_000_000_000 {
		return fmt.Sprintf("%.2f M", float64(m)/1_000_000_000)
	}
	return fmt.Sprintf("%.2f", m)
}
func (n RawNum) String() string { return fmt.Sprintf("%.2f", n) }
