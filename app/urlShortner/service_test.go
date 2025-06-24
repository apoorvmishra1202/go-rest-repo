package urlshortner

import (
	"testing"
)

func TestSetAndGetShortURL(t *testing.T) {
	original := "https://example.com"
	short := SetShortURL(original)
	if short == "" {
		t.Fatal("expected a short code, got empty string")
	}
	got, ok := GetOriginalURL(short)
	if !ok {
		t.Fatalf("expected to find original URL for short code %s", short)
	}
	if got != original {
		t.Errorf("expected %s, got %s", original, got)
	}
}

func TestSetShortURL_Duplicate(t *testing.T) {
	original := "https://duplicate.com"
	short1 := SetShortURL(original)
	short2 := SetShortURL(original)
	if short1 != short2 {
		t.Errorf("expected same short code for duplicate URLs, got %s and %s", short1, short2)
	}
}

func TestGetOriginalURL_NotFound(t *testing.T) {
	_, ok := GetOriginalURL("nonexistent")
	if ok {
		t.Error("expected not to find original URL for nonexistent short code")
	}
}
