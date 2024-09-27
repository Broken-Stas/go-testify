package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Fatalf("expected status code: %d, got %d", http.StatusOK, status)
	}

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	assert.Equal(t, len(list), totalCount, "expected cafe count: %d, got %d", totalCount, len(list))
}
func TestMainHandlerWhenRequestIsCorrect(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe", nil)
	q := req.URL.Query()
	q.Add("city", "moscow")
	q.Add("count", "4")

	req.URL.RawQuery = q.Encode()

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	expectedCode := http.StatusOK
	if responseRecorder.Code != expectedCode {
		assert.Equal(t, expectedCode, responseRecorder.Code)
	}

	body := responseRecorder.Body.String()
	assert.NotEmpty(t, body)
}
func TestMainHandleWhenCityIsNotSupported(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe", nil)
	q := req.URL.Query()
	q.Add("city", "wrong_city") // Неподдерживаемый город
	q.Add("count", "1")         // Пример значения count

	req.URL.RawQuery = q.Encode()

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	expectedCode := http.StatusBadRequest
	if responseRecorder.Code != expectedCode {
		assert.Equal(t, expectedCode, responseRecorder.Code)
	}

	body := responseRecorder.Body.String()
	assert.Equal(t, body, "wrong city value")
}
