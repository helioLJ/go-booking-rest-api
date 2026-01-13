package utils

import (
	"testing"
)

func TestGenerateToken(t *testing.T) {
	email := "test@example.com"
	userId := int64(1)

	token, err := GenerateToken(email, userId)

	if err != nil {
		t.Errorf("GenerateToken failed: %v", err)
	}

	if token == "" {
		t.Error("GenerateToken returned empty token")
	}
}

func TestVerifyToken(t *testing.T) {
	email := "test@example.com"
	userId := int64(123)

	// Generate a valid token
	token, err := GenerateToken(email, userId)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	// Verify the token
	verifiedUserId, err := VerifyToken(token)

	if err != nil {
		t.Errorf("VerifyToken failed: %v", err)
	}

	if verifiedUserId != userId {
		t.Errorf("Expected userId %d, got %d", userId, verifiedUserId)
	}
}

func TestVerifyInvalidToken(t *testing.T) {
	invalidToken := "invalid.token.here"

	_, err := VerifyToken(invalidToken)

	if err == nil {
		t.Error("VerifyToken should fail with invalid token")
	}
}

func TestVerifyEmptyToken(t *testing.T) {
	_, err := VerifyToken("")

	if err == nil {
		t.Error("VerifyToken should fail with empty token")
	}
}

func TestTokenExpiration(t *testing.T) {
	// This test verifies that tokens have an expiration time
	email := "test@example.com"
	userId := int64(1)

	token, err := GenerateToken(email, userId)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	// Verify token is currently valid
	verifiedUserId, err := VerifyToken(token)
	if err != nil {
		t.Errorf("Token should be valid: %v", err)
	}

	if verifiedUserId != userId {
		t.Errorf("Expected userId %d, got %d", userId, verifiedUserId)
	}

	// Note: We can't easily test actual expiration without waiting 2 hours
	// or mocking time. This test just ensures the token works initially.
}

func TestMultipleTokenGeneration(t *testing.T) {
	email1 := "user1@example.com"
	email2 := "user2@example.com"
	userId1 := int64(1)
	userId2 := int64(2)

	token1, err1 := GenerateToken(email1, userId1)
	token2, err2 := GenerateToken(email2, userId2)

	if err1 != nil || err2 != nil {
		t.Fatal("Failed to generate tokens")
	}

	// Tokens should be different
	if token1 == token2 {
		t.Error("Different users produced identical tokens")
	}

	// Each token should verify to correct user
	verifiedId1, err1 := VerifyToken(token1)
	verifiedId2, err2 := VerifyToken(token2)

	if err1 != nil || err2 != nil {
		t.Fatal("Failed to verify tokens")
	}

	if verifiedId1 != userId1 {
		t.Errorf("Token1: expected userId %d, got %d", userId1, verifiedId1)
	}

	if verifiedId2 != userId2 {
		t.Errorf("Token2: expected userId %d, got %d", userId2, verifiedId2)
	}
}
