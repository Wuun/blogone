package api

import (
	"bblog/serializer"
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
)

func Login(c *gin.Context) {
	c.HTML(200,"login.html",nil)
	return
}

func Authenticate(c *gin.Context) {
	var(
		body []byte
		err error
		)
	if body,err = ioutil.ReadAll(c.Request.Body);err != nil {
		c.JSON(200,serializer.Response{
			Status:4002,
			Msg:"nothing you post.",
			Error:err,
		})
		return
	}

	loginAuth := serializer.LoginAuthentication{}
	err = json.Unmarshal(body,&loginAuth)
	if err != nil {
		c.JSON(200,serializer.Response{
			Status:4002,
			Msg:"wrong thing you post.",
			Error:err,
		})
		return
	}

	if loginAuth.Password != os.Getenv("WEB_PASSWORD") {
		c.JSON(200,serializer.Response{
			Status:4002,
			Msg:"wrong password.",
			Error:err,
		})
		return
	}

	session := sessions.Default(c)
	session.Set("password",os.Getenv("WEB_PASSWORD"))
	err = session.Save()
	if err !=nil {
		c.JSON(200,serializer.Response{
			Status:4002,
			Msg:"login failed.",
			Error:err,
		})

		return
	}

	c.JSON(200,&serializer.Response{
		Status:0,
		Msg:"success",
	})
}