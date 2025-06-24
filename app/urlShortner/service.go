package urlshortner

import (
	"math/rand"
	"sync"
	"time"
)

var (
	urlStore = make(map[string]string) // short -> original
	storeMu  sync.RWMutex
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func SetShortURL(original string) string {
	storeMu.Lock()
	defer storeMu.Unlock()
	// Check if already exists
	for k, v := range urlStore {
		if v == original {
			return k
		}
	}
	short := GenerateShortCode(6)
	for {
		if _, exists := urlStore[short]; !exists {
			break
		}
		short = GenerateShortCode(6)
	}
	urlStore[short] = original
	return short
}

func GetOriginalURL(short string) (string, bool) {
	storeMu.RLock()
	defer storeMu.RUnlock()
	orig, ok := urlStore[short]
	return orig, ok
}
