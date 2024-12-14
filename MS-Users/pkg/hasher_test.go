package pkg

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "mysecretpassword"
	salt := "randomsalt"

	hashedPassword, err := HashPassword(password, salt)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if hashedPassword == "" {
		t.Fatalf("expected hashed password, got empty string")
	}

	err = ComparePasswords(hashedPassword, password, salt)
	if err != nil {
		t.Fatalf("expected passwords to match, got %v", err)
	}
}

func TestHashPasswordWithEmptyPassword(t *testing.T) {
	password := ""
	salt := "randomsalt"

	hashedPassword, err := HashPassword(password, salt)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if hashedPassword == "" {
		t.Fatalf("expected hashed password, got empty string")
	}

	err = ComparePasswords(hashedPassword, password, salt)
	if err != nil {
		t.Fatalf("expected passwords to match, got %v", err)
	}
}

func TestHashPasswordWithEmptySalt(t *testing.T) {
	password := "mysecretpassword"
	salt := ""

	hashedPassword, err := HashPassword(password, salt)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if hashedPassword == "" {
		t.Fatalf("expected hashed password, got empty string")
	}

	err = ComparePasswords(hashedPassword, password, salt)
	if err != nil {
		t.Fatalf("expected passwords to match, got %v", err)
	}
}
