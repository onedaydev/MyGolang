package middleware

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/onedaydev/mygolang/complex-server/config"
	"github.com/onedaydev/mygolang/complex-server/handlers"
)

func TestPanicMiddleware(t *testing.T) {
	b := new(bytes.Buffer)
	c := config.InitConfig(b)

	m := http.NewServeMux()
	handlers.Register(m, c)

	h := panicMiddleware(m, c)

	r := httptest.NewRequest("GET", "/panic", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)

	resp := w.Result()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error reading response body: %v", err)
	}

	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected response status: %v, Got: %v",
			http.StatusOK,
			resp.StatusCode,
		)
	}

	expectedResponseBody := "Unexpected server error occurred"

	if string(body) != expectedResponseBody {
		t.Errorf(
			"expected response: %s, Got: %s\n",
			expectedResponseBody,
			string(body),
		)
	}
}
