package ping

import ginGonic "github.com/gin-gonic/gin"

func Ping (context *ginGonic.Context) {

	context.String(200, "pong")

}