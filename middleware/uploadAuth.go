package middleware

import (
	"blogone/conf"
	"blogone/serializer"
	"github.com/gin-gonic/gin"
)

func UploadAuth() gin.HandlerFunc{
	return func(context *gin.Context) {
		if context.Param("password") != conf.G_CONF.WebsitePassWord {
			context.JSON(403,serializer.Response{
				Status:40003,
				Msg:"you are invaild",
			})
			context.Abort()
			return
		}
		context.Next()
	}
}