package database

import (
	"errors"
	"strconv"
	db "taskweb/database"
)

type TbPerformance struct {
	Id	int
	JobId	int
	NodeId	int
	Cpu	float64
	Memory	float64
	Installdirsize	float64
	Updatetime	int64
	Remark	string
}

func ExistTbPerformance(id int) (bool, error) {
	rows, err := db.Dtsc.Query("select count(0) Count from tb_performance where id=?", id)
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

func InsertTbPerformance(tb_performance TbPerformance) (int64, error) {
	result, err := db.Dtsc.Exec("insert into tb_performance(job_id,node_id,cpu,memory,installdirsize,updatetime,remark) values(?,?,?,?,?,?,?)", tb_performance.JobId,tb_performance.NodeId,tb_performance.Cpu,tb_performance.Memory,tb_performance.Installdirsize,tb_performance.Updatetime,tb_performance.Remark)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

func UpdateTbPerformance(tb_performance TbPerformance) (bool, error) {
	result, err := db.Dtsc.Exec("update tb_performance set job_id=?, node_id=?, cpu=?, memory=?, installdirsize=?, updatetime=?, remark=? where id=?", tb_performance.JobId, tb_performance.NodeId, tb_performance.Cpu, tb_performance.Memory, tb_performance.Installdirsize, tb_performance.Updatetime, tb_performance.Remark, tb_performance.Id)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

func GetTbPerformance(id int) (tb_performance TbPerformance, err error) {
	rows, err := db.Dtsc.Query("select id, job_id, node_id, cpu, memory, installdirsize, updatetime, remark from tb_performance where id=?", id)
	if err != nil {
		return tb_performance, err
	}
	if len(rows) <= 0 {
		return tb_performance, nil
	}
	tb_performances, err := _TbPerformanceRowsToArray(rows)
	if err != nil {
		return tb_performance, err
	}
	return tb_performances[0], nil
}

func GetTbPerformanceRowCount() (count int, err error) {
	rows, err := db.Dtsc.Query("select count(0) Count from tb_performance")
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

func GetTbPerformances() (tb_performances []TbPerformance, err error) {
	rows, err := db.Dtsc.Query("select * from tb_performance")
	if err != nil {
		return tb_performances, err
	}
	if len(rows) <= 0 {
		return tb_performances, nil
	}
	return _TbPerformanceRowsToArray(rows)
}


func _TbPerformanceRowsToArray(maps []map[string][]byte) ([]TbPerformance, error) {
	models := make([]TbPerformance, len(maps))
	var err error
	for index, obj := range maps {
		model := TbPerformance{}
		model.Id, err = strconv.Atoi(string(obj["id"]))
		if err != nil {
			return nil, errors.New("parse Id error: " + err.Error())
		}
		model.JobId, err = strconv.Atoi(string(obj["job_id"]))
		if err != nil {
			return nil, errors.New("parse JobId error: " + err.Error())
		}
		model.NodeId, err = strconv.Atoi(string(obj["node_id"]))
		if err != nil {
			return nil, errors.New("parse NodeId error: " + err.Error())
		}
		model.Cpu, err = strconv.ParseFloat(string(obj["cpu"]), 64)
		if err != nil {
			return nil, errors.New("parse Cpu error: " + err.Error())
		}
		model.Memory, err = strconv.ParseFloat(string(obj["memory"]), 64)
		if err != nil {
			return nil, errors.New("parse Memory error: " + err.Error())
		}
		model.Installdirsize, err = strconv.ParseFloat(string(obj["installdirsize"]), 64)
		if err != nil {
			return nil, errors.New("parse Installdirsize error: " + err.Error())
		}
		model.Updatetime, err = strconv.ParseInt(string(obj["updatetime"]), 10, 64)
		if err != nil {
			return nil, errors.New("parse Updatetime error: " + err.Error())
		}
		model.Remark = string(obj["remark"])
		models[index] = model
	}
	return models, err
}
