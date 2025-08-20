package auth

import (
	"errors"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(getSecret())

func getSecret() string {
	if s := os.Getenv("JWT_SECRET"); s != "" {
		return s
	}
	return "dev_secret_change_me"
}

func GenerateJWT(userID string) (string, error) {
	claims := jwt.MapClaims{}
	claims["sub"] = userID
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenStr string) (string, error) {
	t, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		if sub, ok := claims["sub"].(string); ok {
			return sub, nil
		}
	}
	return "", errors.New("invalid token claims")
}
