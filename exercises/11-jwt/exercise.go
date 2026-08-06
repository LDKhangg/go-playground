package jwtx

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"time"
)

var (
	ErrMalformed = errors.New("malformed token")
	ErrExpired   = errors.New("token expired")
)

func nowUnix() int64 { return time.Now().Unix() }

func base64url(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}

func encodeJSON(value any) (string, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return base64url(data), nil
}

func signInput(secret []byte, data string) []byte {
	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(data))
	return mac.Sum(nil)
}

func Sign(secret []byte, claims map[string]any) (string, error) {
	headerPart, err := encodeJSON(map[string]any{"alg": "HS256", "typ": "JWT"})
	if err != nil {
		return "", err
	}
	payloadPart, err := encodeJSON(claims)
	if err != nil {
		return "", err
	}
	sig := base64url(signInput(secret, headerPart+"."+payloadPart))
	return headerPart + "." + payloadPart + "." + sig, nil
}

func Verify(secret []byte, token string) (map[string]any, error) {
	panic("TODO")
}