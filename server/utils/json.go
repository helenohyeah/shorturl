package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog/log"
)

const (
	requestBodyLimit = 1024 * 1024 // 1 MB
)

type JSONResponse struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

// ReadJSON decodes and unmarshals the content from the HTTP request body.
func ReadJSON(r *http.Request, v interface{}) error {
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		return errors.New("Content-Type is not application/json")
	}

	b := bytes.NewBuffer(make([]byte, 0))
	reader := io.TeeReader(r.Body, b)
	r.Body = ioutil.NopCloser(b)

	lr := io.LimitReader(reader, requestBodyLimit)
	if err := json.NewDecoder(lr).Decode(v); err != nil {
		log.Error().Err(err).Str("request body", string(b.Bytes())).Msg("failed to decode json")
		return err
	}

	return nil
}

// WriteJSONError returns a json object with error back to the client
func WriteJSONError(w http.ResponseWriter, errMsg string, status int) {
	resp := JSONResponse{
		Error: errMsg,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(resp)

	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
}

// WriteJSON encodes the provided value and returns it as json back the the client
func WriteJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	result, _ := json.Marshal(v)
	fmt.Println(string(result))
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Error().Err(err).Interface("response", v).Msg("failed to encode json")
	}

	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
}
