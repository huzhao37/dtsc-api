package database

import (
	"errors"
	"strconv"
	db "taskweb/database"
	"strings"
)

type TbJob struct {
	Id	int
	Name	string
	Single	bool
	Datamap	string
	NodeId	int
	CategoryId	int
	UserId	int
	State	 bool
	Version	int
	Runcount	int
	Createtime	int64
	Lastedstart	int64
	Lastedend	int64
	Nextstart	int64
	Remark	string
	Cron	string
}

func ExistTbJob(id int) (bool, error) {
	rows, err := db.Dtsc.Query("select count(0) Count from tb_job where id=?", id)
	if err != nil {
		return false, err
	}
	if len(rows) <= 0 {
		return false, nil
	}
	for _, obj := range rows {
		count, err := strconv.Atoi(string(obj["Count"]))
		if err != nil {
			return false, errors.New("parse Count error: " + err.Error())
		}
		return count > 0, nil
	}
	return false, nil
}

func InsertTbJob(tb_job TbJob) (int64, error) {
	result, err := db.Dtsc.Exec("insert into tb_job(name,single,datamap,node_id,category_id,user_id,state,version,runcount,createtime,lastedstart,lastedend,nextstart,remark,cron) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", tb_job.Name,tb_job.Single,tb_job.Datamap,tb_job.NodeId,tb_job.CategoryId,tb_job.UserId,tb_job.State,tb_job.Version,tb_job.Runcount,tb_job.Createtime,tb_job.Lastedstart,tb_job.Lastedend,tb_job.Nextstart,tb_job.Remark,tb_job.Cron)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

func UpdateTbJob(tb_job TbJob) (bool, error) {
	result, err := db.Dtsc.Exec("update tb_job set name=?, single=?, datamap=?, node_id=?, category_id=?, user_id=?, state=?, version=?, runcount=?, createtime=?, lastedstart=?, lastedend=?, nextstart=?, remark=?, cron=? where id=?", tb_job.Name, tb_job.Single, tb_job.Datamap, tb_job.NodeId, tb_job.CategoryId, tb_job.UserId, tb_job.State, tb_job.Version, tb_job.Runcount, tb_job.Createtime, tb_job.Lastedstart, tb_job.Lastedend, tb_job.Nextstart, tb_job.Remark, tb_job.Cron, tb_job.Id)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

func GetTbJob(id int) (tb_job TbJob, err error) {
	rows, err := db.Dtsc.Query("select id, name, single, datamap, node_id, category_id, user_id, state, version, runcount, createtime, lastedstart, lastedend, nextstart, remark, cron from tb_job where id=?", id)
	if err != nil {
		return tb_job, err
	}
	if len(rows) <= 0 {
		return tb_job, nil
	}
	tb_jobs, err := _TbJobRowsToArray(rows)
	if err != nil {
		return tb_job, err
	}
	return tb_jobs[0], nil
}

func GetTbJobRowCount() (count int, err error) {
	rows, err := db.Dtsc.Query("select count(0) Count from tb_job")
	if err != nil {
		return -1, err
	}
	if len(rows) <= 0 {
		return -1, nil
	}
	for _, obj := range rows {
		count, err := strconv.Atoi(string(obj["Count"]))
		if err != nil {
			return -1, errors.New("parse Count error: " + err.Error())
		}
		return count, nil
	}
	return -1, nil
}

func _TbJobRowsToArray(maps []map[string][]byte) ([]TbJob, error) {
	models := make([]TbJob, len(maps))
	var err error
	for index, obj := range maps {
		model := TbJob{}
		model.Id, err = strconv.Atoi(string(obj["id"]))
		if err != nil {
			return nil, errors.New("parse Id error: " + err.Error())
		}
		model.Name = string(obj["name"])

		if strings.ToLower(string(obj["single"])) == "true" || obj["single"][0] == 1 {
			model.Single = true
		} else {
			model.Single = false
		}
		model.Datamap = string(obj["datamap"])
		model.NodeId, err = strconv.Atoi(string(obj["node_id"]))
		if err != nil {
			return nil, errors.New("parse NodeId error: " + err.Error())
		}
		model.CategoryId, err = strconv.Atoi(string(obj["category_id"]))
		if err != nil {
			return nil, errors.New("parse CategoryId error: " + err.Error())
		}
		model.UserId, err = strconv.Atoi(string(obj["user_id"]))
		if err != nil {
			return nil, errors.New("parse UserId error: " + err.Error())
		}
		if strings.ToLower(string(obj["single"])) == "true" || obj["single"][0] == 1 {
			model.Single = true
		} else {
			model.Single = false
		}
		if strings.ToLower(string(obj["state"])) == "true" || obj["state"][0] == 1 {
			model.State = true
		} else {
			model.State = false
		}
		model.Version, err = strconv.Atoi(string(obj["version"]))
		if err != nil {
			return nil, errors.New("parse Version error: " + err.Error())
		}
		model.Runcount, err = strconv.Atoi(string(obj["runcount"]))
		if err != nil {
			return nil, errors.New("parse Runcount error: " + err.Error())
		}
		model.Createtime, err = strconv.ParseInt(string(obj["createtime"]), 10, 64)
		if err != nil {
			return nil, errors.New("parse Createtime error: " + err.Error())
		}
		model.Lastedstart, err = strconv.ParseInt(string(obj["lastedstart"]), 10, 64)
		if err != nil {
			return nil, errors.New("parse Lastedstart error: " + err.Error())
		}
		model.Lastedend, err = strconv.ParseInt(string(obj["lastedend"]), 10, 64)
		if err != nil {
			return nil, errors.New("parse Lastedend error: " + err.Error())
		}
		model.Nextstart, err = strconv.ParseInt(string(obj["nextstart"]), 10, 64)
		if err != nil {
			return nil, errors.New("parse Nextstart error: " + err.Error())
		}
		model.Remark = string(obj["remark"])
		model.Cron = string(obj["cron"])
		models[index] = model
	}
	return models, err
}
func GetJobs() (jobs []TbJob, err error) {
	rows, err := db.Dtsc.Query("select * from tb_job ")
	if err != nil {
		return jobs, err
	}
	if len(rows) <= 0 {
		return jobs, nil
	}
	return _TbJobRowsToArray(rows)
}
