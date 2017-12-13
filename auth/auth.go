package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/makki0205/vuejs/controller"
	"github.com/makki0205/vuejs/model"
	"github.com/makki0205/vuejs/service"
)

func LoginController(c *gin.Context) {
	var req model.User
	err := c.BindJSON(&req)

	if err != nil {
		controller.BatRequest(err.Error(), c)
		return
	}
	user := service.User.FindByEmail(req.Email)
	if user == nil {
		controller.BatRequest("emailが登録されていません", c)
		return
	}
	if user.Password != req.Password {
		controller.BatRequest("passwordが違います", c)
		return
	}
	token := createToken(user.ID)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
