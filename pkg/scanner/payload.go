// pkg/scanner/payload.go

package scanner

// extracted raw data (use as reference in the future, if you'd like to develop this or would like to insert to your project.)
// valuation
/*
{"columns":["ticker-view","market_cap_basic","type","typespecs","fundamental_currency_code","Perf.1Y.MarketCap","price_earnings_ttm","price_earnings_growth_ttm","price_sales_current","price_book_fq","price_to_cash_f_operating_activities_ttm","price_free_cash_flow_ttm","price_to_cash_ratio","enterprise_value_current","enterprise_value_to_revenue_ttm","enterprise_value_to_ebit_ttm","enterprise_value_ebitda_ttm"],"filter":[{"left":"is_primary","operation":"equal","right":true}],"ignore_unknown_fields":false,"options":{"lang":"en"},"range":[0,100],"sort":{"sortBy":"market_cap_basic","sortOrder":"desc"},"symbols":{},"markets":["indonesia"],"filter2":{"operator":"and","operands":[{"operation":{"operator":"or","operands":[{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"stock"}},{"expression":{"left":"typespecs","operation":"has","right":["common"]}}]}},{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"stock"}},{"expression":{"left":"typespecs","operation":"has","right":["preferred"]}}]}},{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"dr"}}]}},{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"fund"}},{"expression":{"left":"typespecs","operation":"has_none_of","right":["etf","mutual","closedend"]}}]}}]}},{"expression":{"left":"typespecs","operation":"has_none_of","right":["pre-ipo"]}}]}}
*/

// overview
/*
{"columns":["ticker-view","close","type","typespecs","pricescale","minmov","fractional","minmove2","currency","change","volume","relative_volume_10d_calc","market_cap_basic","fundamental_currency_code","price_earnings_ttm","earnings_per_share_diluted_ttm","earnings_per_share_diluted_yoy_growth_ttm","dividends_yield_current","sector.tr","market","sector","AnalystRating","AnalystRating.tr"],"filter":[{"left":"is_primary","operation":"equal","right":true}],"ignore_unknown_fields":false,"options":{"lang":"en"},"range":[0,100],"sort":{"sortBy":"market_cap_basic","sortOrder":"desc"},"symbols":{},"markets":["indonesia"],"filter2":{"operator":"and","operands":[{"operation":{"operator":"or","operands":[{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"stock"}},{"expression":{"left":"typespecs","operation":"has","right":["common"]}}]}},{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"stock"}},{"expression":{"left":"typespecs","operation":"has","right":["preferred"]}}]}},{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"dr"}}]}},{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"fund"}},{"expression":{"left":"typespecs","operation":"has_none_of","right":["etf","mutual","closedend"]}}]}}]}},{"expression":{"left":"typespecs","operation":"has_none_of","right":["pre-ipo"]}}]}}
*/

// extended hours
/*
{"columns":["ticker-view","premarket_close","type","typespecs","pricescale","minmov","fractional","minmove2","currency","premarket_change","premarket_gap","premarket_volume","close","change","gap","volume","volume_change","postmarket_close","postmarket_change","postmarket_volume"],"filter":[{"left":"is_primary","operation":"equal","right":true}],"ignore_unknown_fields":false,"options":{"lang":"en"},"range":[0,100],"sort":{"sortBy":"market_cap_basic","sortOrder":"desc"},"symbols":{},"markets":["indonesia"],"filter2":{"operator":"and","operands":[{"operation":{"operator":"or","operands":[{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"stock"}},{"expression":{"left":"typespecs","operation":"has","right":["common"]}}]}},{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"stock"}},{"expression":{"left":"typespecs","operation":"has","right":["preferred"]}}]}},{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"dr"}}]}},{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"fund"}},{"expression":{"left":"typespecs","operation":"has_none_of","right":["etf","mutual","closedend"]}}]}}]}},{"expression":{"left":"typespecs","operation":"has_none_of","right":["pre-ipo"]}}]}}
*/

// dividends
/*
{"columns":["ticker-view","dps_common_stock_prim_issue_fy","type","typespecs","fundamental_currency_code","dps_common_stock_prim_issue_fq","dividends_yield_current","dividends_yield","dividend_payout_ratio_ttm","dps_common_stock_prim_issue_yoy_growth_fy","continuous_dividend_payout","continuous_dividend_growth"],"filter":[{"left":"is_primary","operation":"equal","right":true}],"ignore_unknown_fields":false,"options":{"lang":"en"},"range":[0,100],"sort":{"sortBy":"market_cap_basic","sortOrder":"desc"},"symbols":{},"markets":["indonesia"],"filter2":{"operator":"and","operands":[{"operation":{"operator":"or","operands":[{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"stock"}},{"expression":{"left":"typespecs","operation":"has","right":["common"]}}]}},{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"stock"}},{"expression":{"left":"typespecs","operation":"has","right":["preferred"]}}]}},{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"dr"}}]}},{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"fund"}},{"expression":{"left":"typespecs","operation":"has_none_of","right":["etf","mutual","closedend"]}}]}}]}},{"expression":{"left":"typespecs","operation":"has_none_of","right":["pre-ipo"]}}]}}
*/

// profitability
/*
{"columns":["ticker-view","gross_margin_ttm","operating_margin_ttm","pre_tax_margin_ttm","net_margin_ttm","free_cash_flow_margin_ttm","return_on_assets_fq","return_on_equity_fq","return_on_invested_capital_fq","research_and_dev_ratio_ttm","sell_gen_admin_exp_other_ratio_ttm"],"filter":[{"left":"is_primary","operation":"equal","right":true}],"ignore_unknown_fields":false,"options":{"lang":"en"},"range":[0,100],"sort":{"sortBy":"market_cap_basic","sortOrder":"desc"},"symbols":{},"markets":["indonesia"],"filter2":{"operator":"and","operands":[{"operation":{"operator":"or","operands":[{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"stock"}},{"expression":{"left":"typespecs","operation":"has","right":["common"]}}]}},{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"stock"}},{"expression":{"left":"typespecs","operation":"has","right":["preferred"]}}]}},{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"dr"}}]}},{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"fund"}},{"expression":{"left":"typespecs","operation":"has_none_of","right":["etf","mutual","closedend"]}}]}}]}},{"expression":{"left":"typespecs","operation":"has_none_of","right":["pre-ipo"]}}]}}
*/

// Income statement
/*
{"columns":["ticker-view","fiscal_period_current","fiscal_period_end_current","total_revenue_ttm","type","typespecs","fundamental_currency_code","total_revenue_yoy_growth_ttm","gross_profit_ttm","oper_income_ttm","net_income_ttm","ebitda_ttm","earnings_per_share_diluted_ttm","earnings_per_share_diluted_yoy_growth_ttm"],"filter":[{"left":"is_primary","operation":"equal","right":true}],"ignore_unknown_fields":false,"options":{"lang":"en"},"range":[0,100],"sort":{"sortBy":"market_cap_basic","sortOrder":"desc"},"symbols":{},"markets":["indonesia"],"filter2":{"operator":"and","operands":[{"operation":{"operator":"or","operands":[{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"stock"}},{"expression":{"left":"typespecs","operation":"has","right":["common"]}}]}},{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"stock"}},{"expression":{"left":"typespecs","operation":"has","right":["preferred"]}}]}},{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"dr"}}]}},{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"fund"}},{"expression":{"left":"typespecs","operation":"has_none_of","right":["etf","mutual","closedend"]}}]}}]}},{"expression":{"left":"typespecs","operation":"has_none_of","right":["pre-ipo"]}}]}}
*/

// balance sheet
/*
{"columns":["ticker-view","fiscal_period_current","fiscal_period_end_current","total_assets_fq","type","typespecs","fundamental_currency_code","total_current_assets_fq","cash_n_short_term_invest_fq","total_liabilities_fq","total_debt_fq","net_debt_fq","total_equity_fq","current_ratio_fq","quick_ratio_fq","debt_to_equity_fq","cash_n_short_term_invest_to_total_debt_fq"],"filter":[{"left":"is_primary","operation":"equal","right":true}],"ignore_unknown_fields":false,"options":{"lang":"en"},"range":[0,100],"sort":{"sortBy":"market_cap_basic","sortOrder":"desc"},"symbols":{},"markets":["indonesia"],"filter2":{"operator":"and","operands":[{"operation":{"operator":"or","operands":[{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"stock"}},{"expression":{"left":"typespecs","operation":"has","right":["common"]}}]}},{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"stock"}},{"expression":{"left":"typespecs","operation":"has","right":["preferred"]}}]}},{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"dr"}}]}},{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"fund"}},{"expression":{"left":"typespecs","operation":"has_none_of","right":["etf","mutual","closedend"]}}]}}]}},{"expression":{"left":"typespecs","operation":"has_none_of","right":["pre-ipo"]}}]}}
*/

// cashflow
/*
{"columns":["ticker-view","fiscal_period_current","fiscal_period_end_current","cash_f_operating_activities_ttm","type","typespecs","fundamental_currency_code","cash_f_investing_activities_ttm","cash_f_financing_activities_ttm","free_cash_flow_ttm","neg_capital_expenditures_ttm"],"filter":[{"left":"is_primary","operation":"equal","right":true}],"ignore_unknown_fields":false,"options":{"lang":"en"},"range":[0,100],"sort":{"sortBy":"market_cap_basic","sortOrder":"desc"},"symbols":{},"markets":["indonesia"],"filter2":{"operator":"and","operands":[{"operation":{"operator":"or","operands":[{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"stock"}},{"expression":{"left":"typespecs","operation":"has","right":["common"]}}]}},{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"stock"}},{"expression":{"left":"typespecs","operation":"has","right":["preferred"]}}]}},{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"dr"}}]}},{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"fund"}},{"expression":{"left":"typespecs","operation":"has_none_of","right":["etf","mutual","closedend"]}}]}}]}},{"expression":{"left":"typespecs","operation":"has_none_of","right":["pre-ipo"]}}]}}
*/

// per share
/*
{"columns":["ticker-view","revenue_per_share_ttm","type","typespecs","fundamental_currency_code","earnings_per_share_basic_ttm","earnings_per_share_diluted_ttm","operating_cash_flow_per_share_ttm","free_cash_flow_per_share_ttm","ebit_per_share_ttm","ebitda_per_share_ttm","book_value_per_share_fq","total_debt_per_share_fq","cash_per_share_fq"],"filter":[{"left":"is_primary","operation":"equal","right":true}],"ignore_unknown_fields":false,"options":{"lang":"en"},"range":[0,100],"sort":{"sortBy":"market_cap_basic","sortOrder":"desc"},"symbols":{},"markets":["indonesia"],"filter2":{"operator":"and","operands":[{"operation":{"operator":"or","operands":[{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"stock"}},{"expression":{"left":"typespecs","operation":"has","right":["common"]}}]}},{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"stock"}},{"expression":{"left":"typespecs","operation":"has","right":["preferred"]}}]}},{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"dr"}}]}},{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"fund"}},{"expression":{"left":"typespecs","operation":"has_none_of","right":["etf","mutual","closedend"]}}]}}]}},{"expression":{"left":"typespecs","operation":"has_none_of","right":["pre-ipo"]}}]}}
*/

// technicals
/*
{"columns":["ticker-view","TechRating_1D","TechRating_1D.tr","MARating_1D","MARating_1D.tr","OsRating_1D","OsRating_1D.tr","RSI","Mom","pricescale","minmov","fractional","minmove2","AO","CCI20","Stoch.K","Stoch.D","Candle.3BlackCrows","Candle.3WhiteSoldiers","Candle.AbandonedBaby.Bearish","Candle.AbandonedBaby.Bullish","Candle.Doji","Candle.Doji.Dragonfly","Candle.Doji.Gravestone","Candle.Engulfing.Bearish","Candle.Engulfing.Bullish","Candle.EveningStar","Candle.Hammer","Candle.HangingMan","Candle.Harami.Bearish","Candle.Harami.Bullish","Candle.InvertedHammer","Candle.Kicking.Bearish","Candle.Kicking.Bullish","Candle.LongShadow.Lower","Candle.LongShadow.Upper","Candle.Marubozu.Black","Candle.Marubozu.White","Candle.MorningStar","Candle.ShootingStar","Candle.SpinningTop.Black","Candle.SpinningTop.White","Candle.TriStar.Bearish","Candle.TriStar.Bullish"],"filter":[{"left":"is_primary","operation":"equal","right":true}],"ignore_unknown_fields":false,"options":{"lang":"en"},"range":[0,100],"sort":{"sortBy":"market_cap_basic","sortOrder":"desc"},"symbols":{},"markets":["indonesia"],"filter2":{"operator":"and","operands":[{"operation":{"operator":"or","operands":[{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"stock"}},{"expression":{"left":"typespecs","operation":"has","right":["common"]}}]}},{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"stock"}},{"expression":{"left":"typespecs","operation":"has","right":["preferred"]}}]}},{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"dr"}}]}},{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"fund"}},{"expression":{"left":"typespecs","operation":"has_none_of","right":["etf","mutual","closedend"]}}]}}]}},{"expression":{"left":"typespecs","operation":"has_none_of","right":["pre-ipo"]}}]}}
*/

/*
{"columns":["ticker-view","close","type","typespecs","pricescale","minmov","fractional","minmove2","currency","change","volume","relative_volume_10d_calc","market_cap_basic","fundamental_currency_code","price_earnings_ttm","earnings_per_share_diluted_ttm","earnings_per_share_diluted_yoy_growth_ttm","dividends_yield_current","sector.tr","market","sector","AnalystRating","AnalystRating.tr"],"filter":[{"left":"float_shares_percent_current","operation":"in_range","right":[0,100]},{"left":"is_primary","operation":"equal","right":true}],"ignore_unknown_fields":false,"options":{"lang":"en"},"range":[0,100],"sort":{"sortBy":"market_cap_basic","sortOrder":"desc"},"symbols":{},"markets":["america"],"filter2":{"operator":"and","operands":[{"operation":{"operator":"or","operands":[{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"stock"}},{"expression":{"left":"typespecs","operation":"has","right":["common"]}}]}},{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"stock"}},{"expression":{"left":"typespecs","operation":"has","right":["preferred"]}}]}},{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"dr"}}]}},{"operation":{"operator":"and","operands":[{"expression":{"left":"type","operation":"equal","right":"fund"}},{"expression":{"left":"typespecs","operation":"has_none_of","right":["etf","mutual","closedend"]}}]}}]}},{"expression":{"left":"typespecs","operation":"has_none_of","right":["pre-ipo"]}}]}}
*/

// M is a convenience type for map[string]interface{}
type M map[string]interface{}

// DefaultColumns defines the fundamental data points to be fetched from TradingView.
var DefaultColumns = []string{
	// 1. Overview & Base Identity
	"name", "sector", "market", "fundamental_currency_code",
	"close", "change", "volume", "relative_volume_10d_calc",
	"market_cap_basic", "float_shares_outstanding_current", "float_shares_percent_current",

	// 2. Valuation
	"price_earnings_ttm", "price_earnings_growth_ttm", "price_sales_current", "price_book_fq",
	"price_to_cash_ratio", "price_free_cash_flow_ttm",
	"enterprise_value_current", "enterprise_value_ebitda_ttm", "enterprise_value_to_revenue_ttm", "enterprise_value_to_ebit_ttm",

	// 3. Profitability
	"gross_margin_ttm", "operating_margin_ttm", "pre_tax_margin_ttm", "net_margin_ttm",
	"free_cash_flow_margin_ttm", "return_on_assets_fq", "return_on_equity_fq", "return_on_invested_capital_fq",

	// 4. Income Statement
	"total_revenue_ttm", "total_revenue_yoy_growth_ttm", "gross_profit_ttm", "oper_income_ttm", "net_income_ttm", "ebitda_ttm",
	"earnings_per_share_diluted_ttm", "earnings_per_share_diluted_yoy_growth_ttm",

	// 5. Balance Sheet
	"total_assets_fq", "total_current_assets_fq", "cash_n_short_term_invest_fq",
	"total_liabilities_fq", "total_debt_fq", "net_debt_fq", "total_equity_fq",
	"current_ratio_fq", "quick_ratio_fq", "debt_to_equity_fq", "cash_n_short_term_invest_to_total_debt_fq",

	// 6. Cashflow
	"cash_f_operating_activities_ttm", "cash_f_investing_activities_ttm", "cash_f_financing_activities_ttm",
	"free_cash_flow_ttm", "neg_capital_expenditures_ttm",

	// 7. Dividends
	"dividends_yield_current", "dividend_payout_ratio_ttm",
	"dps_common_stock_prim_issue_fy", "dps_common_stock_prim_issue_yoy_growth_fy",

	// 8. Technicals & Momentum
	"RSI", "Mom", "AO", "CCI20", "Stoch.K", "Stoch.D", "TechRating_1D",
}

// BuildPayload constructs the JSON payload required by the TradingView API.
// It applies necessary filters to retrieve stock fundamentals.
func BuildPayload(offset, limit int, market string) M {
	return M{
		"columns":               DefaultColumns,
		"ignore_unknown_fields": false,
		"options":               M{"lang": "en"},
		"range":                 []int{offset, limit},
		"sort":                  M{"sortBy": "market_cap_basic", "sortOrder": "desc"},
		"symbols":               M{},
		"markets":               []string{market},
		"filter": []M{
			{"left": "is_primary", "operation": "equal", "right": true},
		},
		"filter2": M{
			"operator": "and",
			"operands": []M{
				{
					"operation": M{
						"operator": "or",
						"operands": []M{
							{
								"operation": M{
									"operator": "and",
									"operands": []M{
										{"expression": M{"left": "type", "operation": "equal", "right": "stock"}},
										{"expression": M{"left": "typespecs", "operation": "has", "right": []string{"common"}}},
									},
								},
							},
							{
								"operation": M{
									"operator": "and",
									"operands": []M{
										{"expression": M{"left": "type", "operation": "equal", "right": "stock"}},
										{"expression": M{"left": "typespecs", "operation": "has", "right": []string{"preferred"}}},
									},
								},
							},
							{
								"operation": M{
									"operator": "and",
									"operands": []M{
										{"expression": M{"left": "type", "operation": "equal", "right": "dr"}},
									},
								},
							},
							{
								"operation": M{
									"operator": "and",
									"operands": []M{
										{"expression": M{"left": "type", "operation": "equal", "right": "fund"}},
										{"expression": M{"left": "typespecs", "operation": "has_none_of", "right": []string{"etf", "mutual", "closedend"}}},
									},
								},
							},
						},
					},
				},
				{
					"expression": M{
						"left":      "typespecs",
						"operation": "has_none_of",
						"right":     []string{"pre-ipo"},
					},
				},
			},
		},
	}
}
