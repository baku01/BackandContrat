package Application

import (
	"BackandContrat/Application/Routes"
	"github.com/gin-gonic/gin"
	"log"
)

func InitWebServer() {
	app := gin.Default()
	Routes.InitRoutes(&app.RouterGroup)
	if err := app.Run(":8080"); err != nil {
		log.Fatalf("Erro a inciar Servidorweb!")
	}
}
