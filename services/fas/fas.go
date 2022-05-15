package fas

import (
	"github.com/gin-gonic/gin"
	"wifidogfas/model"
	"wifidogfas/util"
)

func LoginHandler(c *gin.Context) {
	var loginForm model.FasLogin
	c.Bind(&loginForm)

	if loginForm.Username == "admin" && loginForm.Password == util.Sha256Sum("123456") {
		result := model.JsonTemplate{
			Code: 200,
			Msg:  "登录成功",
			Data: gin.H{
				"token": loginForm.MAC,
			},
		}
		c.JSON(200, &result)
	} else {
		result := model.JsonTemplate{
			Code: 403,
			Msg:  "账号密码错误",
			Data: nil,
		}
		c.JSON(200, &result)
	}

}
