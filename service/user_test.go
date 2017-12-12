package service

import (
	"fmt"
	"testing"
)

func TestUserImpl_FindByEmail(t *testing.T) {
	user := User.FindByEmail("llxo2_5oxaall@icloud.com")
	fmt.Println(user)
}
