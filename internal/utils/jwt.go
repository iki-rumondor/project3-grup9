package utils

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = "project3Secret"

func GenerateToken(data map[string]interface{}) (string, error) {
	claims := jwt.MapClaims(data)
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := parseToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(strToken string) (jwt.MapClaims, error) {
	errResponse := errors.New("please ensure you have the right credentials to proceed")

	token, _ := jwt.Parse(strToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(secretKey), nil
	})

	mapClaims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, errResponse
	}

	return mapClaims, nil
}
