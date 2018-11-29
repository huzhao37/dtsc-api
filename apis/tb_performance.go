package apis

import (
	"net/http"
	"gopkg.in/gin-gonic/gin.v1"
	."taskweb/models"
	"taskweb/core"
)


func GetTbPerformancesApi(c *gin.Context){
	   var tbPerformances = make([]TbPerformance, 0)
		tbPerformances, err :=GetTbPerformances()
		if err != nil {
			core.Logger.Fatalln(err)
		}
	c.JSON(http.StatusOK, gin.H{
		"msg": tbPerformances,
	})

}
