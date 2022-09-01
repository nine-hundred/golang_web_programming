package membership_test

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"golang_web_programming/app/membership"
	"testing"
)

func TestRouterFunc(t *testing.T) {
	t.Run("valid jwt", func(t *testing.T) {
		tmpClaims := jwt.MapClaims{
			"id": "id",
			"pw": "pw",
		}
		testTokenStr, _ := membership.GenJwt(tmpClaims)
		fmt.Println(testTokenStr)
		token, err := jwt.Parse(testTokenStr, membership.ValidJwt)
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			t.Log(claims)
		} else {
			t.Log(err)
		}
		assert.Nil(t, err)
	})
}
