package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/makki0205/vuejs/model"
	"github.com/makki0205/vuejs/service"
)

var User = userImpl{}

type userImpl struct {
}

func (self userImpl) Create(c *gin.Context) {
	req, ok := self.getUser(c)
	if !ok {
		return
	}
	user := service.User.FindByEmail(req.Email)
	if user != nil {
		BatRequest("そのメアド使われてまーす＞＜", c)
		return
	}
	res := service.User.Create(req)
	c.JSON(http.StatusOK, gin.H{
		"user": res,
	})

}
func (self userImpl) getUser(c *gin.Context) (model.User, bool) {
	var req model.User
	err := c.BindJSON(&req)
	if err != nil {
		BatRequest(err.Error(), c)
		return req, false
	}
	return req, true
}
