package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/helen/lumen5_miniurl/models"
	"github.com/helen/lumen5_miniurl/utils"
	"github.com/rs/zerolog/log"
)

// type URLHandle struct {
// }

var data map[int]models.URL

func CreateShortURL(w http.ResponseWriter, r *http.Request) {
	type request struct {
		OriginalURL string `json:"originalUrl"`
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

	fmt.Println(longURL)
	// generate a 7 alphanumeric short url
	// check if it exists
	// create in DB

	utils.WriteJSON(w, struct {
		Data string `json:"data"`
	}{
		Data: longURL,
	})
}
