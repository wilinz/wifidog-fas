package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"wifidogfas/services"
)

var Router = gin.Default()

func CustomRouterMiddle1(c *gin.Context) {
	if gin.Mode() == gin.DebugMode {
		header := c.Writer.Header()
		header.Set("Pragma", "no-cache")
		header.Set("Expires", strconv.Itoa(0))
		header.Add("Cache-Control", "no-cache")
		header.Add("Cache-Control", "no-store")
		header.Add("Cache-Control", "must-revalidate")
	}
}

func Run() {
	//gin.SetMode(gin.ReleaseMode)
	Router.Use(CustomRouterMiddle1)
	Router.Static("/html/", "html")
	Router.LoadHTMLGlob("html/*.html")
	Router.NoRoute(func(c *gin.Context) {
		log.Println("请求不匹配路由", c.Request.URL)
	})

	wifiDog := Router.Group("/wifidog")
	{
		wifiDog.GET("/login", services.WifiDogLoginHandler)
		wifiDog.GET("/auth/", services.WifiDogAuthHandler)
		wifiDog.GET("/ping/",  services.WifiDogPingHandler)
		wifiDog.GET("/portal", services.WifiDogPortalHandler)
	}

	fas := Router.Group("/fas")
	{
		fas.POST("/login",services.FasLoginHandler)
	}

	err := Router.Run(":10010")
	if err != nil {
		log.Println(err)
	}
}