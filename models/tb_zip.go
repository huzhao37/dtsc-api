package database

import (
	"errors"
	"strconv"
	db "taskweb/database"
)

type TbZip struct {
	Id	int
	JobId	int
	Version	int
	Zipfilename	string
	Zipfile	[]byte
	Time	int64
}

func ExistTbZip(id int) (bool, error) {
	rows, err := db.Dtsc.Query("select count(0) Count from tb_zip where id=?", id)
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

func InsertTbZip(tb_zip TbZip) (int64, error) {
	result, err := db.Dtsc.Exec("insert into tb_zip(job_id,version,zipfilename,zipfile,time) values(?,?,?,?,?)", tb_zip.JobId,tb_zip.Version,tb_zip.Zipfilename,tb_zip.Zipfile,tb_zip.Time)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

func UpdateTbZip(tb_zip TbZip) (bool, error) {
	result, err := db.Dtsc.Exec("update tb_zip set job_id=?, version=?, zipfilename=?, zipfile=?, time=? where id=?", tb_zip.JobId, tb_zip.Version, tb_zip.Zipfilename, tb_zip.Zipfile, tb_zip.Time, tb_zip.Id)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

func GetTbZip(id int) (tb_zip TbZip, err error) {
	rows, err := db.Dtsc.Query("select id, job_id, version, zipfilename, zipfile, time from tb_zip where id=?", id)
	if err != nil {
		return tb_zip, err
	}
	if len(rows) <= 0 {
		return tb_zip, nil
	}
	tb_zips, err := _TbZipRowsToArray(rows)
	if err != nil {
		return tb_zip, err
	}
	return tb_zips[0], nil
}

func GetTbZipRowCount() (count int, err error) {
	rows, err := db.Dtsc.Query("select count(0) Count from tb_zip")
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

func _TbZipRowsToArray(maps []map[string][]byte) ([]TbZip, error) {
	models := make([]TbZip, len(maps))
	var err error
	for index, obj := range maps {
		model := TbZip{}
		model.Id, err = strconv.Atoi(string(obj["id"]))
		if err != nil {
			return nil, errors.New("parse Id error: " + err.Error())
		}
		model.JobId, err = strconv.Atoi(string(obj["job_id"]))
		if err != nil {
			return nil, errors.New("parse JobId error: " + err.Error())
		}
		model.Version, err = strconv.Atoi(string(obj["version"]))
		if err != nil {
			return nil, errors.New("parse Version error: " + err.Error())
		}
		model.Zipfilename = string(obj["zipfilename"])
		model.Zipfile = []byte(obj["zipfile"])
		model.Time, err = strconv.ParseInt(string(obj["time"]), 10, 64)
		if err != nil {
			return nil, errors.New("parse Time error: " + err.Error())
		}
		models[index] = model
	}
	return models, err
}
