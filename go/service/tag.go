package service

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"hidakkathon/model"
	"hidakkathon/repository"
)

type Tag struct {
	DB *sqlx.DB
}

func NewTagService(db *sqlx.DB) *Tag {
	return &Tag{db}
}

func (t *Tag)IsExists(tid int)bool{
	tag,_:=repository.GetTagInfo(t.DB,tid)
	if tag.Id==0{
		fmt.Println("not exists",tid)
		return false
	}
	fmt.Println("exists",tag.Id)
	return true
}

func (t *Tag)IsExistsByName(tname string)bool{
	tag,_:=repository.GetTagInfoByName(t.DB,tname)
	if tag.Id==0{
		fmt.Println("not exists",tname)
		return false
	}
	fmt.Println("exists",tag.Name)
	return true
}

func (t *Tag)IsExistsTypeByName(tname string)bool{
	types,_:=repository.GetTypeInfoByName(t.DB,tname)
	if types.Id==0{
		fmt.Println("not exists",tname)
		return false
	}
	fmt.Println("exists",types.Name)
	return true
}

func (t *Tag)IsExistsTypeById(tid int)bool {
	types, _ := repository.GetTypeInfoById(t.DB, tid)
	if types.Id == 0 {
		fmt.Println("not exists", tid)
		return false
	}
	fmt.Println("exists", types.Name)
	return true
}

func (t *Tag)GetAll()(*[]model.Tag,error){
	tags, err := repository.GetTags(t.DB)
	if err != nil {
		return nil, err
	}
	return tags,nil
}

func (t *Tag)GetAllTargetUserType()(*[]model.TargetUserType,error){
	types, err := repository.GetTargetUserTypes(t.DB)
	if err != nil {
		return nil, err
	}
	return types,nil
}
