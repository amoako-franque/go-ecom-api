package auth

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	hash, err := HashUserPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if hash == "" {
		t.Error("expected hash to be not empty")
	}

	if hash == "password" {
		t.Error("expected hash to be different from password")
	}
}

func TestComparePasswords(t *testing.T) {
	hash, err := HashUserPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if !CompareUserPassword(hash, []byte("password")) {
		t.Errorf("expected password to match hash")
	}
	if CompareUserPassword(hash, []byte("not password")) {
		t.Errorf("expected password to not match hash")
	}
}
