package controller

import "github.com/gin-gonic/gin"

var Auth = authImpl{}

type authImpl struct {
}

func (self *authImpl) create(c *gin.Context) {

}
