package repository

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"hidakkathon/model"
)

func GetTagInfo(db *sqlx.DB,tid int)(*model.Tag,error){
	tag:=model.Tag{}
	rows,err:=db.Queryx(`SELECT tag_id, tag_name FROM m_event_tag WHERE tag_id = ?`,tid)
	if err!=nil{
		fmt.Println("SERVER ERROR : ",err)
		return nil,err
	}
	for rows.Next(){
		err:=rows.StructScan(&tag)
		if err!=nil{
			fmt.Println(err)
			return nil,err
		}
	}
	return &tag, nil
}

func GetTagInfoByName(db *sqlx.DB,tname string)(*model.Tag,error) {
	tag := model.Tag{}
	rows, err := db.Queryx(`SELECT tag_id, tag_name FROM m_event_tag WHERE tag_name = ?`, tname)
	if err != nil {
		fmt.Println("SERVER ERROR : ", err)
		return nil, err
	}
	for rows.Next() {
		err := rows.StructScan(&tag)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return &tag, nil
}

func GetTypeInfoByName(db *sqlx.DB,tname string)(*model.TargetUserType,error) {
	types := model.TargetUserType{}
	rows, err := db.Queryx(`SELECT target_user_type_id, target_user_type_name FROM m_target_user_type WHERE target_user_type_name = ?`, tname)
	if err != nil {
		fmt.Println("SERVER ERROR : ", err)
		return nil, err
	}
	for rows.Next() {
		err := rows.StructScan(&types)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return &types, nil
}

func GetTypeInfoById(db *sqlx.DB,tid int)(*model.TargetUserType,error) {
	types := model.TargetUserType{}
	rows, err := db.Queryx(`SELECT target_user_type_id, target_user_type_name FROM m_target_user_type WHERE target_user_type_id = ?`, tid)
	if err != nil {
		fmt.Println("SERVER ERROR : ", err)
		return nil, err
	}
	for rows.Next() {
		err := rows.StructScan(&types)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return &types, nil
}

func DeleteTag(db *sqlx.DB,tid int)(sql.Result,error){
	stmt, err := db.Prepare(`
DELETE FROM m_event_tag WHERE tag_id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(tid)
}

func GetAll(db *sqlx.DB)(*[]model.Tag,error){
	tags := []model.Tag{}
	if err:=db.Select(&tags,`
SELECT 
	*
FROM m_event_tag
	`);err!=nil{
		return nil,err
	}
	return  &tags,nil
}

func CreateTag(db *sqlx.DB,tname string)(sql.Result,error){
	stmt, err := db.Prepare(`
INSERT INTO m_event_tag (tag_name) VALUES (?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(tname)

}

func CreateTargetUserType(db *sqlx.DB,tname string,colorcode string)(sql.Result,error) {
	stmt, err := db.Prepare(`
INSERT INTO m_target_user_type (target_user_type_name,color_code) VALUES (?,?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(tname, colorcode)
}

func UpdateTargetUserType(db *sqlx.DB,tid int,tname string,color string)(sql.Result,error){
	stmt, err := db.Prepare(`
UPDATE m_target_user_type SET target_user_type_name = ?, color_code = ? WHERE target_user_type_id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(tname, color,tid)
}

func UpdateUserFull(db *sqlx.DB,id string,name string,comment string,pass string,is bool)(sql.Result,error){
	stmt, err := db.Prepare(`
UPDATE i_user SET user_name = ?, user_comment = ?, password = ?, is_admin = ? WHERE user_id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(name,comment,pass,is,id)
}

func UpdateUserWithoutPass(db *sqlx.DB,id string,name string,comment string,is bool)(sql.Result,error){
	stmt, err := db.Prepare(`
UPDATE i_user SET user_name = ?, user_comment = ?, is_admin = ? WHERE user_id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(name,comment,is,id)
}

func UpdateUserSelf(db *sqlx.DB,id string,name string,comment string)(sql.Result,error){
	fmt.Println(id,name,comment)
	stmt, err := db.Prepare(`
UPDATE i_user SET user_name = ?, user_comment = ? WHERE user_id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(name,comment,id)
}
