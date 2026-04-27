package scanner

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

/*
"price_earnings_ttm",
	"earnings_per_share_diluted_yoy_growth_ttm",
	"dividends_yield_current",
	"fundamental_currency_code",
*/

type TVAsset struct {
	Ticker    string
	Price     float64
	MarketCap float64
	// PBV       float64
	NetIncome               float64
	PERatio                 float64
	EPS                     float64
	CurrentDividenYield     float64
	FundamentalCurrencyCode string
}

type TVPayload struct {
	Filter  []map[string]interface{} `json:"filter"`
	Options map[string]interface{}   `json:"options"`
	Markets []string                 `json:"markets"`
	Symbols map[string]interface{}   `json:"symbols"`
	Columns []string                 `json:"columns"`
	Sort    map[string]string        `json:"sort"`
	Range   []int                    `json:"range"`
}

func FetchIDXData() ([]TVAsset, error) {
	url := "https://scanner.tradingview.com/indonesia/scan"

	payload := BuildIDXPayload(0, 100)
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal json: %v", err)
	}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))

	// 2. HEADER DILENGKAPIN
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	req.Header.Set("Referer", "https://www.tradingview.com/")
	req.Header.Set("Origin", "https://www.tradingview.com")

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 3. BACA ERRORNYA KALAU DITOLAK
	if resp.StatusCode != 200 {
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
		return nil, err
	}

	var assets []TVAsset
	for _, item := range result.Data {
		if len(item.D) < 7 {
			continue
		}

		assets = append(assets, TVAsset{
			Ticker:    fmt.Sprintf("%v", item.D[0]),
			Price:     convertToFloat(item.D[1]),
			MarketCap: convertToFloat(item.D[2]),
			// PBV:       convertToFloat(item.D[3]),
			NetIncome:               convertToFloat(item.D[3]),
			PERatio:                 convertToFloat(item.D[4]),
			EPS:                     convertToFloat(item.D[5]),
			CurrentDividenYield:     convertToFloat(item.D[6]),
			FundamentalCurrencyCode: fmt.Sprintf("%v", item.D[7]),
		})
	}

	return assets, nil
}

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
