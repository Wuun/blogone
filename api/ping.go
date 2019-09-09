package api

import (
	"blogone/serializer"
	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context){
	ctx.JSON(200,serializer.BuildResponse(0,"pong",nil,nil))
}
