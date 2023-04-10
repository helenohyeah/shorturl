package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/helen/lumen5_miniurl/db"
	"github.com/helen/lumen5_miniurl/models"
	"github.com/helen/lumen5_miniurl/utils"
	"github.com/rs/zerolog/log"
)

var (
	errInvalidUsernamePassword = errors.New("incorrect username or password")
)

type UserAccountHandle struct {
	DB *db.DB
}

func (h *UserAccountHandle) Register(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var req request
	if err := utils.ReadJSON(r, &req); err != nil {
		log.Error().Err(err).Msg("Register - failed to parse json request")
		utils.WriteJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Username == "" || req.Password == "" {
		log.Error().Msg("Register - received empty username/password")
		utils.WriteJSONError(w, errInvalidUsernamePassword.Error(), http.StatusBadRequest)
		return
	}

	sqlStatement := `
		SELECT EXISTS (SELECT 1 FROM users WHERE username = $1)
	`
	exists := false
	if err := h.DB.QueryRow(sqlStatement, req.Username).Scan(&exists); err != nil {
		log.Error().Err(err).Msg("Register - failed to query user exists")
		utils.WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if exists {
		log.Error().Msg("Register - username exists")
		utils.WriteJSONError(w, errInvalidUsernamePassword.Error(), http.StatusBadRequest)
		return
	}

	sqlStatement = `
		INSERT INTO users (
			username
			,password
		) VALUES ($1, $2)
		RETURNING id
	`
	var id int
	if err := h.DB.QueryRow(sqlStatement, req.Username, req.Password).Scan(&id); err != nil {
		log.Error().Err(err).Msg("Register - failed to write user to DB")
		utils.WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// todo:
	// hash password
	// return session id or something for a logged in session

	utils.WriteJSON(w, struct {
		Data models.User `json:"data"`
	}{
		Data: models.User{
			ID:       id,
			Username: req.Username,
			Password: req.Password,
		},
	})
}

func (h *UserAccountHandle) Login(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var req request
	if err := utils.ReadJSON(r, &req); err != nil {
		log.Error().Err(err).Msg("Login - failed to parse json request")
		utils.WriteJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Username == "" || req.Password == "" {
		log.Error().Msg("Login - received empty username/password")
		utils.WriteJSONError(w, errInvalidUsernamePassword.Error(), http.StatusBadRequest)
		return
	}

	sqlStatement := `
		SELECT id, password FROM users WHERE username = $1
	`
	var (
		id       int
		password string
	)
	if err := h.DB.QueryRow(sqlStatement, req.Username).Scan(&id, &password); err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, sql.ErrNoRows) {
			status = http.StatusUnauthorized
		}
		log.Error().Err(err).Msg("Login - failed to look up user")
		utils.WriteJSONError(w, errInvalidUsernamePassword.Error(), status)
		return
	}

	if strings.TrimSpace(password) != req.Password {
		log.Error().Msg("Login - password did not match record")
		utils.WriteJSONError(w, errInvalidUsernamePassword.Error(), http.StatusUnauthorized)
		return
	}

	// todo:
	// return session id or something for a logged in session

	utils.WriteJSON(w, struct {
		Data models.User `json:"data"`
	}{
		Data: models.User{
			ID:       id,
			Username: req.Username,
			Password: req.Password,
		},
	})
}
