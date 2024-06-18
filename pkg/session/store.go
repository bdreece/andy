package session

import (
	"encoding/base64"

	"github.com/gorilla/sessions"
)

func NewCookieStore(cfg *Config) (*sessions.CookieStore, error) {
	signingKey, err := base64.StdEncoding.DecodeString(cfg.SigningKey)
	if err != nil {
		return nil, err
	}

	encryptingKey, err := base64.StdEncoding.DecodeString(cfg.EncryptingKey)
	if err != nil {
		return nil, err
	}

	return sessions.NewCookieStore(signingKey, encryptingKey), nil
}
