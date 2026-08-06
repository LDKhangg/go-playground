//go:build exercise

package jwtx

import (
	"encoding/base64"
	"errors"
	"strings"
	"testing"
	"time"
)

var testSecret = []byte("super-secret")

func TestSignVerifyRoundTrip(t *testing.T) {
	claims := map[string]any{
		"sub": "user-1",
		"exp": time.Now().Add(time.Hour).Unix(),
	}

	token, err := Sign(testSecret, claims)
	if err != nil {
		t.Fatalf("Sign: %v", err)
	}
	if parts := strings.Split(token, "."); len(parts) != 3 {
		t.Fatalf("token has %d parts, want 3", len(parts))
	}

	got, err := Verify(testSecret, token)
	if err != nil {
		t.Fatalf("Verify: %v", err)
	}
	if got["sub"] != "user-1" {
		t.Fatalf("sub = %v, want user-1", got["sub"])
	}
}

func TestVerifyRejectsWrongSecret(t *testing.T) {
	token, err := Sign(testSecret, map[string]any{"sub": "user-1"})
	if err != nil {
		t.Fatalf("Sign: %v", err)
	}

	if _, err := Verify([]byte("wrong-secret"), token); err == nil {
		t.Fatal("expected an error for the wrong secret")
	}
}

func TestVerifyRejectsTamperedPayload(t *testing.T) {
	token, err := Sign(testSecret, map[string]any{"sub": "user-1"})
	if err != nil {
		t.Fatalf("Sign: %v", err)
	}

	parts := strings.Split(token, ".")
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		t.Fatalf("decode payload: %v", err)
	}
	tamperedPayload := base64.RawURLEncoding.EncodeToString(
		[]byte(strings.ReplaceAll(string(payload), "user-1", "user-2")))
	tampered := parts[0] + "." + tamperedPayload + "." + parts[2]

	if _, err := Verify(testSecret, tampered); err == nil {
		t.Fatal("expected an error for a tampered payload")
	}
}

func TestVerifyRejectsExpired(t *testing.T) {
	claims := map[string]any{
		"sub": "user-1",
		"exp": time.Now().Add(-time.Hour).Unix(),
	}
	token, err := Sign(testSecret, claims)
	if err != nil {
		t.Fatalf("Sign: %v", err)
	}

	if _, err := Verify(testSecret, token); !errors.Is(err, ErrExpired) {
		t.Fatalf("err = %v, want ErrExpired", err)
	}
}