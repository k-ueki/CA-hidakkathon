package model

type Tag struct{
	Id int `json:"tag_id" db:"tag_id"`
	Name string `json:"tag_name" db:"tag_name"`
}

type TagIds struct {
	TagId int `json:"tag_id" db:"tag_id"`
}

type RespTagList struct{
	Tags []Tag `json:"tag_list"`
}
