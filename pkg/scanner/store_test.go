// pkg/scanner/store_test.go
package scanner

import (
	"strings"
	"testing"
)

func TestUpdateAndGet(t *testing.T) {
	mockData := []TVAsset{
		{Ticker: "BBRI", Description: "Bank Rakyat Indonesia"},
		{Ticker: "TLKM", Description: "Telkom Indonesia"},
	}

	GlobalStore.UpdateData(mockData)

	assets := GlobalStore.GetData()
	if len(assets) != 2 {
		t.Errorf("Expected 2 assets, got %d", len(assets))
	}
	t.Run("Case Insensitivity Search", func(t *testing.T) {
		s := GlobalStore
		s.RLock()
		defer s.RUnlock()

		searchKey := strings.ToUpper("bbri")
		_, found := s.Index[searchKey]

		if !found {
			t.Errorf("Failed to find BBRI. Map has %d items", len(s.Index))
		}
	})
}
