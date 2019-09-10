package api

import (
	"bblog/serializer"

	"github.com/gin-gonic/gin"
)

//Ping is ping to server.
func Ping(ctx *gin.Context) {
	ctx.JSON(200, serializer.BuildResponse(0, "pong", nil, nil))
}
