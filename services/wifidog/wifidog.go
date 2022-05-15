package wifidog

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"wifidogfas/enum"
	"wifidogfas/model"
	"wifidogfas/util"
)

func PostAuthHandler(c *gin.Context) {
	var p model.AuthInfo
	c.Bind(&p)

	fmt.Printf("%#v\n", p)

	authOptions := make([]model.AuthOp, 0)
	for _, client := range p.Clients {
		fmt.Println("下载：", util.ConvertFlowUnit(client.Incoming), "上传：", util.ConvertFlowUnit(client.Outgoing))

		authOption := &model.AuthOp{
			ID:       client.ID,
			AuthCode: enum.AuthTypes.AuthAllowed,
		}
		//下载超过300MB断网
		//if client.Incoming > 1000*1024*1024 {
		//	fmt.Println("断网")
		//	authOption.AuthCode = enum.AuthTypes.AuthDenied
		//}

		authOptions = append(authOptions, *authOption)
	}
	resp := &model.AuthResponse{
		GwID:   p.GwID,
		AuthOp: authOptions,
	}
	fmt.Printf("返回信息：%#v\n", resp)
	c.JSON(200, resp)
}

func GetAuthHandler(c *gin.Context) {
	var p model.AuthInfoQuery
	c.BindQuery(&p)
	fmt.Printf("%#v\n", p)
	response := fmt.Sprint("Auth: ", enum.AuthTypes.AuthAllowed)
	c.String(200, response)
}

func LoginHandler(c *gin.Context) {
	var params model.WifiDogLogin
	c.BindQuery(&params)
	fmt.Printf("%#v\n", params)
	c.HTML(200, "login.html", &params)
}

func PingHandler(c *gin.Context) {
	info := new(model.PingInfo)
	c.BindQuery(info)
	fmt.Printf("%#v\n", info)
	c.String(200, "Pong")
}

func PortalHandler(c *gin.Context) {
	c.HTML(200, "portal.html", nil)
}
