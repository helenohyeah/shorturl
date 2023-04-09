package utils

import (
	"errors"
	"net/url"
	"regexp"
)

const (
	HTTP  = "http"
	HTTPS = "https"
)

var (
	supportedSchemes = map[string]bool{
		HTTP:  true,
		HTTPS: true,
	}
)

// ParseURL will parse a string into valid url with a scheme / protocol
// Will return error if protocol is not https or http
func ParseURL(str string) (string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return "", err
	}
	// If we don't have a scheme, set it to https and then parse it again
	if u.Scheme == "" {
		u.Scheme = HTTPS
		u, err = url.Parse(u.String())
	}
	if err := validate(u); err != nil {
		return "", err
	}

	return u.String(), nil
}

func validate(u *url.URL) error {
	if u.Scheme != "" && !supportedSchemes[u.Scheme] {
		return errors.New("invalid protocol")
	}

	if u.Host == "" {
		return errors.New("missing host")
	}

	// Check for top level domain
	tldRegex, err := regexp.Compile("[^.]*\\.[^.]{2,3}(?:\\.[^.]{2,3})?$")
	if err != nil {
		return err
	}
	match := tldRegex.MatchString(u.Host)
	if !match {
		return errors.New("missing top level domain")
	}

	return nil
}
