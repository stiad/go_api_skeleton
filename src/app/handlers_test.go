package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

/*****************
	 Tests
 *****************/

func TestServer_HandlerHelloWorld(T *testing.T) {
	server := setup()

	tt := []struct {
		expectedStatusCode int
	}{
		{http.StatusOK},
	}

	for _, t := range tt {
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/hello", nil)
		server.ServeHTTP(rec, req)

		if t.expectedStatusCode != rec.Code {
			T.Errorf("expected a status code of %d, but got %d", t.expectedStatusCode, rec.Code)
		}
	}
}

func setup() *Server {
	server := NewServer()
	server.routes()
	server.testing = true
	return server
}
