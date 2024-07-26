package routes

import (
	"github.com/chi-net/kirara/core/class/tokens"
	"github.com/chi-net/kirara/core/utils"
	"github.com/gin-gonic/gin"
)

type UserLoginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func HandleUserLogin(c *gin.Context) {
	loginForm := UserLoginForm{}

	if c.ShouldBind(&loginForm) != nil {
		c.JSON(400, gin.H{
			"code":   400,
			"status": "no enough content",
		})
		return
	}

	if utils.CheckLogin(loginForm.Username, loginForm.Password) {
		token := tokens.Create(1)
		c.JSON(200, gin.H{
			"code":  200,
			"token": token,
		})
	} else {
		c.JSON(400, gin.H{
			"code":   400,
			"status": "failed",
		})
	}
}
