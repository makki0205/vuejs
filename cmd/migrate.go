package main

import "github.com/makki0205/vuejs/model"

var db = model.GetDBConn()

func main() {
	db.DropTableIfExists(&model.User{})

	db.AutoMigrate(&model.User{})
}
