package app

import (
	ginGonic "github.com/gin-gonic/gin"
	"github.com/mercadolibre/taller-go/src/api/controllers/myml"
	"github.com/mercadolibre/taller-go/src/api/controllers/ping"
)

const (
	port = ":8080"
)

func setupRouter() *ginGonic.Engine {
	r := ginGonic.Default()
	r.GET("/ping", ping.Ping)
	r.GET("/myml/:id", myml.GetUserDataReceiver)
	return r
}

func StartApp() {

	r := setupRouter()
	r.Run(":8080")
}


