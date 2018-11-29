package database

import (
	"errors"
	"strconv"
	db "taskweb/database"
)

type TbError struct {
	Id	int
	JobId	int
	Msg	string
	Createtime	int64
}

func ExistTbError(id int) (bool, error) {
	rows, err := db.Dtsc.Query("select count(0) Count from tb_error where id=?", id)
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

func InsertTbError(tb_error TbError) (int64, error) {
	result, err := db.Dtsc.Exec("insert into tb_error(job_id,msg,createtime) values(?,?,?)", tb_error.JobId,tb_error.Msg,tb_error.Createtime)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

func UpdateTbError(tb_error TbError) (bool, error) {
	result, err := db.Dtsc.Exec("update tb_error set job_id=?, msg=?, createtime=? where id=?", tb_error.JobId, tb_error.Msg, tb_error.Createtime, tb_error.Id)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

func GetTbError(id int) (tb_error TbError, err error) {
	rows, err := db.Dtsc.Query("select id, job_id, msg, createtime from tb_error where id=?", id)
	if err != nil {
		return tb_error, err
	}
	if len(rows) <= 0 {
		return tb_error, nil
	}
	tb_errors, err := _TbErrorRowsToArray(rows)
	if err != nil {
		return tb_error, err
	}
	return tb_errors[0], nil
}

func GetTbErrors() (tb_errors []TbError, err error) {
	rows, err := db.Dtsc.Query("select * from tb_error  order by createtime desc")
	if err != nil {
		return tb_errors, err
	}
	if len(rows) <= 0 {
		return tb_errors, nil
	}
	return _TbErrorRowsToArray(rows)
}
func GetTbErrorsByJob(jobid int) (tb_errors []TbError, err error) {
	rows, err := db.Dtsc.Query("select * from tb_error where job_id=? order by createtime desc",jobid)
	if err != nil {
		return tb_errors, err
	}
	if len(rows) <= 0 {
		return tb_errors, nil
	}
	return _TbErrorRowsToArray(rows)
}


func GetTbErrorRowCount() (count int, err error) {
	rows, err := db.Dtsc.Query("select count(0) Count from tb_error")
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

func _TbErrorRowsToArray(maps []map[string][]byte) ([]TbError, error) {
	models := make([]TbError, len(maps))
	var err error
	for index, obj := range maps {
		model := TbError{}
		model.Id, err = strconv.Atoi(string(obj["id"]))
		if err != nil {
			return nil, errors.New("parse Id error: " + err.Error())
		}
		model.JobId, err = strconv.Atoi(string(obj["job_id"]))
		if err != nil {
			return nil, errors.New("parse JobId error: " + err.Error())
		}
		model.Msg = string(obj["msg"])
		model.Createtime, err = strconv.ParseInt(string(obj["createtime"]), 10, 64)
		if err != nil {
			return nil, errors.New("parse Createtime error: " + err.Error())
		}
		models[index] = model
	}
	return models, err
}
