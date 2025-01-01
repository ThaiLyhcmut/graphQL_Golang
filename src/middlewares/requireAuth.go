package middlewares

import (
	"ThaiLy/configs"
	helper "ThaiLy/helpers"
	"ThaiLy/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func RequireAuth(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(401, gin.H{
			"code":    "error",
			"massage": "Nguoi dung chua dang nhap 1",
		})
		c.Abort()
		return
	}
	if len(token) > 7 && strings.HasPrefix(token, "Bearer ") {
		token = token[7:] // Lấy token sau "Bearer "
	} else {
		c.JSON(401, gin.H{
			"code":    "error",
			"message": "Nguoi dung chua dang nhap 2",
		})
		c.Abort()
		return
	}
	Claims, _ := helper.ParseJWT(token)
	if Claims != nil {
		var account = models.Account{}
		result := configs.GetDB().Where("id = ?", Claims.ID).First(&account)
		if result.Error != nil {
			c.JSON(401, gin.H{
				"code":    "error",
				"massage": "Nguoi dung chua dang nhap 4",
			})
			c.Abort()
			return
		}
		c.Set("account", account)
	}
	c.Next()
}
