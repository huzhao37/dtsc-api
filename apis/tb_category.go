package apis

import (
	"net/http"
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	."taskweb/models"
	"taskweb/core"
	"time"
	"strconv"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func AddTbCategoryApi(c *gin.Context) {
	name := c.Request.FormValue("name")

	createtime:=time.Now().Unix()//,_:=strconv.ParseInt(c.Request.FormValue("createtime"),10,64)
	remark := c.Request.FormValue("remark")

	 p:= TbCategory{Name: name, Createtime: createtime,Remark:remark}

	ra, err := InsertTbCategory(p)
	if err != nil {
		core.Logger.Fatalln(err)
	}
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
func GetTbCategorysApi(c *gin.Context){

	var tbcategorys = make([]TbCategory, 0)
	tbcategorys, err :=GetTbCategorys()
	if err != nil {
		core.Logger.Fatalln(err)
	}
	//for _,v := range tbcategorys {
	//	fmt.Sprintf(strconv.Itoa(v.Id)+":"+v.Name)
	//}
	//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")//允许访问所有域
	//c.Writer.Header().Add("Access-Control-Allow-Headers","Content-Type")//header的类型
	//c.Writer.Header().Set("content-type","application/json") //返回数据格式是json
	//msg := fmt.Sprintf("get successful %d", len(persons))
	c.JSON(http.StatusOK, gin.H{

		"msg": tbcategorys,
	})
	//c.Writer.WriteString("success data")
}

func DelCategoryApi(c *gin.Context){
	id ,err:=strconv.Atoi(c.Request.FormValue("id"))
	if err!=nil {
		core.Logger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	rows, err :=DelTbCategory(id)
	if err != nil {
		core.Logger.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": rows,
	})
}

