package model

type ReqLogin struct {
	UserId   string `json:"user_id"`
	Password string `json:"password"`
}

type User struct {
	UserId      string `json:"user_id"`
	UserName    string `json:"user_name"`
	Password    string `json:"password"`
	UserComment string `json:"user_comment"`
	IsAdmin     int    `json:"is_admin"`
}

type ResLoginUser struct {
	IsAdmin  bool   `json:"is_admin" db:"is_admin"`
	UserId   string `json:"user_id" db:"user_id"`
	UserName string `json:"user_name" db:"user_name"`
}

type UserInfo struct {
	Id      string `json:"user_id" db:"user_id"`
	Name    string `json:"user_name" db:"user_name"`
	IsAdmin int   `json:"is_admin" db:"is_admin"`
}

type RespSelfInfo struct{
	Name string `json:"user_name" db:"user_name"`
	Comment string `json:"user_commment" db:"user_comment"`
	EventInfoList []EventInfo `json:"event_info_list"`
}

type SelfUser struct{
	Name string `json:"user_name" db:"user_name"`
	Comment string `json:"user_commment" db:"user_comment"`
}

type RespUsersList struct{
	List []UserForGetUsers `json:"user_list"`
}

type UserForGetUsers struct{
	UserId      string `json:"user_id" db:"user_id"`
	UserName    string `json:"user_name" db:"user_name"`
	UserComment string `json:"user_comment" db:"user_comment"`
	IsAdmin     bool    `json:"is_admin" db:"is_admin"`
}

type RespOther struct{
	Name string `json:"user_name"`
	Comment string `json:"user_comment"`
	EventInfoList []EventInfo `json:"event_info_list"`
}