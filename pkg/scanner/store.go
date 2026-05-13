// pkg/scanner/store.go

package scanner

import (
	"sync"
	"time"
)

type StockStore struct {
	sync.RWMutex
	Assets     []TVAsset
	LastUpdate time.Time
}

var GlobalStore = &StockStore{
	Assets: []TVAsset{},
}

func (s *StockStore) UpdateData(newData []TVAsset) {
	s.Lock()
	defer s.Unlock()
	s.Assets = newData
	s.LastUpdate = time.Now()
}

func (s *StockStore) GetData() []TVAsset {
	s.RLock()
	defer s.RUnlock()
	return s.Assets
}
