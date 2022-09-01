package membership

import (
	"errors"
	"github.com/golang-jwt/jwt"
)

var key = []byte("signed")

func GenJwt(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return "", errors.New("fail to signed jwt")
	}
	return tokenStr, nil
}

func ParseJwt(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, ValidJwt)
	if err != nil {
		return nil, errors.New("fail to valid")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); !ok {
		return nil, errors.New("fail to parse token")
	} else {
		return claims, nil
	}
}

func ValidJwt(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("unexpected signing method")
	}
	return key, nil
}
