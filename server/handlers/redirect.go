package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog/log"
)

// type RedirectHandle struct {
// }

func Redirect(w http.ResponseWriter, r *http.Request) {

	shortURL := chi.URLParam(r, "shortURL")
	fmt.Println("here")
	log.Debug().Msgf("got shortURL: %s", shortURL)
	// decode short URL
	// fetch from DB
	// return not found or original url
	// use go routines and threads
	longURL := "https://www.google.com"
	log.Debug().Msgf("redirecting to: %s", longURL)
	http.Redirect(w, r, longURL, http.StatusTemporaryRedirect)
}
