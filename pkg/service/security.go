package service

import (
	"fmt"

	"github.com/darksasori/finance/pkg/utils"
	jwt "github.com/dgrijalva/jwt-go"
)

func encode(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})

	return token.SignedString([]byte(utils.GetEnv("SECRET", "finance")))
}

func decode(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected method")
		}
		return []byte(utils.GetEnv("SECRET", "finance")), nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", fmt.Errorf("Token invalid")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if username, ok := claims["username"].(string); ok {
			return username, nil
		}
		return "", fmt.Errorf("Token invalid")
	}

	return "", fmt.Errorf("Token invalid")
}
