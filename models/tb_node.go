package database

import (
	"errors"
	"strconv"
	db "taskweb/database"
)

type TbNode struct {
	Id	int
	Name	string
	Ip	string
	Createtime	int64
	Remark	string
}

func ExistTbNode(id int) (bool, error) {
	rows, err := db.Dtsc.Query("select count(0) Count from tb_node where id=?", id)
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

func InsertTbNode(tb_node TbNode) (int64, error) {
	result, err := db.Dtsc.Exec("insert into tb_node(name,ip,createtime,remark) values(?,?,?,?)", tb_node.Name,tb_node.Ip,tb_node.Createtime,tb_node.Remark)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

func UpdateTbNode(tb_node TbNode) (bool, error) {
	result, err := db.Dtsc.Exec("update tb_node set name=?, ip=?, createtime=?, remark=? where id=?", tb_node.Name, tb_node.Ip, tb_node.Createtime, tb_node.Remark, tb_node.Id)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

func GetTbNode(id int) (tb_node TbNode, err error) {
	rows, err := db.Dtsc.Query("select id, name, ip, createtime, remark from tb_node where id=?", id)
	if err != nil {
		return tb_node, err
	}
	if len(rows) <= 0 {
		return tb_node, nil
	}
	tb_nodes, err := _TbNodeRowsToArray(rows)
	if err != nil {
		return tb_node, err
	}
	return tb_nodes[0], nil
}
func GetTbNodes() (tb_nodes []TbNode, err error) {
	rows, err := db.Dtsc.Query("select * from tb_node ")
	if err != nil {
		return tb_nodes, err
	}
	if len(rows) <= 0 {
		return tb_nodes, nil
	}
	return _TbNodeRowsToArray(rows)
}
func GetTbNodeRowCount() (count int, err error) {
	rows, err := db.Dtsc.Query("select count(0) Count from tb_node")
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

func _TbNodeRowsToArray(maps []map[string][]byte) ([]TbNode, error) {
	models := make([]TbNode, len(maps))
	var err error
	for index, obj := range maps {
		model := TbNode{}
		model.Id, err = strconv.Atoi(string(obj["id"]))
		if err != nil {
			return nil, errors.New("parse Id error: " + err.Error())
		}
		model.Name = string(obj["name"])
		model.Ip = string(obj["ip"])
		model.Createtime, err = strconv.ParseInt(string(obj["createtime"]), 10, 64)
		if err != nil {
			return nil, errors.New("parse Createtime error: " + err.Error())
		}
		model.Remark = string(obj["remark"])
		models[index] = model
	}
	return models, err
}

func DelTbNode(nodeid int) (ra int64, err error) {
	exist, err := db.Dtsc.Query("select * from tb_job where node_id=? ",nodeid)
	if len(exist)>0{
		return 0,nil
	}
	rows, err := db.Dtsc.Exec("delete from tb_node where id=?",nodeid)
	if err != nil {
		return 0, err
	}
	return rows.RowsAffected()
}
