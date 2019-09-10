package middleware

import (
	"bblog/conf"
	"fmt"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//UploadAuth is middleware use to judge use is admin or not.
func UploadAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		pas := session.Get("password")
		fmt.Println(os.Getenv("WEB_PASSWORD"))
		if pas != nil {
			if pas.(string) == os.Getenv("WEB_PASSWORD") {
				c.Next()
				return
			}
			loginPage := "http://" + conf.G_CONF.WebAddr + "/api/v1/login"
			fmt.Println(pas.(string))
			c.Redirect(200, loginPage)
			c.Abort()
		}
		c.Redirect(307, "http://127.0.0.1:8948/api/v1/login")
		c.Abort()
		return
	}
}
