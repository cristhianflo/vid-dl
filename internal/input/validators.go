package input

import (
	"errors"
	"net/url"
	"strings"
)

func IsEmpty(str string) error {
	trimmed := strings.TrimSpace(str)
	if trimmed == "" {
		return errors.New("Input is empty.")
	}

	return nil
}

func IsEmail(str string) (*url.URL, error) {
	u, err := url.Parse(str)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return nil, errors.New("Input is not a valid URL.")
	}
	return u, nil
}
