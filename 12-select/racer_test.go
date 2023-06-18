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
	
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	
	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusOK)
	}))
	
	defer slowServer.Close()
	defer fastServer.Close()
	
	slowUrl := slowServer.URL
	fastUrl := fastServer.URL

	want := fastUrl

	got := Racer(slowUrl, fastUrl)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}