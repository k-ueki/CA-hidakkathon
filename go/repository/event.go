package repository

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"hidakkathon/model"
	"log"
)

func GetEventInfoLists(db *sqlx.DB, date model.Date) (*[]model.EventInfo, error) {
	list := []model.EventInfo{}
	if err := db.Select(&list, `
SELECT 
	i_event.event_id,
	i_event.event_name,
	i_event.start_date,
	i_event.end_date,
	m_target_user_type.color_code
FROM i_event
INNER JOIN i_event_target_user_type 
ON i_event.event_id = i_event_target_user_type.event_id
INNER JOIN m_target_user_type
ON i_event_target_user_type.target_user_type_id = m_target_user_type.target_user_type_id
WHERE i_event.start_date > ? AND i_event.end_date < ?
ORDER BY i_event.event_id ASC
	`, date.Start, date.Finish); err != nil {
		return nil, err
	}
	return &list, nil
}

func GetAllEventInfoListsByUid(db *sqlx.DB,uid string)(*[]model.EventInfo,error){
	list := []model.EventInfo{}
	if err:=db.Select(&list,`
SELECT 
	i_event.event_id,
	i_event.event_name,
	i_event.start_date,
	i_event.end_date,
	m_target_user_type.color_code
FROM i_event
INNER JOIN i_participate_event 
ON i_event.event_id = i_participate_event.event_id
INNER JOIN i_user
ON i_participate_event.user_id = i_user.user_id
INNER JOIN i_event_target_user_type
ON i_event_target_user_type.event_id = i_event.event_id
INNER JOIN m_target_user_type
ON m_target_user_type.target_user_type_id = i_event_target_user_type.target_user_type_id
WHERE i_user.user_id = ?
ORDER BY i_event.event_id ASC
	`,uid);err!=nil{
	return nil,err
	}
	return &list,nil
}

func GetEventListByUid(db *sqlx.DB,uid string)(*[]model.EventInfo,error){
	list := []model.EventInfo{}

//	rows,err:=db.Queryx(`
//SELECT DISTINCT
//	i_event.event_id,
//	i_event.event_name,
//	i_event.start_date,
//	i_event.end_date,
//	m_target_user_type.color_code
//FROM i_event
//INNER JOIN i_participate_event
//ON i_event.event_id = i_participate_event.event_id
//INNER JOIN i_user
//ON i_participate_event.user_id = i_user.user_id
//INNER JOIN i_event_target_user_type
//ON i_event_target_user_type.event_id = i_event.event_id
//INNER JOIN m_target_user_type
//ON m_target_user_type.target_user_type_id = i_event_target_user_type.target_user_type_id
//WHERE i_user.user_id = "abigayle957"
//ORDER BY i_event.event_id ASC
//`,uid)
//	if err!=nil{
//		return nil,err
//	}
//	for rows.Next(){
//		tmp := model.EventInfo{}
//		err:=rows.StructScan(&tmp)
//		if err!=nil{
//			return nil,err
//		}
//		list=append(list,tmp)
//	}

	if err:=db.Select(&list,`
SELECT DISTINCT
	i_event.event_id,
	i_event.event_name,
	i_event.start_date,
	i_event.end_date,
	m_target_user_type.color_code
FROM i_event
INNER JOIN i_participate_event
ON i_event.event_id = i_participate_event.event_id
INNER JOIN i_user
ON i_participate_event.user_id = i_user.user_id
INNER JOIN i_event_target_user_type
ON i_event_target_user_type.event_id = i_event.event_id
INNER JOIN m_target_user_type
ON m_target_user_type.target_user_type_id = i_event_target_user_type.target_user_type_id
WHERE i_user.user_id = ?
ORDER BY i_event.event_id ASC
	`,uid);err!=nil{
		return nil,err
	}
	fmt.Println("リストだよ",list)
	return &list,nil
}
//func GetTargetUserTypes(db *sqlx.DB, date model.Date) (*[]model.TargetUserType, error) {
//	types := []model.TargetUserType{}
//	if err := db.Select(&types, `
//SELECT
//	m_target_user_type.target_user_type_id,
//	m_target_user_type.target_user_type_name,
//	m_target_user_type.color_code
//FROM m_target_user_type
//INNER JOIN i_event_target_user_type
//ON i_event_target_user_type.target_user_type_id = m_target_user_type.target_user_type_id
//INNER JOIN i_event
//ON i_event.event_id = i_event_target_user_type.event_id
//WHERE i_event.start_date > ? AND i_event.end_date < ?
//ORDER BY m_target_user_type.target_user_type_id ASC
//	`, date.Start, date.Finish); err != nil {
//		return nil, err
//	}
//	return &types, nil
//
//}
func GetTargetUserTypes(db *sqlx.DB) (*[]model.TargetUserType, error) {
	types := []model.TargetUserType{}
	if err := db.Select(&types, `
SELECT 
	*
FROM m_target_user_type
ORDER BY target_user_type_id ASC
	`); err != nil {
		return nil, err
	}
	return &types, nil
}

//func GetTags(db *sqlx.DB, date model.Date) (*[]model.Tag, error) { tags := []model.Tag{} if err := db.Select(&tags, `
//SELECT
//	m_event_tag.tag_id,
//	m_event_tag.tag_name
//FROM m_event_tag
//INNER JOIN i_event_tag
//ON i_event_tag.tag_id = m_event_tag.tag_id
//INNER JOIN i_event
//ON i_event.event_id = i_event_tag.event_id
//WHERE i_event.start_date > ? AND i_event.end_date < ?
//ORDER BY m_event_tag.tag_id ASC
//	`, date.Start, date.Finish); err != nil {
//		return nil, err
//	}
//	return &tags, nil
//}

func GetTags(db *sqlx.DB) (*[]model.Tag, error) {
	tags := []model.Tag{}
	if err := db.Select(&tags, `
SELECT
	*
FROM m_event_tag
ORDER BY m_event_tag.tag_id ASC
	`); err != nil {
		return nil, err
	}
	return &tags, nil
}

func GetList(db *sqlx.DB,uid string)(*model.EventInfo,error){
	resp := model.EventInfo{}
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
func GetImage(db *sqlx.DB,event_id int)[]byte{
	img := model.Img{}
	fmt.Println("event_id",event_id)
	rows,err := db.Queryx( `SELECT img_binary FROM i_event_image WHERE event_id = ?`, event_id)
	fmt.Println(err)
	for rows.Next(){
		err := rows.StructScan(&img)
        if err != nil {
            log.Fatalln(err)
        }
	}
	//if err != nil {
	//	fmt.Println("NOT FOUND")
	//	fmt.Println(err)
	//	return nil
	//}
	fmt.Printf("Img %#v\n",img)
	return img.ImgBinary
}

func GetEventInfo(db *sqlx.DB,eid int)(*model.Event,error){
	resp := model.Event{}
	rows,err:=db.Queryx(`SELECT created_user_id FROM i_event WHERE event_id = ?`,eid)
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

func Delete(db *sqlx.DB,eid int)(sql.Result,error){
	stmt, err := db.Prepare(`
DELETE FROM i_event WHERE event_id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(eid)
}
//import (
//<<<<<<< HEAD
//	"fmt"
//=======
//>>>>>>> master
//	"hidakkathon/model"
//
//	"github.com/jmoiron/sqlx"
//)
//
//<<<<<<< HEAD
//func GetBasicById(db *sqlx.DB, eid int) (*model.Basic, error) {
//	basic := model.Basic{}
//	if err := db.Get(&basic, `
//SELECT
//	event_id,
//	event_name,
//	start_date,
//	end_date,
//	location,
//	target_user,
//	participant_limit_num,
//	created_user_id,
//	event_detail AS detail_comment
//FROM i_event
//WHERE event_id = ?
//	`, eid); err != nil {
//		return nil, err
//	}
//	return &basic, nil
//}
//
//func GetTagList(db *sqlx.DB, eid int) (*[]int, error) {
//	tags := []model.Tag{}
//	if err := db.Select(&tags, `
//SELECT
//	tag_id
//FROM  i_event_tag
//WHERE event_id = ?
//	`, eid); err != nil {
//		return nil, err
//	}
//
//	resp := make([]int, len(tags))
//	for i := range tags {
//		resp[i] = tags[i].TagId
//	}
//	return &resp, nil
//}
//
//// 多分改善んおよち
//func GetRegisteredUser(db *sqlx.DB, eid int) (*[]model.UserInfo, error) {
//	rusers := []model.UserInfo{}
//	if err := db.Select(&rusers, `
//SELECT
//	i_user.user_id,
//	i_user.user_name,
//	i_user.is_admin
//FROM i_user
//INNER JOIN i_participate_event
//ON i_participate_event.event_id = ?
//	`, eid); err != nil {
//		return nil, err
//	}
//	fmt.Println(rusers)
//	return &rusers, nil
//}
//
//// func GetTargetUserType(db *sqlx.DB, eid int) (*[]int, error) {
//// }
//
//// func GetAttendUserList(db *sqlx.DB, eid int) (*[]model.UserInfo, error) {
//// 	list := []model.UserInfo{}
//// 	if err := db.Select(&list, `
////
//// 	`, eid); err != nil {
//// 		return nil, err
//// 	}
//// 	return &list, nil
//// }
//
//func GetTargetUserTypes(db *sqlx.DB, eid int) (*[]int, error) {
//	types := []model.Types{}
//	if err := db.Select(&types, `
//SELECT
//	target_user_type_id
//FROM i_event_target_user_type
//WHERE event_id = ?
//ORDER BY target_user_type_id ASC
//	`, eid); err != nil {
//		return nil, err
//	}
//
//	resp := make([]int, len(types))
//	for i := range types {
//		resp[i] = types[i].TargetUserTypeId
//	}
//	return &resp, nil
//}
//
//func IsParticipatedById(db *sqlx.DB, user_id string) (bool, error) {
//	user := []model.UserInfo{}
//	if err := db.Select(&user, `
//SELECT
//	i_user.user_id,
//	i_user.user_name,
//	i_user.is_admin
//FROM i_user
//INNER JOIN i_participate_event
//ON i_participate_event.user_id = ?
//	`, user_id); err != nil {
//		return false, err
//	}
//	if len(user) == 0 {
//		return false, nil
//	}
//	return true, nil
//=======
//func GetEventInfoLists(db *sqlx.DB, date model.Date) (*[]model.EventInfo, error) {
//	list := []model.EventInfo{}
//	if err := db.Select(&list, `
//SELECT
//	i_event.event_id,
//	i_event.event_name,
//	i_event.start_date,
//	i_event.end_date,
//	m_target_user_type.color_code
//FROM i_event
//INNER JOIN i_event_target_user_type
//ON i_event.event_id = i_event_target_user_type.event_id
//INNER JOIN m_target_user_type
//ON i_event_target_user_type.target_user_type_id = m_target_user_type.target_user_type_id
//WHERE i_event.start_date > ? AND i_event.end_date < ?
//	`, date.Start, date.Finish); err != nil {
//		return nil, err
//	}
//	return &list, nil
//}
//
//func GetTargetUserTypes(db *sqlx.DB, date model.Date) (*[]model.TargetUserType, error) {
//	types := []model.TargetUserType{}
//	if err := db.Select(&types, `
//SELECT
//	m_target_user_type.target_user_type_id,
//	m_target_user_type.target_user_type_name,
//	m_target_user_type.color_code
//FROM m_target_user_type
//INNER JOIN i_event_target_user_type
//ON i_event_target_user_type.target_user_type_id = m_target_user_type.target_user_type_id
//INNER JOIN i_event
//ON i_event.event_id = i_event_target_user_type.event_id
//WHERE i_event.start_date > ? AND i_event.end_date < ?
//	`, date.Start, date.Finish); err != nil {
//		return nil, err
//	}
//	return &types, nil
//
//}
//
//func GetTags(db *sqlx.DB, date model.Date) (*[]model.Tag, error) {
//	tags := []model.Tag{}
//	if err := db.Select(&tags, `
//SELECT
//	m_event_tag.tag_id,
//	m_event_tag.tag_name
//FROM m_event_tag
//INNER JOIN i_event_tag
//ON i_event_tag.tag_id = m_event_tag.tag_id
//INNER JOIN i_event
//ON i_event.event_id = i_event_tag.event_id
//WHERE i_event.start_date > ? AND i_event.end_date < ?
//	`, date.Start, date.Finish); err != nil {
//		return nil, err
//	}
//	return &tags, nil
//
//>>>>>>> master
//}
