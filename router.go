package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	. "taskweb/apis"
)

func initRouter() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	//cors
	router.Use(MiddleWare())

	router.GET("/", IndexApi)

	router.POST("/category", AddTbCategoryApi)

	router.POST("/delcategory", DelCategoryApi)

	router.GET("/categorys", GetTbCategorysApi)

	router.POST("/login", LoginApi)

	router.OPTIONS("/categorys", GetTbCategorysApi)

	router.OPTIONS("/login", LoginApi)

	router.OPTIONS("/jobs", GetJobsApi)

	router.GET("/jobs", GetJobsApi)

	router.POST("/job", AddJobApi)

	router.PUT("/job", PutJobApi)

	router.OPTIONS("/job", PutJobApi)

	router.POST("/node", AddTbNodeApi)

	router.POST("/delnode", DelNodeApi)

	router.GET("/nodes", GetTbNodesApi)

	router.OPTIONS("/nodes", GetTbNodesApi)

	router.POST("/command", AddTbCommandApi)

	router.GET("/commands", GetTbCommandsApi)

	router.OPTIONS("/commands", GetTbCommandsApi)

	router.OPTIONS("/logs", GetTbErrorsApi)

	router.GET("/logs", GetTbErrorsApi)

	router.OPTIONS("/performances", GetTbPerformancesApi)

	router.GET("/performances", GetTbPerformancesApi)
	//
	//router.PUT("/person/:id", ModPersonApi)
	//

	return router
}
//cors
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
			//c.Request.SetBasicAuth("x","x")
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")//允许访问所有域
			c.Writer.Header().Add("Access-Control-Allow-Headers","Content-Type")//header的类型
			//c.Writer.Header().Set("content-type","application/json") //返回数据格式是json
			c.Next()
		//}

	}
}
