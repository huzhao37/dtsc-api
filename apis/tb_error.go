package apis

import (
	"net/http"
	"gopkg.in/gin-gonic/gin.v1"
	."taskweb/models"
	"taskweb/core"
	"strconv"
)


func GetTbErrorsApi(c *gin.Context){
	jobid,err:=strconv.Atoi(c.Request.FormValue("jobid"))
	if err!=nil {
		core.Logger.Println(err)
	}
	var tbErrors = make([]TbError, 0)
	if jobid==0{
		tbErrors, err =GetTbErrors()
		if err != nil {
			core.Logger.Fatalln(err)
		}
	}else {
		tbErrors, err =GetTbErrorsByJob(jobid)
		if err != nil {
			core.Logger.Fatalln(err)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": tbErrors,
	})

}
