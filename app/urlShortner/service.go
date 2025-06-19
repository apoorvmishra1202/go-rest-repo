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

func generateShortCode() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
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
	short := generateShortCode()
	for {
		if _, exists := urlStore[short]; !exists {
			break
		}
		short = generateShortCode()
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
