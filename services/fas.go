package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/url"
	"wifidogfas/dao"
	"wifidogfas/util"
)

func FasLoginHandler(c *gin.Context) {
	var loginForm dao.FasLogin
	c.Bind(&loginForm)

	if loginForm.Username == "admin" && loginForm.Password == util.Sha256Sum("123456") {
		redirectURL := url.URL{
			Scheme: "http",
			Host:   fmt.Sprint(loginForm.GatewayAddress, ":", loginForm.GatewayPort),
			Path:   "/wifidog/auth",
		}

		query := redirectURL.Query()
		query.Set("token", loginForm.MAC)
		redirectURL.RawQuery = query.Encode()

		//在前端重定向，ajax重定向有跨域问题
		result := dao.JsonTemplate{
			Code: 200,
			Msg:  "登录成功",
			Data: redirectURL.String(),
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
