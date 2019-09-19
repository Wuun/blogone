package middleware

import (
	"bblog/conf"
	"fmt"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//UploadAuth is middleware use to judge use is admin or not.
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		loginPage := "http://" + conf.G_CONF.WebAddr + "/api/v1/login"
		session := sessions.Default(c)
		pas := session.Get("password")
		if pas != nil {
			if pas.(string) == os.Getenv("WEB_PASSWORD") {
				c.Next()
				return
			}
			fmt.Println(pas.(string))
			c.Redirect(200, loginPage)
			c.Abort()
		}
		c.Redirect(307, loginPage)
		c.Abort()
		return
	}
}
