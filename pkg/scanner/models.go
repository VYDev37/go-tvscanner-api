package scanner

import "fmt"

// TVAsset represents a single financial asset retrieved from the TradingView screener.
// It contains raw float64 values for precise calculations and fundamental metrics.
type TVAsset struct {
	// --- 1. Overview & Base Identity ---
	Ticker                  string  `json:"ticker"`                    // Ticker symbol (e.g., "BBCA")
	Sector                  string  `json:"sector"`                    // Sector classification (e.g., "Finance")
	Market                  string  `json:"market"`                    // Exchange market (e.g., "indonesia")
	FundamentalCurrencyCode string  `json:"fundamental_currency_code"` // The currency code (e.g., "IDR")
	Description             string  `json:"description"`               // Company name
	Price                   Money   `json:"price"`                     // Current close price
	Change                  RawNum  `json:"change"`                    // Price change value
	Volume                  RawNum  `json:"volume"`                    // Trading volume
	RelativeVolume10D       Ratio   `json:"relative_volume_10d"`       // Relative volume compared to 10-day average
	MarketCap               Money   `json:"market_cap"`                // Basic market capitalization
	FreeFloatQty            RawNum  `json:"free_float_qty"`            // Free Float (in qty)
	FreeFloatPercentage     Percent `json:"free_float_percentage"`     // Free Float percentage

	// --- 2. Valuation ---
	PERatio         Ratio `json:"pe_ratio"`         // Price to Earnings (TTM)
	PEGRatio        Ratio `json:"peg_ratio"`        // PEG Ratio (TTM) - PEG < 1 means undervalued
	PSRatio         Ratio `json:"ps_ratio"`         // Price to Sales (Current)
	PBV             Ratio `json:"pbv"`              // Price to Book Value (FQ)
	PriceToCash     Ratio `json:"price_to_cash"`    // Price to Cash Ratio
	PriceToFCF      Ratio `json:"price_to_fcf"`     // Price to Free Cash Flow (TTM)
	EnterpriseValue Money `json:"enterprise_value"` // Enterprise Value (Current)
	EVEBITDA        Ratio `json:"ev_ebitda"`        // Enterprise Value / EBITDA (TTM)
	EVToRevenue     Ratio `json:"ev_to_revenue"`    // Enterprise Value to Revenue (TTM)
	EVToEBIT        Ratio `json:"ev_to_ebit"`       // Enterprise Value to EBIT (TTM)

	// --- 3. Profitability ---
	GrossMargin     Percent `json:"gross_margin"`     // Gross Margin (TTM)
	OperatingMargin Percent `json:"operating_margin"` // Operating Margin (TTM)
	PreTaxMargin    Percent `json:"pre_tax_margin"`   // Pre-Tax Margin (TTM)
	NetMargin       Percent `json:"net_margin"`       // Net Margin (TTM)
	FCFMargin       Percent `json:"fcf_margin"`       // Free Cash Flow Margin (TTM)
	ROA             Percent `json:"roa"`              // Return on Assets (FQ)
	ROE             Percent `json:"roe"`              // Return on Equity (FQ)
	ROIC            Percent `json:"roic"`             // Return on Invested Capital (FQ)

	// --- 4. Income Statement ---
	TotalRevenue    Money   `json:"total_revenue"`     // Total Revenue (TTM)
	RevenueGrowth   Percent `json:"revenue_growth"`    // Total Revenue YoY Growth (TTM)
	GrossProfit     Money   `json:"gross_profit"`      // Gross Profit (TTM)
	OperatingIncome Money   `json:"operating_income"`  // Operating Income (TTM)
	NetIncome       Money   `json:"net_income"`        // Net Income (TTM)
	EBITDA          Money   `json:"ebitda"`            // EBITDA (TTM)
	EPS             Money   `json:"eps"`               // Earnings Per Share Diluted (TTM)
	EPSGrowth       Percent `json:"eps_growth"`        // EPS Diluted YoY Growth (TTM)
	FiscalPeriod    string  `json:"fiscal_period"`     // Example: "2026-Q1"
	FiscalPeriodEnd Epoch   `json:"fiscal_period_end"` // Fiscal end date (unix timestamp)

	// --- 5. Balance Sheet ---
	TotalAssets            Money `json:"total_assets"`               // Total Assets (FQ)
	TotalCurrentAssets     Money `json:"total_current_assets"`       // Total Current Assets (FQ)
	CashAndShortTermInvest Money `json:"cash_and_short_term_invest"` // Cash & Short Term Investments (FQ)
	TotalLiabilities       Money `json:"total_liabilities"`          // Total Liabilities (FQ)
	TotalDebt              Money `json:"total_debt"`                 // Total Debt (FQ)
	NetDebt                Money `json:"net_debt"`                   // Net Debt (FQ)
	TotalEquity            Money `json:"total_equity"`               // Total Equity (FQ)
	CurrentRatio           Ratio `json:"current_ratio"`              // Current Ratio (FQ) - Current Assets vs Current Debt
	QuickRatio             Ratio `json:"quick_ratio"`                // Quick Ratio (FQ) - Strict liquidity metric
	DERatio                Ratio `json:"de_ratio"`                   // Debt to Equity Ratio (FQ)
	CashToTotalDebt        Ratio `json:"cash_to_total_debt"`         // Cash to Total Debt Ratio (FQ)

	// --- 6. Cashflow ---
	OperatingCashflow   Money `json:"operating_cashflow"`   // Cash from Operating Activities (TTM)
	InvestingCashflow   Money `json:"investing_cashflow"`   // Cash from Investing Activities (TTM)
	FinancingCashflow   Money `json:"financing_cashflow"`   // Cash from Financing Activities (TTM)
	FreeCashflow        Money `json:"free_cashflow"`        // Free Cash Flow (TTM)
	CapitalExpenditures Money `json:"capital_expenditures"` // Capital Expenditures (TTM)

	// --- 7. Dividends ---
	CurrentDividendYield Percent `json:"current_dividend_yield"` // Current Dividend Yield
	DividendPayoutRatio  Percent `json:"dividend_payout_ratio"`  // Dividend Payout Ratio (TTM)
	DPS                  Money   `json:"dps"`                    // Dividends Per Share (FY)
	DPSYoYGrowth         Percent `json:"dps_yoy_growth"`         // Dividends Per Share YoY Growth (FY)

	// --- 8. Technicals & Momentum ---
	RSI               RawNum `json:"rsi"`                // Relative Strength Index (14)
	Momentum          RawNum `json:"momentum"`           // Momentum indicator
	AwesomeOscillator RawNum `json:"awesome_oscillator"` // Awesome Oscillator
	CCI20             RawNum `json:"cci20"`              // Commodity Channel Index (20)
	StochK            RawNum `json:"stoch_k"`            // Stochastic %K
	StochD            RawNum `json:"stoch_d"`            // Stochastic %D
	TechRating1D      string `json:"tech_rating_1d"`     // AI Tech Rating 1D (-1 Strong Sell to 1 Strong Buy)
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
