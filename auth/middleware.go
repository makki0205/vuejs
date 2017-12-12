package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": "認証よろwwwww",
		})
		c.Abort()
		return
	}
	userId, err := perseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": "トークンがちがいまーす笑笑おつーーー",
		})
		return
	}
	c.Set("user_id", userId)
}
