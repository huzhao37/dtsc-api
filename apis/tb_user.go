package apis

import (
	"net/http"
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	."taskweb/models"
	"strconv"
	"taskweb/core"
)
type user struct {
	data data     `json:"data"`
}

type data struct{
	Name string     `json:"name"`
	Pwd     string `json:"pwd"`
}
func LoginApi(c *gin.Context) {

	name := c.Request.PostFormValue("name")
	pwd:=c.Request.PostFormValue("pwd")
	user, err := UserLogin(name,pwd)
	if err != nil {
		core.Logger.Fatalln(err)
		c.JSON(http.StatusBadGateway, gin.H{
			"msg": "server exception",
		})
	}
	if user.Id==0{
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "login failed",
		})
	}
	c.JSON(http.StatusOK, gin.H{
			"msg": user,
		})
}
func GetUsersApi(c *gin.Context){

	var persons = make([]TbCategory, 0)
	persons, err :=GetTbCategorys()
	if err != nil {
		core.Logger.Fatalln(err)
	}
	for _,v := range persons {
		fmt.Sprintf(strconv.Itoa(v.Id)+":"+v.Name)
	}
	msg := fmt.Sprintf("get successful %d", len(persons))
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
