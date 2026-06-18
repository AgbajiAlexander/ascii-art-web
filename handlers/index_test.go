package handlers

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"
)

// helper: safe template init for tests
func setupIndexTest() {
	tmpl := template.Must(template.New("index.html").Parse("OK"))
	InitTemplates(tmpl)
}

func TestHandleIndex_OK(t *testing.T) {
	setupIndexTest()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	HandleIndex(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rec.Code)
	}
}

func TestHandleIndex_NotFound(t *testing.T) {
	setupIndexTest()

	req := httptest.NewRequest(http.MethodGet, "/wrong", nil)
	rec := httptest.NewRecorder()

	HandleIndex(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", rec.Code)
	}
}

func TestHandleIndex_MethodNotAllowed(t *testing.T) {
	setupIndexTest()

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()

	HandleIndex(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected 405, got %d", rec.Code)
	}
}
