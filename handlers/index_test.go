package handlers

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleIndex_Get(t *testing.T) {
	Templates = template.Must(template.New("index.html").Parse("OK"))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	HandleIndex(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rec.Code)
	}
}

func TestHandleIndex_WrongMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()

	HandleIndex(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected 405, got %d", rec.Code)
	}
}
