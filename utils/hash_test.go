package utils

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "testPassword123"

	hashedPassword, err := HashPassword(password)

	if err != nil {
		t.Errorf("HashPassword failed: %v", err)
	}

	if hashedPassword == "" {
		t.Error("HashPassword returned empty string")
	}

	if hashedPassword == password {
		t.Error("HashPassword returned the same password (not hashed)")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "testPassword123"
	wrongPassword := "wrongPassword"

	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	// Test with correct password
	if !CheckPasswordHash(password, hashedPassword) {
		t.Error("CheckPasswordHash failed with correct password")
	}

	// Test with wrong password
	if CheckPasswordHash(wrongPassword, hashedPassword) {
		t.Error("CheckPasswordHash succeeded with wrong password")
	}
}

func TestHashPasswordConsistency(t *testing.T) {
	password := "testPassword123"

	hash1, err1 := HashPassword(password)
	hash2, err2 := HashPassword(password)

	if err1 != nil || err2 != nil {
		t.Fatal("Failed to hash password")
	}

	// Hashes should be different (bcrypt uses salt)
	if hash1 == hash2 {
		t.Error("Same password produced identical hashes (salt not working)")
	}

	// But both should validate the original password
	if !CheckPasswordHash(password, hash1) || !CheckPasswordHash(password, hash2) {
		t.Error("Hashes don't validate the original password")
	}
}
