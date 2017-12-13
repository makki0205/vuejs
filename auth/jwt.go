package auth

import (
	"time"

	jwt_lib "github.com/dgrijalva/jwt-go"
)

var key = "test"
var exp = time.Hour * 100

func createToken(userID uint) string {
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
	token.Claims = jwt_lib.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(exp).Unix(),
	}
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		panic(err)
	}
	return tokenString
}

func perseToken(token string) (uint, error) {
	t, err := jwt_lib.Parse(token, func(t *jwt_lib.Token) (interface{}, error) { return []byte(key), nil })
	if err != nil {
		return 0, err
	}
	id := uint(t.Claims.(jwt_lib.MapClaims)["id"].(float64))
	return id, nil
}
