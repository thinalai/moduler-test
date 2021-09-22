package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		player := &PlayerServer{}
		player.ServeHTTP(response, request)

		got := response.Body.String()
		want := "hello wiredcraft"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
