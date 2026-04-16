package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// This is a single key JWT creation method
type JWTService struct {
	secretKey []byte
	issuer    string
}

func NewJwtService(secret, appName string) *JWTService {
	return &JWTService{
		secretKey: []byte(secret),
		issuer:    appName,
	}
}

// Custom claims are like meta data that we pass in to jwt to generate token
type CustomClaims struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

// Generating JWT tokens
func (s *JWTService) GenerateToken(userId string) (string, error) {
	// Creating a custom claims instance
	claims := CustomClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			// Issuer will hold the app name
			Issuer: s.issuer,
			// ExpiresAt will hold the JWT expiration time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			// IssuedAt will hold the time at which the token was generated
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	// This will return us a token that is signed using ES256 method
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	// Here, we are signing the string with the secret key
	return token.SignedString(s.secretKey)
}

// Validate Token
func (s *JWTService) ValidateToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return s.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
