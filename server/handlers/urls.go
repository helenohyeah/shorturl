package handlers

import (
	"errors"
	"net/http"

	"github.com/helen/lumen5_miniurl/db"
	"github.com/helen/lumen5_miniurl/models"
	"github.com/helen/lumen5_miniurl/utils"
	"github.com/rs/zerolog/log"
)

type URLHandle struct {
	DB *db.DB
}

var data map[int]models.URL

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
	url.EncodedURL = utils.ToBase62(url.ID)

	utils.WriteJSON(w, struct {
		Data models.URL `json:"data"`
	}{
		Data: url,
	})
}
