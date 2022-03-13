package main

import (
	"ginweb/controller"
	"ginweb/tools"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := tools.ParesConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}
	_, erro := tools.OrmEngine(cfg)
	if erro != nil {
		panic(err.Error())
	}
	engins := gin.Default()
	regsiterRouter(engins)
	engins.Run(cfg.AppHost + ":" + cfg.AppPort)
}

func regsiterRouter(router *gin.Engine) {
	new(controller.MemberController).Router(router)
}
