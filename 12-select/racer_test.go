package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	// slowUrl := "https://www.facebook.com"
	// fastUrl := "https://blog.theswarnim.tech"

    // want := fastUrl
	// got := Racer(slowUrl, fastUrl)

	// if got != want {
	// 	t.Errorf("got %q, want %q", got, want)
	// }
	
	t.Run("returns the url of the fastest server", func(t *testing.T){
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl

		got, _ := Racer(slowUrl, fastUrl)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T){
		server := makeDelayedServer(25 * time.Second)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20 * time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}