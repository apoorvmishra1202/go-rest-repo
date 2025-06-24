package network

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSendSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	SendSuccess(c, gin.H{"foo": "bar"})
	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	if w.Body.String() != "{\"foo\":\"bar\"}\n" {
		t.Errorf("unexpected body: %s", w.Body.String())
	}
}

func TestSendFailure(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	SendFailure(c, http.StatusBadRequest, "fail")
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
	if w.Body.String() != "{\"error\":\"fail\"}\n" {
		t.Errorf("unexpected body: %s", w.Body.String())
	}
}
