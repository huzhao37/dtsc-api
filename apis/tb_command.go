package apis

import (
	"net/http"
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	."taskweb/models"
	"strconv"
	"taskweb/core"
	"time"
	redis "taskweb/redis"
)
///// <summary>
///// 删除
///// </summary>
//Delete=-1,
///// <summary>
///// 停止
///// </summary>
//Stop=0,
///// <summary>
///// 启动
///// </summary>
//Start=1,
///// <summary>
///// 重启
///// </summary>
//ReStart=2

func AddTbCommandApi(c *gin.Context) {
	jobid,err :=strconv.Atoi(c.Request.FormValue("jobid"))
	if err != nil {
		core.Logger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	commandtype,err:=strconv.Atoi(c.Request.FormValue("commandtype"))
	if err != nil {
		core.Logger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	time:=time.Now().Unix()

	 p:= TbCommand{Jobid: jobid, Time: time,Commandtype:commandtype,Success:false}

	ra, err := InsertTbCommand(p)
	if err != nil {
		core.Logger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	p.Id=int(ra)
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
	//cmd
	redis.Redis.Set(strconv.Itoa(jobid),p)
}
func GetTbCommandsApi(c *gin.Context){

	var tbCommands = make([]TbCommand, 0)
	tbCommands, err :=GetCommands()
	if err != nil {
		core.Logger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{

		"msg": tbCommands,
	})
}
