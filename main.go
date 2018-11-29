package main

import (
	db "taskweb/database"
    r "taskweb/redis"
)


func init()  {
	r.Redis.Init(0)
}

func main() {

	defer db.SqlDB.Close()
	router := initRouter()
	router.Run(":8001")
	//outer.RunTLS(":8000", "./testdata/server.pem", "./testdata/server.key")
}

