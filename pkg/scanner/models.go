package scanner

import "fmt"

// TVAsset represents a single financial asset retrieved from the TradingView screener.
// It contains raw float64 values for precise calculations and fundamental metrics.
type TVAsset struct {
	// --- 1. Overview & Base Identity ---
	Ticker                  string  // Ticker symbol (e.g., "BBCA")
	Sector                  string  // Sector classification (e.g., "Finance")
	Market                  string  // Exchange market (e.g., "indonesia")
	FundamentalCurrencyCode string  // The currency code (e.g., "IDR")
	Price                   Money   // Current close price
	Change                  RawNum  // Price change value
	Volume                  RawNum  // Trading volume
	RelativeVolume10D       Ratio   // Relative volume compared to 10-day average
	MarketCap               Money   // Basic market capitalization
	FreeFloatQty            RawNum  // Free Float (in qty)
	FreeFloatPercentage     Percent // Free Float percentage

	// --- 2. Valuation ---
	PERatio         Ratio // Price to Earnings (TTM)
	PEGRatio        Ratio // PEG Ratio (TTM) - PEG < 1 means undervalued
	PSRatio         Ratio // Price to Sales (Current)
	PBV             Ratio // Price to Book Value (FQ)
	PriceToCash     Ratio // Price to Cash Ratio
	PriceToFCF      Ratio // Price to Free Cash Flow (TTM)
	EnterpriseValue Money // Enterprise Value (Current)
	EVEBITDA        Ratio // Enterprise Value / EBITDA (TTM)
	EVToRevenue     Ratio // Enterprise Value to Revenue (TTM)
	EVToEBIT        Ratio // Enterprise Value to EBIT (TTM)

	// --- 3. Profitability ---
	GrossMargin     Percent // Gross Margin (TTM)
	OperatingMargin Percent // Operating Margin (TTM)
	PreTaxMargin    Percent // Pre-Tax Margin (TTM)
	NetMargin       Percent // Net Margin (TTM)
	FCFMargin       Percent // Free Cash Flow Margin (TTM)
	ROA             Percent // Return on Assets (FQ)
	ROE             Percent // Return on Equity (FQ)
	ROIC            Percent // Return on Invested Capital (FQ)

	// --- 4. Income Statement ---
	TotalRevenue    Money   // Total Revenue (TTM)
	RevenueGrowth   Percent // Total Revenue YoY Growth (TTM)
	GrossProfit     Money   // Gross Profit (TTM)
	OperatingIncome Money   // Operating Income (TTM)
	NetIncome       Money   // Net Income (TTM)
	EBITDA          Money   // EBITDA (TTM)
	EPS             Money   // Earnings Per Share Diluted (TTM)
	EPSGrowth       Percent // EPS Diluted YoY Growth (TTM)

	// --- 5. Balance Sheet ---
	TotalAssets            Money // Total Assets (FQ)
	TotalCurrentAssets     Money // Total Current Assets (FQ)
	CashAndShortTermInvest Money // Cash & Short Term Investments (FQ)
	TotalLiabilities       Money // Total Liabilities (FQ)
	TotalDebt              Money // Total Debt (FQ)
	NetDebt                Money // Net Debt (FQ)
	TotalEquity            Money // Total Equity (FQ)
	CurrentRatio           Ratio // Current Ratio (FQ) - Current Assets vs Current Debt
	QuickRatio             Ratio // Quick Ratio (FQ) - Strict liquidity metric
	DERatio                Ratio // Debt to Equity Ratio (FQ)
	CashToTotalDebt        Ratio // Cash to Total Debt Ratio (FQ)

	// --- 6. Cashflow ---
	OperatingCashflow   Money // Cash from Operating Activities (TTM)
	InvestingCashflow   Money // Cash from Investing Activities (TTM)
	FinancingCashflow   Money // Cash from Financing Activities (TTM)
	FreeCashflow        Money // Free Cash Flow (TTM)
	CapitalExpenditures Money // Capital Expenditures (TTM)

	// --- 7. Dividends ---
	CurrentDividendYield Percent // Current Dividend Yield
	DividendPayoutRatio  Percent // Dividend Payout Ratio (TTM)
	DPS                  Money   // Dividends Per Share (FY)
	DPSYoYGrowth         Percent // Dividends Per Share YoY Growth (FY)

	// --- 8. Technicals & Momentum ---
	RSI               RawNum // Relative Strength Index (14)
	Momentum          RawNum // Momentum indicator
	AwesomeOscillator RawNum // Awesome Oscillator
	CCI20             RawNum // Commodity Channel Index (20)
	StochK            RawNum // Stochastic %K
	StochD            RawNum // Stochastic %D
	TechRating1D      RawNum // AI Tech Rating 1D (-1 Strong Sell to 1 Strong Buy)
}

type Percent float64
type Ratio float64
type Money float64
type RawNum float64

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
