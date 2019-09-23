package service

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"hidakkathon/model"
	"hidakkathon/repository"
	"net/http"
	"strconv"
)

//import (
//	"fmt"
//	"github.com/jmoiron/sqlx"
//	"hidakkathon/model"
//	"hidakkathon/repository"
//	"net/http"
//	"strconv"
//)

//import (
//	"fmt"
//	"hidakkathon/model"
//	"hidakkathon/repository"
//<<<<<<< HEAD
//	"net/http"
//=======
//>>>>>>> master
//	"strconv"
//
//	"github.com/jmoiron/sqlx"
//)

type Event struct {
	DB *sqlx.DB
}

func NewEventService(db *sqlx.DB) *Event {
	return &Event{db}
}

//<<<<<<< HEAD
//func (e *Event) GetEventDetailById(eid string, r *http.Request) (*model.EventDetail, error) {
//	co, err := r.Cookie("user_id")
//	fmt.Println("A", co)
//	if err != nil {
//		detail.IsAuthor = false
//		detail.IsAttend = false
//	}
//	fmt.Println("B")
//	// user_id := co.Value
//	user_id := "hoge"
//
//	eventId, _ := strconv.Atoi(eid)
//
//	respBasic, err := repository.GetBasicById(e.DB, eventId)
//	if err != nil {
//		return nil, err
//	}
//	taglist, err := repository.GetTagList(e.DB, eventId)
//=======
func (e *Event) GetAllEventsTargetYM(events *model.EventsResp, year string, month string) (*model.EventsResp, error) {
	tyear,_:=strconv.Atoi(year)
	tmonth, _ := strconv.Atoi(month)
	date := model.Date{
		Start:  fmt.Sprintf("%d-%d-01", tyear, tmonth),
		Finish: fmt.Sprintf("%d-%d-01", tyear, tmonth+1),
	}
	if tmonth==12{
		date.Finish=fmt.Sprintf("%d-%d-01", tyear+1, 1)
	}
	fmt.Println(date)

	eventInfoLists, err := repository.GetEventInfoLists(e.DB, date)
	if err != nil {
		return nil, err
	}

	targetUserTypes, err := repository.GetTargetUserTypes(e.DB)
	if err != nil {
		return nil, err
	}

	tags, err := repository.GetTags(e.DB)
	if err != nil {
		return nil, err
	}

	events.EventInfoList = *eventInfoLists
	events.TargetUserTypeList = *targetUserTypes
	events.TagList = *tags

	return events, nil
}

func (e *Event)Delete(eid int,uid string)int{
	event,_:=repository.GetEventInfo(e.DB,eid)

	if event.CreatedUserId != uid {
		fmt.Println(":delete: not author")
		return http.StatusForbidden
	}
	fmt.Println(":delete: Author")

	_,err := repository.Delete(e.DB,eid)
	if err != nil {
		return http.StatusInternalServerError
	}
	return http.StatusOK
}

func (e *Event)IsExists(eid int)bool{
	event,_:=repository.GetEventInfo(e.DB,eid)
	fmt.Println("isExists event ",*event,event.CreatedUserId)
	if event.CreatedUserId==""{
		return false
	}
	return true
}

func (e *Event)GetListByUid(uid string)(*model.EventInfo,error){
	list,err:=repository.GetList(e.DB,uid)
	if err!=nil{
		return nil,err
	}
	return list,nil
}

// API4あとはregistered_userその他調整のみ
//=======
//	events.EventInfoList = *eventInfoLists
//	events.TargetUserTypeList = *targetUserTypes
//	events.TagList = *tags
//
//	return events, nil
//}
//>>>>>>> master
