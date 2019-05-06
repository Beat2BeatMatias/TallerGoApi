package main

import (
	ginGonic "github.com/gin-gonic/gin"
	"github.com/mercadolibre/taller-go/src/api/controllers/myml"
	"github.com/mercadolibre/taller-go/src/api/controllers/ping"
)


const (
	port = ":8080"
)

var (
	router = ginGonic.Default()
)


func main() {
	router.GET("/ping", ping.Ping)
	router.GET("/myml/:id", myml.GetUserDataReceiver)

	router.Run(port)
}

