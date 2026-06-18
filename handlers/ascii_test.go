package handlers

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAsciiArt_Post_OK(t *testing.T) {
	templates = template.Must(template.New("index.html").Parse("OK"))

	body := strings.NewReader("text=A&banner=standard")

	req := httptest.NewRequest(http.MethodPost, "/ascii-art", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rec := httptest.NewRecorder()

	HandleAsciiArt(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rec.Code)
	}
}
