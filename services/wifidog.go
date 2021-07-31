package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"wifidogfas/dao"
	"wifidogfas/enum"
	"wifidogfas/util"
)

func WifiDogAuthHandler(c *gin.Context) {
	var p dao.WifiDogAuth
	c.Bind(&p)

	if p.Stage == "login" {
		response := fmt.Sprint("Auth: ", enum.AuthTypes.AuthAllowed)
		c.String(200, response)
	} else {
		fmt.Println("下载：", util.ConvertFlowUnit(p.Incoming), "上传：", util.ConvertFlowUnit(p.Outgoing))

		//下载超过300MB断网
		if p.Incoming > 300*1024*1024 {
			response := fmt.Sprint("Auth: ", enum.AuthTypes.AuthDenied)
			c.String(200, response)
		}

		c.String(200, "")
	}

}

func WifiDogLoginHandler(c *gin.Context) {
	var params dao.WifiDogLogin
	c.Bind(&params)
	c.HTML(200, "login.html", &params)
}

func WifiDogPingHandler(c *gin.Context) {
	c.String(200, "Pong")
}

func WifiDogPortalHandler(c *gin.Context) {
	c.HTML(200, "portal.html", nil)
}
