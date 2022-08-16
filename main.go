package main

import (
	"TREgitim/Config"
	"TREgitim/Routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	Config.Connect()
	Routes.LoginRoute(router)
	Routes.RegisterRoute(router)
	Routes.AdvertRoute(router)
	Routes.Application(router)
	Routes.Company(router)
	Routes.Todo(router)
	router.Run(":8080")
}
