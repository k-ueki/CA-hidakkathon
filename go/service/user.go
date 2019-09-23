package service

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"hidakkathon/model"
	"hidakkathon/repository"
)

type User struct {
	DB *sqlx.DB
}

func NewUserService(db *sqlx.DB) *User {
	return &User{db}
}

func (u *User) Login(req model.ReqLogin) (*model.ResLoginUser, error) {
	usr, _ := repository.Login(u.DB, req)

	return usr, nil
}

func (u *User)IsAdmin(uid string) (*model.UserInfo,bool){
	usr,_:=repository.GetUserInfo(u.DB,uid)
	if usr.IsAdmin==0{
		return usr,false
	}
	return nil,true
}

func (u *User)IsExists(uid string)bool{
	user,_:=repository.GetUserInfo(u.DB,uid)
	fmt.Println("isExists user ",user.Id)
	if user.Id==""{
		fmt.Println("not exists",user.Id)
		return false
	}
	fmt.Println("exists",user.Id)
	return true
}

func (u *User)IsExistsTargetUseType(id int)bool{
	usertype,_:=repository.GetUserTargetUserType(u.DB,id)
	if usertype.Id==0{
		fmt.Println("not exists target_user_type",usertype.Id)
		return false
	}
	fmt.Println("exists target_user_type",usertype.Id)
	return true
}

func (u *User)GetSelfInfo(uid string)(string,string,error){
	info,err:=repository.GetSelfInfo(u.DB,uid)
	if err!=nil{
		return "","",err
	}
	return info.Name,info.Comment,nil
}

func (u *User)GetAll()(*[]model.UserForGetUsers,error){
	users,err:=repository.GetAllUsers(u.DB)
	if err!=nil{
		return nil,err
	}
	return users,nil
}
