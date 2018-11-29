package database

import (
	"errors"
	"strconv"
	db "taskweb/database"
	"strings"
)

type TbCommand struct {
	Id	int
	Jobid	int
	Commandtype	int
	Success	bool
	Time	int64
}

func ExistTbCommand(id int) (bool, error) {
	rows, err := db.Dtsc.Query("select count(0) Count from tb_command where id=?", id)
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

func InsertTbCommand(tb_command TbCommand) (int64, error) {
	result, err := db.Dtsc.Exec("insert into tb_command(jobid,commandtype,success,time) values(?,?,?,?)", tb_command.Jobid,tb_command.Commandtype,tb_command.Success,tb_command.Time)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

func UpdateTbCommand(tb_command TbCommand) (bool, error) {
	result, err := db.Dtsc.Exec("update tb_command set jobid=?, commandtype=?, success=?, time=? where id=?", tb_command.Jobid, tb_command.Commandtype, tb_command.Success, tb_command.Time, tb_command.Id)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

func GetTbCommand(id int) (tb_command TbCommand, err error) {
	rows, err := db.Dtsc.Query("select id, jobid, commandtype, success, time from tb_command where id=?", id)
	if err != nil {
		return tb_command, err
	}
	if len(rows) <= 0 {
		return tb_command, nil
	}
	tb_commands, err := _TbCommandRowsToArray(rows)
	if err != nil {
		return tb_command, err
	}
	return tb_commands[0], nil
}

func GetTbCommandRowCount() (count int, err error) {
	rows, err := db.Dtsc.Query("select count(0) Count from tb_command")
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

func _TbCommandRowsToArray(maps []map[string][]byte) ([]TbCommand, error) {
	models := make([]TbCommand, len(maps))
	var err error
	for index, obj := range maps {
		model := TbCommand{}
		model.Id, err = strconv.Atoi(string(obj["id"]))
		if err != nil {
			return nil, errors.New("parse Id error: " + err.Error())
		}
		model.Jobid, err = strconv.Atoi(string(obj["jobid"]))
		if err != nil {
			return nil, errors.New("parse Jobid error: " + err.Error())
		}
		model.Commandtype, err = strconv.Atoi(string(obj["commandtype"]))
		if err != nil {
			return nil, errors.New("parse Commandtype error: " + err.Error())
		}
		if strings.ToLower(string(obj["success"])) == "true" || obj["success"][0] == 1 {
			model.Success = true
		} else {
			model.Success = false
		}
		model.Time, err = strconv.ParseInt(string(obj["time"]), 10, 64)
		if err != nil {
			return nil, errors.New("parse Time error: " + err.Error())
		}
		models[index] = model
	}
	return models, err
}
func GetCommands() (cmds []TbCommand, err error) {
	rows, err := db.Dtsc.Query("select * from tb_command ")
	if err != nil {
		return cmds, err
	}
	if len(rows) <= 0 {
		return cmds, nil
	}
	return _TbCommandRowsToArray(rows)
}
