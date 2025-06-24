package urlshortner

import (
	"testing"
)

func TestGenerateShortCode_Length(t *testing.T) {
	code := GenerateShortCode(8)
	if len(code) != 8 {
		t.Errorf("expected length 8, got %d", len(code))
	}
}

func TestGenerateShortCode_Uniqueness(t *testing.T) {
	codes := make(map[string]struct{})
	for i := 0; i < 1000; i++ {
		code := GenerateShortCode(6)
		if _, exists := codes[code]; exists {
			t.Errorf("duplicate code generated: %s", code)
		}
		codes[code] = struct{}{}
	}
}
