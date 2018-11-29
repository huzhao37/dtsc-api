package database

import (
	"errors"
	"strconv"
	db "taskweb/database"
)

type TbUser struct {
	Id	int
	Name	string
	Pwd	string
	Email	string
	Createtime	int64
	Remark	string
}

func ExistTbUser(id int) (bool, error) {
	rows, err := db.Dtsc.Query("select count(0) Count from tb_user where id=?", id)
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

func InsertTbUser(tb_user TbUser) (int64, error) {
	result, err := db.Dtsc.Exec("insert into tb_user(name,pwd,email,createtime,remark) values(?,?,?,?,?)", tb_user.Name,tb_user.Pwd,tb_user.Email,tb_user.Createtime,tb_user.Remark)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

func UpdateTbUser(tb_user TbUser) (bool, error) {
	result, err := db.Dtsc.Exec("update tb_user set name=?, pwd=?, email=?, createtime=?, remark=? where id=?", tb_user.Name, tb_user.Pwd, tb_user.Email, tb_user.Createtime, tb_user.Remark, tb_user.Id)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

func GetTbUser(id int) (tb_user TbUser, err error) {
	rows, err := db.Dtsc.Query("select id, name, pwd, email, createtime, remark from tb_user where id=?", id)
	if err != nil {
		return tb_user, err
	}
	if len(rows) <= 0 {
		return tb_user, nil
	}
	tb_users, err := _TbUserRowsToArray(rows)
	if err != nil {
		return tb_user, err
	}
	return tb_users[0], nil
}

func GetTbUserRowCount() (count int, err error) {
	rows, err := db.Dtsc.Query("select count(0) Count from tb_user")
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

func _TbUserRowsToArray(maps []map[string][]byte) ([]TbUser, error) {
	models := make([]TbUser, len(maps))
	var err error
	for index, obj := range maps {
		model := TbUser{}
		model.Id, err = strconv.Atoi(string(obj["id"]))
		if err != nil {
			return nil, errors.New("parse Id error: " + err.Error())
		}
		model.Name = string(obj["name"])
		model.Pwd = string(obj["pwd"])
		model.Email = string(obj["email"])
		model.Createtime, err = strconv.ParseInt(string(obj["createtime"]), 10, 64)
		if err != nil {
			return nil, errors.New("parse Createtime error: " + err.Error())
		}
		model.Remark = string(obj["remark"])
		models[index] = model
	}
	return models, err
}
func UserLogin(name string,pwd string) (tb_user TbUser, err error ) {
	rows, err := db.Dtsc.Query("select * from tb_user where name=? and pwd=?", name,pwd)
	if err != nil {
		return tb_user, err
	}
	if len(rows) <= 0 {
		return tb_user, nil
	}
	tb_users, err := _TbUserRowsToArray(rows)
	if err != nil {
		return tb_user, err
	}
	return tb_users[0], nil
}
