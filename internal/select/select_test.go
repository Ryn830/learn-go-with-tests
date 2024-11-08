package selectpkg

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("Returns url of faster API", func(t *testing.T) {
		slowServer := createDelayServer(20 * time.Millisecond)
		fastServer := createDelayServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		expected := fastServer.URL
		actual, _ := Racer(slowServer.URL, fastServer.URL)

		if actual != expected {
			t.Errorf("Actual: %q not equal %q", actual, expected)
		}
	})

	t.Run("Returns timeout error after 10 seconds", func(t *testing.T) {
		server := createDelayServer(20 * time.Millisecond)
		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 10*time.Millisecond)

		if err == nil {
			t.Error("Expected an error, but didn't get one.")
		}
	})
}

func createDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
