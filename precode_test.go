package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusOK, "http status invalid")

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	assert.Len(t, list, totalCount)
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
	require.Equal(t, expectedCode, responseRecorder.Code)

	assert.NotEmpty(t, responseRecorder.Body)
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
	require.Equal(t, expectedCode, responseRecorder.Code)

	body := responseRecorder.Body.String()
	assert.Equal(t, body, "wrong city value")
} // СПАСИБО ЗА ВАШУ РАБОТУ, НАДЕЮСЬ Я ПРАВИЛЬНО ВАС ПОНЯЛ)
