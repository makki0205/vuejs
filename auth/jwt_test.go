package auth

import (
	"time"

	"testing"

	"fmt"

	jwt_lib "github.com/dgrijalva/jwt-go"
)

func TestHoge(t *testing.T) {
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
	token.Claims = jwt_lib.MapClaims{
		"id":  int(2),
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}
	tokenString, err := token.SignedString([]byte("hoge"))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(tokenString)
	token, err = jwt_lib.Parse(tokenString, func(t *jwt_lib.Token) (interface{}, error) {
		return []byte("hoge"), nil
	})
	if err != nil {
		panic(err)
	}
	id := token.Claims.(jwt_lib.MapClaims)["id"]
	fmt.Println(id)
}
