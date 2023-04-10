package handlers

import (
	"database/sql"
	"net/http"
	"net/url"

	"github.com/go-chi/chi"
	"github.com/helen/lumen5_miniurl/db"
	"github.com/helen/lumen5_miniurl/utils"
	"github.com/rs/zerolog/log"
)

type RedirectHandle struct {
	DB *db.DB
}

func (h *RedirectHandle) Redirect(w http.ResponseWriter, r *http.Request) {
	shortURL := chi.URLParam(r, "shortURL")

	log.Debug().Msgf("got shortURL: %s", shortURL)

	decodedID, err := utils.FromBase62(shortURL)
	if err != nil {
		log.Error().Err(err).Msgf("Redirect - failed to decode shortURL: %v", shortURL)
		utils.WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sqlStatement := `
		SELECT id, redirect_url FROM urls
		WHERE id = $1
	`
	var (
		id          string
		redirectURL string
	)
	if err = h.DB.QueryRow(sqlStatement, decodedID).Scan(&id, &redirectURL); err != nil {
		log.Error().Err(err).Msg("CreateShortURL - failed to query url")
		status := http.StatusInternalServerError
		if err == sql.ErrNoRows {
			status = http.StatusNotFound
		}
		utils.WriteJSONError(w, err.Error(), status)
		return
	}

	// todo:
	// use go routines and threads

	log.Debug().Msgf("redirecting to: %s\n", redirectURL)
	u, _ := url.Parse(redirectURL)
	log.Debug().Msgf("parse url: %+v\n", u)
	http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
}
