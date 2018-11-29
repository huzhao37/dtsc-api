package database

import (
	"errors"
	"strconv"
	db "taskweb/database"
)

type TbCategory struct {
	Id	int
	Name	string
	Createtime	int64
	Remark	string
}

func ExistTbCategory(id int) (bool, error) {
	rows, err := db.Dtsc.Query("select count(0) Count from tb_category where id=?", id)
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

func InsertTbCategory(tb_category TbCategory) (int64, error) {
	result, err := db.Dtsc.Exec("insert into tb_category(name,createtime,remark) values(?,?,?)", tb_category.Name,tb_category.Createtime,tb_category.Remark)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

func UpdateTbCategory(tb_category TbCategory) (bool, error) {
	result, err := db.Dtsc.Exec("update tb_category set name=?, createtime=?, remark=? where id=?", tb_category.Name, tb_category.Createtime, tb_category.Remark, tb_category.Id)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

func GetTbCategory(id int) (tb_category TbCategory, err error) {
	rows, err := db.Dtsc.Query("select id, name, createtime, remark from tb_category where id=?", id)
	if err != nil {
		return tb_category, err
	}
	if len(rows) <= 0 {
		return tb_category, nil
	}
	tb_categorys, err := _TbCategoryRowsToArray(rows)
	if err != nil {
		return tb_category, err
	}
	return tb_categorys[0], nil
}

func GetTbCategorys() (tb_categorys []TbCategory, err error) {
	rows, err := db.Dtsc.Query("select id, name, createtime, remark from tb_category ")
	if err != nil {
		return tb_categorys, err
	}
	if len(rows) <= 0 {
		return tb_categorys, nil
	}
	return _TbCategoryRowsToArray(rows)
}

func GetTbCategoryRowCount() (count int, err error) {
	rows, err := db.Dtsc.Query("select count(0) Count from tb_category")
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

func _TbCategoryRowsToArray(maps []map[string][]byte) ([]TbCategory, error) {
	models := make([]TbCategory, len(maps))
	var err error
	for index, obj := range maps {
		model := TbCategory{}
		model.Id, err = strconv.Atoi(string(obj["id"]))
		if err != nil {
			return nil, errors.New("parse Id error: " + err.Error())
		}
		model.Name = string(obj["name"])
		model.Createtime, err = strconv.ParseInt(string(obj["createtime"]), 10, 64)
		if err != nil {
			return nil, errors.New("parse Createtime error: " + err.Error())
		}
		model.Remark = string(obj["remark"])
		models[index] = model
	}
	return models, err
}

func DelTbCategory(categoryid int) (ra int64, err error) {
	exist, err := db.Dtsc.Query("select * from tb_job where category_id=? ",categoryid)
	if len(exist)>0{
		return 0,nil
	}
	rows, err := db.Dtsc.Exec("delete from tb_category where id=?",categoryid)
	if err != nil {
		return 0, err
	}
	return rows.RowsAffected()
}
