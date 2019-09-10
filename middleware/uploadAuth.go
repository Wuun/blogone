package middleware

import (
	"bblog/conf"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"os"
)

func UploadAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		pas := session.Get("password")
		fmt.Println(os.Getenv("WEB_PASSWORD"))
		if pas.(string) == os.Getenv("WEB_PASSWORD") {
			c.Next()
			return
		}
		loginPage := "http://" + conf.G_CONF.WebAddr + "/api/v1/login"
		fmt.Println(pas.(string))
		c.Redirect(200,loginPage)
		c.Abort()
	}
}
