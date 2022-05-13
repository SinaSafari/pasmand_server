package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("no_one_can_find_this_secret_ha_ha_ha")

type Claims struct {
	Phone string `json:"phone"`
	jwt.StandardClaims
}

func parse(token string) (*jwt.Token, error) {

	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return jwtKey, nil
	})
}

func GenerateJwtToken(phone string, time time.Time) (string, error) {
	tokenClaims := &Claims{
		Phone: phone,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: time.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Verify(token string) (interface{}, error) {
	parsed, err := parse(token)
	if err != nil {
		return "", err
	}
	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return "", err
	}
	return claims["phone"], nil
}
