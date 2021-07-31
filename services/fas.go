package services

import (
	"github.com/gin-gonic/gin"
	"wifidogfas/dao"
	"wifidogfas/util"
)

func FasLoginHandler(c *gin.Context) {
	var loginForm dao.FasLogin
	c.Bind(&loginForm)

	if loginForm.Username == "admin" && loginForm.Password == util.Sha256Sum("123456") {
		result := dao.JsonTemplate{
			Code: 200,
			Msg:  "登录成功",
			Data: gin.H{
				"token":loginForm.MAC,
			},
		}
		c.JSON(200, &result)
	} else {
		result := dao.JsonTemplate{
			Code: 403,
			Msg:  "账号密码错误",
			Data: nil,
		}
		c.JSON(200, &result)
	}

}
