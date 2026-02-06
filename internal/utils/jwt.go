package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims represents the JWT payload.
// We store user_id and standard JWT claims (exp, iat, etc.)
type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateToken creates a JWT token for a given user ID.
func GenerateToken(userID uint, secret string, expiresInMinutes int) (string, error) {

	// STEP 1:
	// Calculate expiration time (now + expiresInMinutes)
	expirationTime := jwt.NewNumericDate(time.Now().Add(time.Duration(expiresInMinutes) * time.Minute))

	// STEP 2:
	// Create claims (payload) with user_id and expiration time
	
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expirationTime,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "finance-tracker-api",
		},
	}


	// STEP 3:
	// Create a new JWT token with HS256 signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)


	// STEP 4:
	// Sign the token using the secret key
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	// STEP 5:
	// Return the signed token
	return signedToken, nil

}


// ValidateToken verifies the JWT token and returns the claims.
func ValidateToken(tokenString string, secret string) (*Claims, error) {

	// STEP 1:
	// Parse the token with claims
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})


	// STEP 2:
	// Check if token is valid


	// STEP 3:
	// Extract claims and return

}