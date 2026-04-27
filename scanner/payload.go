package scanner

type M map[string]interface{}

var DefaultColumns = []string{
	"name",
	"close",
	"market_cap_basic",
	"net_income",
	"price_earnings_ttm",
	"earnings_per_share_diluted_yoy_growth_ttm",
	"dividends_yield_current",
	"fundamental_currency_code",
}

func BuildIDXPayload(offset, limit int) M {
	return M{
		"columns":               DefaultColumns,
		"ignore_unknown_fields": false,
		"options":               M{"lang": "en"},
		"range":                 []int{offset, limit},
		"sort":                  M{"sortBy": "market_cap_basic", "sortOrder": "desc"},
		"symbols":               M{},
		"markets":               []string{"indonesia"},
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
