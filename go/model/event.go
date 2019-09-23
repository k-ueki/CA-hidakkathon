package model

type Event struct{
	CreatedUserId string `json:"created_user_id" db:"created_user_id"`
}
//<<<<<<< HEAD
//// object と　sulice の違いがわからん
//type EventDetail struct {
//	Id                  int        `json:"event_id" db:"event_id"`
//	Name                string     `json:"event_name" db:"event_name"`
//	StartDate           string     `json:"start_date" db:"start_date"`
//	EndDate             string     `json:"end_date" db:"end_date"`
//	Location            string     `json:"location" db:"location"`
//	TargetUserType      []int      `json:"target_user_type"`
//	TargetUser          string     `json:"target_user"`
//	RegisteredUser      UserInfo   `json:"registered_user"`
//	ParticipantLimitNum int        `json:"participant_limit_num"`
//	DetailComment       string     `json:"detail_comment"`
//	TagList             []int      `json:"tag_list"`
//	AttendUserList      []UserInfo `json:"attend_user_list"`
//	IsAuthor            bool       `json:"is_author"`
//	IsAttend            bool       `json:"is_attend"`
//}
//
//type Basic struct {
//	Id                  int    `json:"event_id" db:"event_id"`
//	Name                string `json:"event_name" db:"event_name"`
//	StartDate           string `json:"start_date" db:"start_date"`
//	EndDate             string `json:"end_date" db:"end_date"`
//	Location            string `json:"location" db:"location"`
//	TargetUser          string `json:"target_user" db:"target_user"`
//	ParticipantLimitNum int    `json:"participant_limit_num" db:"participant_limit_num"`
//	DetailComment       string `json:"detail_comment" db:"detail_comment"`
//	AuthId              string `json:"" db:"created_user_id"`
//}
//
//type Types struct {
//	TargetUserTypeId int `json:"target_user_type_id" db:"target_user_type_id"`
//=======
type EventsResp struct {
	EventInfoList      []EventInfo      `json:"event_info_list"`
	TargetUserTypeList []TargetUserType `json:"target_user_type_list"`
	TagList            []Tag            `json:"tag_list" `
	UserId             string           `json:"user_id" db:"user_id"`
	UserName           string           `json:"user_name" db:"user_name"`
	IsAdmin            bool             `json:"is_admin" db:"is_admin"`
}

type EventInfo struct {
	Id        int    `json:"event_id" db:"event_id"`
	Name      string `json:"event_name" db:"event_name"`
	StartDate string `json:"start_date" db:"start_date"`
	EndDate   string `json:"end_date" db:"end_date"`
	ColorCode string `json:"color_code" db:"color_code"`
}

type TargetUserType struct {
	Id        int    `json:"target_user_type_id" db:"target_user_type_id"`
	Name      string `json:"target_user_type_name" db:"target_user_type_name"`
	ColorCode string `json:"color_code" db:"color_code"`
}

type Date struct {
	Start  string
	Finish string
}

type Img struct{
	ImgBinary []byte `db:"img_binary"`
}

type RespTargetUserTypeList struct{
	List []TargetUserType `json:"target_user_type_list"`
}