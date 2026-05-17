// pkg/scanner/store.go

package scanner

import (
	"strings"
	"sync"
	"time"
)

type StockStore struct {
	sync.RWMutex
	Assets     []TVAsset
	Index      map[string]TVAsset
	LastUpdate time.Time
}

var GlobalStore = &StockStore{
	Assets: []TVAsset{},
	Index:  make(map[string]TVAsset),
}

func (s *StockStore) UpdateData(newData []TVAsset) {
	s.Lock()
	defer s.Unlock()
	s.Assets = newData

	newId := make(map[string]TVAsset)
	for _, data := range newData {
		newId[strings.ToUpper(data.Ticker)] = data
	}
	s.Index = newId
	s.LastUpdate = time.Now()
}

func (s *StockStore) GetData() []TVAsset {
	s.RLock()
	defer s.RUnlock()
	return s.Assets
}
