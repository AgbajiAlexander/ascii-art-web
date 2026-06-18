package handlers

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestHandleAsciiArt_MethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/ascii-art", nil)
	rec := httptest.NewRecorder()

	HandleAsciiArt(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected 405, got %d", rec.Code)
	}
}

func TestHandleAsciiArt_EmptyText(t *testing.T) {
	Templates = template.Must(template.New("index.html").Parse("{{.Error}}"))

	form := url.Values{}
	form.Add("text", "")
	form.Add("banner", "standard")

	req := httptest.NewRequest(
		http.MethodPost,
		"/ascii-art",
		strings.NewReader(form.Encode()),
	)

	req.Header.Set(
		"Content-Type",
		"application/x-www-form-urlencoded",
	)

	rec := httptest.NewRecorder()

	HandleAsciiArt(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rec.Code)
	}
}
