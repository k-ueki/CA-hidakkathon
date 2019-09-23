package repository

import (
	"database/sql"
	"fmt"
	"hidakkathon/model"

	"github.com/jmoiron/sqlx"
)

func Login(db *sqlx.DB, req model.ReqLogin) (*model.ResLoginUser, error) {
	usr := model.ResLoginUser{}
	if err := db.Get(&usr, `SELECT user_id, user_name, is_admin FROM i_user WHERE user_id =? AND password=?`, req.UserId, req.Password); err != nil {
		fmt.Println("err2", err)
		return nil, err
	}
	return &usr, nil
}

func GetAllUsers(db *sqlx.DB)(*[]model.UserForGetUsers,error){
	users := []model.UserForGetUsers{}
	if err := db.Select(&users, `
SELECT 
	user_id, user_name, user_comment, is_admin
FROM i_user
	`); err != nil {
		return nil, err
	}
	return &users, nil
}

func GetUserInfo(db *sqlx.DB,uid string)(*model.UserInfo,error){
	info:=model.UserInfo{}
	rows,err:=db.Queryx(`SELECT user_id, user_name, is_admin FROM i_user WHERE user_id = ?`,uid)
	if err!=nil{
		fmt.Println("SERVER ERROR : ",err)
		return nil,err
	}
	for rows.Next(){
		err:=rows.StructScan(&info)
		if err!=nil{
			fmt.Println(err)
			return nil,err
		}
	}
	return &info, nil
}

func GetSelfInfo(db *sqlx.DB,uid string)(*model.SelfUser,error){
	resp:=model.SelfUser{}
	rows,err:=db.Queryx(`SELECT user_name, user_comment FROM i_user WHERE user_id = ?`,uid)
	if err!=nil{
		fmt.Println("SERVER ERROR : ",err)
		return nil,err
	}
	for rows.Next(){
		err:=rows.StructScan(&resp)
		if err!=nil{
			fmt.Println(err)
			return nil,err
		}
	}
	return &resp, nil

}

func GetUserTargetUserType(db *sqlx.DB,id int)(*model.TargetUserType,error){
	info:=model.TargetUserType{}
	rows,err:=db.Queryx(`SELECT target_user_type_id, target_user_type_name, color_code FROM m_target_user_type WHERE target_user_type_id = ?`,id)
	if err!=nil{
		fmt.Println("SERVER ERROR : ",err)
		return nil,err
	}
	for rows.Next(){
		err:=rows.StructScan(&info)
		if err!=nil{
			fmt.Println(err)
			return nil,err
		}
	}
	return &info, nil
}

func DeleteUser(db *sqlx.DB,uid string)(sql.Result,error){
	stmt, err := db.Prepare(`
DELETE FROM i_user WHERE user_id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(uid)
}

func DeleteTargetUserType(db *sqlx.DB,id int)(sql.Result,error){
	stmt, err := db.Prepare(`
DELETE FROM m_target_user_type WHERE target_user_type_id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(id)
}

func CreateUser(db *sqlx.DB,id,name,comment,pass string,is bool)(sql.Result,error){
	stmt, err := db.Prepare(`
INSERT INTO i_user (user_id, user_name, password, user_comment, is_admin) VALUES (?,?,?,?,?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(id,name,pass,comment,is)
}
