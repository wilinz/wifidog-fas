package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/url"
	"strconv"
)

//0 - AUTH_DENIED - 用户防火墙用户被删除，用户被移除。
//6 - AUTH_VALIDATION_FAILED - 用户电子邮件验证超时，用户/防火墙被删除
//1 - AUTH_ALLOWED - 用户有效，如果不存在则添加防火墙规则
//5 - AUTH_VALIDATION - 允许用户访问电子邮件以在默认规则下获取验证电子邮件
//-1 - AUTH_ERROR - 验证过程中发生错误
type AuthTypes struct {
	AuthDenied           int
	AuthValidationFailed int
	AuthAllowed          int
	AuthValidation       int
	AuthError            int
}

var authTypes = AuthTypes{0, 6, 1, 5, -1}

type WifiDogAuth struct {
	Stage    string `form:"stage"`
	IP       string `form:"ip"`
	MAC      string `form:"mac"`
	Token    string `form:"token"`
	Incoming int64 `form:"incoming"`
	Outgoing int64 `form:"outgoing"`
	GwId     string `form:"gw_id"`
}

type WifiDogLogin struct {
	GatewayAddress string `form:"gw_address"`
	GatewayPort    int    `form:"gw_port"`
	GatewayID      string `form:"gw_id"`
	IP             string `form:"ip"`
	MAC            string `form:"mac"`
	URL            string `form:"url"`
}

type FasLogin struct {
	GatewayAddress string `form:"gw_address"`
	GatewayPort    int    `form:"gw_port"`
	Username       string `form:"username"`
	Password       string `form:"password"`
	MAC            string `form:"mac"`
}

type JsonTemplate struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func CustomRouterMiddle1(c *gin.Context) {
	if gin.Mode()==gin.DebugMode {
		header := c.Writer.Header()
		header.Set("Pragma", "no-cache")
		header.Set("Expires", strconv.Itoa(0))
		header.Add("Cache-Control", "no-cache")
		header.Add("Cache-Control", "no-store")
		header.Add("Cache-Control", "must-revalidate")
	}
}

func main() {

	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(CustomRouterMiddle1)
	r.Static("/html/", "html")
	r.LoadHTMLGlob("html/*.html")

	wifiDog := r.Group("/wifidog")
	{
		wifiDog.GET("/login", func(c *gin.Context) {
			var params WifiDogLogin
			c.Bind(&params)
			c.HTML(200, "login.html", &params)
		})

		wifiDog.GET("/auth/", func(c *gin.Context) {
			var p WifiDogAuth
			c.Bind(&p)

			if p.Stage == "login" {
				response := fmt.Sprint("Auth: ", authTypes.AuthAllowed)
				c.String(200, response)
			}else {
				fmt.Println("下载：",ConvertFlowUnit(p.Incoming),"上传：",ConvertFlowUnit(p.Outgoing))

				//下载超过300MB断网
				if p.Incoming>300*1024*1024 {
					response := fmt.Sprint("Auth: ", authTypes.AuthDenied)
					c.String(200, response)
				}

				c.String(200, "")
			}

		})

		wifiDog.GET("/ping/", func(c *gin.Context) {
			c.String(200, "Pong")
		})

		wifiDog.GET("/portal", func(c *gin.Context) {
			c.HTML(200,"portal.html",nil)
		})
	}

	fas := r.Group("/fas")
	{
		fas.Any("/login", func(c *gin.Context) {
			var loginForm FasLogin
			c.Bind(&loginForm)

			if loginForm.Username == "admin" && loginForm.Password == Sha256Sum("123456") {
				redirectURL := url.URL{
					Scheme: "http",
					Host:   fmt.Sprint(loginForm.GatewayAddress, ":", loginForm.GatewayPort),
					Path:   "/wifidog/auth",
				}

				query := redirectURL.Query()
				query.Set("token", loginForm.MAC)
				redirectURL.RawQuery = query.Encode()

				//在前端重定向，ajax重定向有跨域问题
				result := JsonTemplate{
					Code: 200,
					Msg:  "登录成功",
					Data: redirectURL.String(),
				}
				c.JSON(200, &result)
			} else {
				result := JsonTemplate{
					Code: 403,
					Msg:  "账号密码错误",
					Data: nil,
				}
				c.JSON(200, &result)
			}

		})
	}

	r.NoRoute(func(c *gin.Context) {
		log.Println("请求不匹配路由", c.Request.URL)
	})

	err := r.Run(":10010")
	if err != nil {
		log.Println(err)
	}
}
