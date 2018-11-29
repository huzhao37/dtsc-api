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


func AddTbNodeApi(c *gin.Context) {
	name := c.Request.FormValue("name")
	ip := c.Request.FormValue("ip")
	createtime:=time.Now().Unix()//,_:=strconv.ParseInt(c.Request.FormValue("createtime"),10,64)
	remark := c.Request.FormValue("remark")

	 p:= TbNode{Name: name, Createtime: createtime,Remark:remark,Ip:ip}

	ra, err := InsertTbNode(p)
	if err != nil {
		core.Logger.Fatalln(err)
	}
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func GetTbNodesApi(c *gin.Context){

	var tbnodes = make([]TbNode, 0)
	tbnodes, err :=GetTbNodes()
	if err != nil {
		core.Logger.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{

		"msg": tbnodes,
	})
}

func DelNodeApi(c *gin.Context){
	id ,err:=strconv.Atoi(c.Request.FormValue("id"))
	if err!=nil {
		core.Logger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	rows, err :=DelTbNode(id)
	if err != nil {
		core.Logger.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": rows,
	})
}
