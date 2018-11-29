package apis

import (
	"net/http"
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	."taskweb/models"
	"strconv"
	"taskweb/core"
	"time"
	"os"
	"io"

)


func AddJobApi(c *gin.Context) {
	name := c.Request.FormValue("name")
    single,err:=strconv.ParseBool(c.Request.FormValue("single"))
    if err!=nil {
		core.Logger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	createtime:=time.Now().Unix()//strconv.ParseInt(c.Request.FormValue("createtime"),10,64)

	remark := c.Request.FormValue("remark")
	nodeid,err:=strconv.Atoi(c.Request.FormValue("nodeid"))
	if err!=nil {
		core.Logger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	categoryid,err:=strconv.Atoi(c.Request.FormValue("categoryid"))
	if err!=nil {
		core.Logger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	userid,err:=strconv.Atoi(c.Request.FormValue("userid"))
	if err!=nil {
		core.Logger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}

	state,err:=strconv.ParseBool(c.Request.FormValue("state"))
	if err!=nil {
		core.Logger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	version,err:=strconv.Atoi(c.Request.FormValue("version"))
	if err!=nil {
		core.Logger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	datamap := c.Request.FormValue("datamap")
	cron := c.Request.FormValue("cron")

	id,err:=strconv.Atoi(c.Request.FormValue("id"))
	if err!=nil {
		core.Logger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	 p:= TbJob{Id:id,Name: name, Createtime: createtime,Remark:remark,Single:single,NodeId:nodeid,
	 CategoryId:categoryid,UserId:userid,State:state,Version:version,
	 Datamap:datamap,Cron:cron}
if(id==0) {
	ra, err := InsertTbJob(p)
	if err != nil {
		core.Logger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	UpLoadApi(c,int(ra),p.Version)
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,})
}else{
	ra, err := UpdateTbJob(p)
	if err!=nil {
		core.Logger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	UpLoadApi(c,id,p.Version)

	msg := fmt.Sprintf("update successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}


		//cmd
		//redis.Redis.Set(strconv.Itoa(p.Id),p.State)
}

func GetJobsApi(c *gin.Context){

	var jobs = make([]TbJob, 0)
	jobs, err :=GetJobs()
	if err != nil {
		core.Logger.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": jobs,
	})
	//c.Writer.WriteString("success data")
}

func PutJobApi(c *gin.Context) {

	name := c.Request.FormValue("name")
	single,err:=strconv.ParseBool(c.Request.FormValue("single"))
	if err!=nil {
		core.Logger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}

	id,err:=strconv.Atoi(c.Request.FormValue("id"))
	if err!=nil {
		core.Logger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	createtime,err:=strconv.ParseInt(c.Request.FormValue("createtime"),10,64)
	if err!=nil {
		core.Logger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	remark := c.Request.FormValue("remark")
	nodeid,err:=strconv.Atoi(c.Request.FormValue("nodeid"))
	if err!=nil {
		core.Logger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	categoryid,err:=strconv.Atoi(c.Request.FormValue("categoryid"))
	if err!=nil {
		core.Logger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	userid,err:=strconv.Atoi(c.Request.FormValue("userid"))
	if err!=nil {
		core.Logger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}

	state,err:=strconv.ParseBool(c.Request.FormValue("state"))
	if err!=nil {
		core.Logger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	version,err:=strconv.Atoi(c.Request.FormValue("version"))
	if err!=nil {
		core.Logger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	runcount,err:=strconv.Atoi(c.Request.FormValue("runcount"))
	if err!=nil {
		core.Logger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	lastedstart,err:=strconv.ParseInt(c.Request.FormValue("lastedstart"),10,64)
	if err!=nil {
		core.Logger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	lastedend,err:=strconv.ParseInt(c.Request.FormValue("lastedstart"),10,64)
	if err!=nil {
		core.Logger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	nextstart,err:=strconv.ParseInt(c.Request.FormValue("nextstart"),10,64)
	if err!=nil {
		core.Logger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	datamap := c.Request.FormValue("datamap")
	cron := c.Request.FormValue("cron")
	p:= TbJob{Id: id, Name: name, Createtime: createtime,Remark:remark,Single:single,NodeId:nodeid,
		CategoryId:categoryid,UserId:userid,State:state,Version:version,Runcount:runcount,
		Lastedend:lastedend,Lastedstart:lastedstart,Nextstart:nextstart,Datamap:datamap,Cron:cron}

	ra, err := UpdateTbJob(p)
	if err!=nil {
		core.Logger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	//UpLoadApi(c,id,p.Version)

	msg := fmt.Sprintf("update successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}


// 处理/upload 逻辑
func UpLoadApi(c *gin.Context,jobid int,version int) {//w http.ResponseWriter, r *http.Request
        r:=c.Request
        w:=c.Writer
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./dll/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			core.Logger.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		var zip= TbZip{JobId:jobid,Version:version	}
		zip.Zipfilename=handler.Filename
		zip.Time=time.Now().Unix()
		zip.Zipfile,err=core.ReadAll("./dll/"+handler.Filename)
		if err!=nil {
			core.Logger.Println(err)
		}
		InsertTbZip(zip)

}
