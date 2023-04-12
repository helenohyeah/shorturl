package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/helen/lumen5_miniurl/db"
	"github.com/helen/lumen5_miniurl/models"
	"github.com/helen/lumen5_miniurl/utils"
	"github.com/rs/zerolog/log"
)

type URLHandle struct {
	DB *db.DB
}

func (h *URLHandle) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	type request struct {
		OriginalURL string `json:"originalUrl"`
		UserID      int64  `json:"userId"`
	}
	var req request
	if err := utils.ReadJSON(r, &req); err != nil {
		log.Error().Err(err).Msg("CreateShortURL - failed to parse json request")
		utils.WriteJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.OriginalURL == "" {
		err := errors.New("url cannot be empty")
		log.Error().Err(err).Msg("CreateShortURL - received empty original url")
		utils.WriteJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	longURL, err := utils.ParseURL(req.OriginalURL)
	if err != nil {
		log.Error().Err(err).Msg("CreateShortURL - error calling ParseURL")
		utils.WriteJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// todo: validate userID
	userID := models.NewNullInt64(req.UserID, req.UserID > 0)

	sqlStatement := `
		INSERT INTO urls (redirect_url, user_id)
		VALUES ($1, $2)
		RETURNING id, redirect_url, user_id
	`

	url := models.URL{}
	if err = h.DB.QueryRow(sqlStatement, longURL, userID).Scan(&url.ID, &url.RedirectURL, &url.UserID); err != nil {
		log.Error().Err(err).Msg("CreateShortURL - failed to write to DB")
		utils.WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	url.EncodedURL = utils.EncodeURL(url.ID)

	utils.WriteJSON(w, struct {
		Data models.URL `json:"data"`
	}{
		Data: url,
	})
}

func (h *URLHandle) GetURLsByUserID(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	// todo: authenticate

	u, err := strconv.Atoi(userID)
	if err != nil {
		log.Error().Err(err).Msgf("GetURLsByUserID - failed to convert userID %s to int", userID)
		utils.WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	urls, err := h.urlsByUserID(uint64(u))
	if err != nil {
		log.Error().Err(err).Msgf("GetURLsByUserID - failed to get urls for userID %s", userID)
		utils.WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, url := range urls {
		url.EncodedURL = utils.EncodeURL(url.ID)
	}

	utils.WriteJSON(w, struct {
		Data []*models.URL `json:"data,omitempty"`
	}{
		Data: urls,
	})
}

func (h *URLHandle) urlsByUserID(userID uint64) ([]*models.URL, error) {
	rows, err := h.DB.Query(`
		SELECT
			id
			,user_id
			,redirect_url
		FROM urls WHERE user_id = $1
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// hold data from returned rows.
	var urls []*models.URL

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var url models.URL
		if err := rows.Scan(&url.ID, &url.UserID, &url.RedirectURL); err != nil {
			return urls, err
		}
		urls = append(urls, &url)
	}
	if err = rows.Err(); err != nil {
		return urls, err
	}
	return urls, nil
}
