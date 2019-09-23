package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"hidakkathon/httputil"
	"hidakkathon/model"
	"hidakkathon/repository"
	"hidakkathon/service"
	"net/http"
	"strconv"
	"time"
	// "github.com/gorilla/mux"
)

type DBHandler struct {
	DB *sqlx.DB
}

func (d *DBHandler) Test(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	return http.StatusOK, "OK", nil
}

func (d *DBHandler) Login(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	reqParam := &model.ReqLogin{}
	reqParam.UserId = r.FormValue("user_id")
	reqParam.Password = r.FormValue("password")

	if reqParam.UserId == "" || reqParam.Password == "" {
		return http.StatusUnauthorized, nil, nil
	}

	Services := service.NewUserService(d.DB)
	resp, _ := Services.Login(*reqParam)
	if resp == nil {
		return http.StatusUnauthorized, nil, nil
	}

	cookie := &http.Cookie{
		//Name:  "sugori_rendez_vous.session=user_id",
		Name:  "user_id",
		Value: resp.UserId,
	}

	http.SetCookie(w, cookie)

	return http.StatusOK, resp, nil
}

func (d *DBHandler) Logout(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	fmt.Println("KK")
	cookie, _ := r.Cookie("user_id")
	if cookie == nil {
		fmt.Println("cookie", cookie)
		return http.StatusUnauthorized, nil, nil
	}
	fmt.Println("KKKKKK")

	http.SetCookie(w, &http.Cookie{
		Name:  "user_id",
		Value: "",
	})

	return http.StatusOK, nil, nil
}

func (d *DBHandler) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	events := model.EventsResp{}
	//fmt.Println(events)
	usr, _ := r.Cookie("user_id")
	if usr != nil {
		uid := usr.Value
		userInfo, _ := repository.GetUserInfo(d.DB, uid)
		// cookie, _ = r.Cookie("user_id")
		// cookie_user_id := cookie.Value
		// cookie_user_name := cookie.Value
		// fmt.Println(cookie_is_admin)
		//
		events.UserId = userInfo.Id
		events.UserName = userInfo.Name
		events.IsAdmin = false
		if userInfo.IsAdmin == 1 {
			events.IsAdmin = true
		}
		// if cookie_is_admin == "0" {
		// 	events.IsAdmin = false
		// }
		// fmt.Println(cookie_is_admin)
	}

	tyear := r.FormValue("target_year")
	tmonth := r.FormValue("target_month")

	if tyear == "" {
		fmt.Println("tyear is empty")
		nowYear := time.Now().Year()
		tyear = strconv.Itoa(nowYear)
	}
	if tmonth == "" {
		fmt.Println("tmonth is empty")
		nowMonth := int(time.Now().Month())
		tmonth = strconv.Itoa(nowMonth)
	}

	fmt.Println(tyear,tmonth)

	eventService := service.NewEventService(d.DB)
	resp, err := eventService.GetAllEventsTargetYM(&events, tyear, tmonth)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	//fmt.Println("response top",resp)

	return http.StatusOK, resp, nil
}

func (d *DBHandler) GetEventImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eid, ok := vars["event_id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	eventId, _ := strconv.Atoi(eid)
	fmt.Println("Eventid", eventId)

	image := repository.GetImage(d.DB, eventId)
	fmt.Println("Image controller", image)
	if image == nil {
		fmt.Println("Image Not Found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(image)))
	w.WriteHeader(http.StatusOK)
	w.Write(image)
	return
}

func (d *DBHandler)GetSelf(w http.ResponseWriter,r *http.Request)(int,interface{},error){
	fmt.Println(":get_self:")
	usr, _ := r.Cookie("user_id")
	if usr == nil {
		fmt.Println("not logined")
		return http.StatusUnauthorized, nil, nil
	}
	userId := usr.Value
	fmt.Println("userid", userId)

	resp := model.RespSelfInfo{}

	userService := service.NewUserService(d.DB)
	name,comment,err:=userService.GetSelfInfo(userId)
	if err!=nil{
		fmt.Println(":get_self: failed to getuserinfo")
		return http.StatusInternalServerError,nil,err
	}
	fmt.Println(":get_self: get userinfo")
	fmt.Println(":get_delf:",name,comment)

	list,err:=repository.GetAllEventInfoListsByUid(d.DB,userId)
	if err!=nil{
		fmt.Println(":get_self: failed to get eventlist")
		return http.StatusInternalServerError,nil,err
	}
	fmt.Println(":get_self: get eventlist")
	fmt.Println(":get_self: ",list)

	resp.Name = name
	resp.Comment = comment
	resp.EventInfoList = *list

	fmt.Println(":get_self: fin",resp)


	return http.StatusOK,resp,nil
}

func (d *DBHandler)GetTag(w http.ResponseWriter,r *http.Request)(int,interface{},error){
	fmt.Println(":get_tag: ")
	usr, _ := r.Cookie("user_id")
	if usr == nil {
		fmt.Println("not logined")
		return http.StatusUnauthorized, nil, nil
	}
	userId := usr.Value
	fmt.Println("userid", userId)

	//管理者かどうか
	userService := service.NewUserService(d.DB)
	usrinfo, isAdmin := userService.IsAdmin(userId)
	if !isAdmin {
		if usrinfo.Id == "" {
			return http.StatusBadRequest, nil, nil
		}
		fmt.Println(":get_tag: not admin")
		return http.StatusForbidden, nil, nil
	}
	fmt.Println(":get_tag: admin!")

	tagService := service.NewTagService(d.DB)
	tags,err:=tagService.GetAll()
	if err!=nil{
		return http.StatusInternalServerError,nil,err
	}

	resp:=model.RespTagList{
		Tags: *tags,
	}

	return http.StatusOK,resp,nil
}

func (d *DBHandler)GetTargetUserType(w http.ResponseWriter,r *http.Request)(int,interface{},error){
	fmt.Println(":get_target_user_type: ")
	usr, _ := r.Cookie("user_id")
	if usr == nil {
		fmt.Println("not logined")
		return http.StatusUnauthorized, nil, nil
	}
	userId := usr.Value
	fmt.Println("userid", userId)

	//管理者かどうか
	userService := service.NewUserService(d.DB)
	usrinfo, isAdmin := userService.IsAdmin(userId)
	if !isAdmin {
		if usrinfo.Id == "" {
			return http.StatusBadRequest, nil, nil
		}
		fmt.Println(":get_target_user_type: not admin")
		return http.StatusForbidden, nil, nil
	}
	fmt.Println(":get_target_user_type: admin!")

	tagService := service.NewTagService(d.DB)
	types,err:=tagService.GetAllTargetUserType()
	if err!=nil{
		return http.StatusInternalServerError,nil,err
	}

	resp:=model.RespTargetUserTypeList{
		List: *types,
	}

	return http.StatusOK,resp,nil
}

func (d *DBHandler)GetAllUsers(w http.ResponseWriter,r *http.Request)(int,interface{},error){
	fmt.Println(":get_users:")

	usr, _ := r.Cookie("user_id")
	if usr == nil {
		fmt.Println("not logined")
		return http.StatusUnauthorized, nil, nil
	}
	userId := usr.Value
	fmt.Println("userid", userId)

	//管理者かどうか
	userService := service.NewUserService(d.DB)
	usrinfo, isAdmin := userService.IsAdmin(userId)
	if !isAdmin {
		if usrinfo.Id == "" {
			return http.StatusBadRequest, nil, nil
		}
		fmt.Println(":get_users: not admin")
		return http.StatusForbidden, nil, nil
	}
	fmt.Println(":get_users: admin!")

	users,err:=userService.GetAll()
	if err!=nil{
		return http.StatusInternalServerError,nil,err
	}

	resp := model.RespUsersList{
		List:*users,
	}
	return http.StatusOK,resp,nil
}

func (d *DBHandler)GetUser(w http.ResponseWriter,r *http.Request)(int,interface{},error){
	fmt.Println(":get_user:")

	usr, _ := r.Cookie("user_id")
	if usr == nil {
		fmt.Println("not logined")
		return http.StatusUnauthorized, nil, nil
	}
	userId := usr.Value
	fmt.Println("userid", userId)

	vars := mux.Vars(r)
	tUid, ok := vars["user_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter user_id"}
	}
	fmt.Println("target_userid", tUid)

	//targetuser isExists
	userService := service.NewUserService(d.DB)
	if !userService.IsExists(tUid) {
		fmt.Println(":get_user: not exists user")
		return http.StatusOK, nil, nil
	}
	fmt.Println(":get_user: exists user")



	info,err:=repository.GetSelfInfo(d.DB,tUid)
	if err!=nil{
		return http.StatusInternalServerError,nil,err
	}
	fmt.Println("infooooooooo",info)

	list,err:=repository.GetEventListByUid(d.DB,tUid)
	if err!=nil{
		fmt.Println(":get_user: failed to get eventlist")
		return http.StatusInternalServerError,nil,err
	}
	fmt.Println(":get_user: get eventlist")
	fmt.Println(":get_user: ",list)

	resp := model.RespOther{
		Name:info.Name,
		Comment:info.Comment,
		EventInfoList:*list,
	}
	fmt.Println("resp!!",resp)
	return http.StatusOK,resp,nil
}

func (d *DBHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	fmt.Println("delete")
	vars := mux.Vars(r)
	eid, ok := vars["event_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter event_id"}
	}
	eventId, _ := strconv.Atoi(eid)
	fmt.Println("event_id", eventId)

	eventService := service.NewEventService(d.DB)
	if !eventService.IsExists(eventId) {
		fmt.Println("not exists event")
		return http.StatusBadRequest, nil, nil
	}
	fmt.Println("exists event")

	usr, _ := r.Cookie("user_id")
	if usr == nil {
		fmt.Println("not logined")
		return http.StatusUnauthorized, nil, nil
	}
	userId := usr.Value
	fmt.Println("userid", userId)

	status := eventService.Delete(eventId, userId)

	return status, nil, nil

}

func (d *DBHandler) DeleteUser(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	// cookie取得
	// is_admin判定
	//
	fmt.Println(":delete_user:")
	userService := service.NewUserService(d.DB)

	usr, _ := r.Cookie("user_id")
	if usr == nil {
		fmt.Println(":delete_user: not logined")
		return http.StatusUnauthorized, nil, nil
	}
	userId := usr.Value
	fmt.Println(":delete_user: login userid", userId)

	//管理者かどうか
	usrinfo, isAdmin := userService.IsAdmin(userId)
	if !isAdmin {
		if usrinfo.Id == "" {
			return http.StatusBadRequest, nil, nil
		}
		fmt.Println(":delete_user: not admin")
		return http.StatusForbidden, nil, nil
	}
	fmt.Println(":delete_user: admin!")

	vars := mux.Vars(r)
	delUid, ok := vars["user_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter user_id"}
	}
	fmt.Println(":delete_user: del user_id", delUid)
	//

	if !userService.IsExists(delUid) {
		fmt.Println(":delete_user: not exists user")
		return http.StatusBadRequest, nil, nil
	}
	fmt.Println(":delete_user: exists user")

	_, err := repository.DeleteUser(d.DB, delUid)
	if err != nil {
		fmt.Println(":delete_user: failed to delete")
		return http.StatusInternalServerError, nil, nil
	}
	fmt.Println(":delete_user: deleted")

	return http.StatusOK, nil, nil

}

func (d *DBHandler) DeleteTag(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	fmt.Println(":delete_tag: S")
	userService := service.NewUserService(d.DB)
	tagService := service.NewTagService(d.DB)

	usr, _ := r.Cookie("user_id")
	if usr == nil {
		fmt.Println(":delete_tag: not logined")
		return http.StatusUnauthorized, nil, nil
	}
	userId := usr.Value
	fmt.Println(":delete_tag: login userid", userId)

	//管理者かどうか
	usrinfo, isAdmin := userService.IsAdmin(userId)
	if !isAdmin {
		if usrinfo.Id == "" {
			fmt.Println(":delete_tag: not found user")
			return http.StatusBadRequest, nil, nil
		}
		fmt.Println(":delete_tag: not admin")
		return http.StatusForbidden, nil, nil
	}
	fmt.Println(":delete_tag: admin!")

	vars := mux.Vars(r)
	tmp, ok := vars["tag_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter tag_id"}
	}
	delTid, _ := strconv.Atoi(tmp)
	fmt.Println(":delete_tag: del tag_id", delTid)
	//

	if !tagService.IsExists(delTid) {
		fmt.Println(":delete_tag: not exists tag")
		return http.StatusBadRequest, nil, nil
	}
	fmt.Println(":delete_tag: exists tag")

	_, err := repository.DeleteTag(d.DB, delTid)
	if err != nil {
		fmt.Println(":delete_user: failed to delete")
		return http.StatusInternalServerError, nil, nil
	}
	fmt.Println(":delete_user: deleted")

	return http.StatusOK, nil, nil

}

func (d *DBHandler) DeleteTargetUserType(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	fmt.Println(":delete_usertype: S")
	userService := service.NewUserService(d.DB)
	//tagService := service.NewTagService(d.DB)
	//
	usr, _ := r.Cookie("user_id")
	if usr == nil {
		fmt.Println(":delete_usertype: not logined")
		return http.StatusUnauthorized, nil, nil
	}
	userId := usr.Value
	fmt.Println(":delete_usertype: login userid", userId)

	//管理者かどうか
	usrinfo, isAdmin := userService.IsAdmin(userId)
	if !isAdmin {
		if usrinfo.Id == "" {
			fmt.Println(":delete_usertype: not found user")
			return http.StatusBadRequest, nil, nil
		}
		fmt.Println(":delete_usertype: not admin")
		return http.StatusForbidden, nil, nil
	}
	fmt.Println(":delete_usertype: admin!")

	vars := mux.Vars(r)
	tmp, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter id"}
	}
	delId, _ := strconv.Atoi(tmp)
	fmt.Println(":delete_usertype: del usertype_id", delId)

	if !userService.IsExistsTargetUseType(delId) {
		fmt.Println(":delete_usertype: not exists tag")
		return http.StatusBadRequest, nil, nil
	}
	fmt.Println(":delete_usertype: exists tag")

	_, err := repository.DeleteTargetUserType(d.DB, delId)
	if err != nil {
		fmt.Println(":delete_usertype: failed to delete")
		return http.StatusInternalServerError, nil, nil
	}
	fmt.Println(":delete_usertype: deleted")

	return http.StatusOK, nil, nil
}

func (d *DBHandler)CreateTag(w http.ResponseWriter,r *http.Request)(int,interface{},error){
	fmt.Println(":create_tag:")
	usr, _ := r.Cookie("user_id")
	if usr == nil {
		fmt.Println(":create_tag: not logined")
		return http.StatusUnauthorized, nil, nil
	}
	userId := usr.Value
	fmt.Println(":create_tag: login userid", userId)

	//管理者かどうか
	userService := service.NewUserService(d.DB)
	usrinfo, isAdmin := userService.IsAdmin(userId)
	if !isAdmin {
		if usrinfo.Id == "" {
			fmt.Println(":create_tag: not found user")
			return http.StatusBadRequest, nil, nil
		}
		fmt.Println(":create_tag: not admin")
		return http.StatusForbidden, nil, nil
	}
	fmt.Println(":create_tag: admin!")

	tagName := r.FormValue("tag_name")

	//tag isExists
	tagService := service.NewTagService(d.DB)
	if tagService.IsExistsByName(tagName) {
		fmt.Println(":delete_tag: exists tag")
		return http.StatusBadRequest, nil, nil
	}
	fmt.Println(":delete_tag: not exists tag")

	//register
	_,err:=repository.CreateTag(d.DB,tagName)
	if err!=nil{
		return http.StatusInternalServerError,nil,nil
	}

	return http.StatusOK,nil,nil
}

func (d *DBHandler)CreateTargetUserType(w http.ResponseWriter,r *http.Request)(int,interface{},error){
	fmt.Println(":create_target_usertype:")
	usr, _ := r.Cookie("user_id")
	if usr == nil {
		fmt.Println(":create_target_user_type: not logined")
		return http.StatusUnauthorized, nil, nil
	}
	userId := usr.Value
	fmt.Println(":create_target_user_type: login userid", userId)

	//管理者かどうか
	userService := service.NewUserService(d.DB)
	usrinfo, isAdmin := userService.IsAdmin(userId)
	if !isAdmin {
		if usrinfo.Id == "" {
			fmt.Println(":create_target_user_type: not found user")
			return http.StatusBadRequest, nil, nil
		}
		fmt.Println(":create_target_user_type: not admin")
		return http.StatusForbidden, nil, nil
	}
	fmt.Println(":create_target_user_type: admin!")

	targetUserTypeName := r.FormValue("target_user_type")
	colorCode := r.FormValue("color_code")

	//target_user_type isExists
	tagService := service.NewTagService(d.DB)
	if tagService.IsExistsTypeByName(targetUserTypeName) {
		fmt.Println(":delete_tag: exists tag")
		return http.StatusBadRequest, nil, nil
	}
	fmt.Println(":delete_tag: not exists tag")

	//register
	_,err:=repository.CreateTargetUserType(d.DB,targetUserTypeName,colorCode)
	if err!=nil{
		return http.StatusInternalServerError,nil,nil
	}

	return http.StatusOK,nil,nil
}

func (d *DBHandler)CreateUser(w http.ResponseWriter,r *http.Request)(int,interface{},error) {
	fmt.Println(":create_user:")
	usr, _ := r.Cookie("user_id")
	if usr == nil {
		fmt.Println(":create_user: not logined")
		return http.StatusUnauthorized, nil, nil
	}
	userId := usr.Value
	fmt.Println(":create_user: login userid", userId)

	//管理者かどうか
	userService := service.NewUserService(d.DB)
	usrinfo, isAdmin := userService.IsAdmin(userId)
	if !isAdmin {
		if usrinfo.Id == "" {
			fmt.Println(":create_user: not found user")
			return http.StatusBadRequest, nil, nil
		}
		fmt.Println(":create_user: not admin")
		return http.StatusForbidden, nil, nil
	}
	fmt.Println(":create_user: admin!")

	uId := r.FormValue("user_id")
	userName := r.FormValue("user_name")
	userComment := r.FormValue("user_comment")
	pass := r.FormValue("password")
	tmp := r.FormValue("is_admin")

	isa,_:=strconv.ParseBool(tmp)

	//register
	_, err := repository.CreateUser(d.DB,uId,userName,userComment,pass,isa)
	if err != nil {
		fmt.Println(":create_user:",err)
		return http.StatusInternalServerError, nil, nil
	}
	fmt.Println(":create_user: created!")

	return http.StatusOK, nil, nil
}

func (d *DBHandler)UpdateTargetUserType(w http.ResponseWriter,r *http.Request)(int,interface{},error){
	fmt.Println(":update_target_usertype:")
	usr, _ := r.Cookie("user_id")
	if usr == nil {
		fmt.Println(":update_target_user_type: not logined")
		return http.StatusUnauthorized, nil, nil
	}
	userId := usr.Value
	fmt.Println(":update_target_user_type: login userid", userId)

	//管理者かどうか
	userService := service.NewUserService(d.DB)
	usrinfo, isAdmin := userService.IsAdmin(userId)
	if !isAdmin {
		if usrinfo.Id == "" {
			fmt.Println(":update_target_user_type: not found user")
			return http.StatusBadRequest, nil, nil
		}
		fmt.Println(":update_target_user_type: not admin")
		return http.StatusForbidden, nil, nil
	}
	fmt.Println(":update_target_user_type: admin!")

	tmpId := r.FormValue("target_user_type_id")
	targetUserTypeName := r.FormValue("target_user_type_name")
	colorCode := r.FormValue("color_code")

	targetUserTypeId,_:=strconv.Atoi(tmpId)

	tagService := service.NewTagService(d.DB)
	if !tagService.IsExistsTypeById(targetUserTypeId) {
		fmt.Println(":delete_type: not exists tag")
		return http.StatusBadRequest, nil, nil
	}
	fmt.Println(":delete_type: exists tag")

	//update
	_,err:=repository.UpdateTargetUserType(d.DB,targetUserTypeId,targetUserTypeName,colorCode)
	if err!=nil{
		return http.StatusInternalServerError,nil,nil
	}

	return http.StatusOK,nil,nil
}

func (d *DBHandler)UpdateUser(w http.ResponseWriter,r *http.Request)(int,interface{},error){
	fmt.Println(":update_user:")
	usr, _ := r.Cookie("user_id")
	if usr == nil {
		fmt.Println(":update_user: not logined")
		return http.StatusUnauthorized, nil, nil
	}
	userId := usr.Value
	fmt.Println(":update_user: login userid", userId)

	//管理者かどうか
	userService := service.NewUserService(d.DB)
	usrinfo, isAdmin := userService.IsAdmin(userId)
	if !isAdmin {
		if usrinfo.Id == "" {
			fmt.Println(":update_user: not found user")
			return http.StatusBadRequest, nil, nil
		}
		fmt.Println(":update_use: not admin")
		return http.StatusForbidden, nil, nil
	}
	fmt.Println(":update_user: admin!")

	uId := r.FormValue("user_id")
	userName := r.FormValue("user_name")
	userComment := r.FormValue("user_comment")
	pass := r.FormValue("password")
	tmp := r.FormValue("is_admin")

	isa,_:=strconv.ParseBool(tmp)

	//user isExists
	if !userService.IsExists(uId) {
		fmt.Println(":update_user: not exists user")
		return http.StatusBadRequest, nil, nil
	}
	fmt.Println(":update_user: exists user")

	//update
	//with full
	if pass!="" {
		_, err := repository.UpdateUserFull(d.DB, uId, userName, userComment, pass, isa)
		if err != nil {
			return http.StatusInternalServerError, nil, nil
		}
	}else{
		_, err := repository.UpdateUserWithoutPass(d.DB, uId, userName, userComment, isa)
		if err != nil {
			return http.StatusInternalServerError, nil, nil
		}

	}
	return http.StatusOK,nil,nil
}

func (d *DBHandler)UpdateUserSelf(w http.ResponseWriter,r *http.Request)(int,interface{},error){
	fmt.Println(":update_user_self:")
	usr, _ := r.Cookie("user_id")
	if usr == nil {
		fmt.Println(":update_user_self: not logined")
		return http.StatusUnauthorized, nil, nil
	}
	userId := usr.Value
	fmt.Println(":update_user_self: login userid", userId)

	userName := r.FormValue("user_name")
	userComment := r.FormValue("user_comment")

	fmt.Println(userName,userComment)

	//update
	_, err := repository.UpdateUserSelf(d.DB, userId, userName, userComment)
	if err != nil {
		return http.StatusInternalServerError, nil, nil
	}
	fmt.Println(":update_user_self: updated!")
	return http.StatusOK,nil,nil
}
//func (d *DBHandler) GetEventDetail(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
//	vars := mux.Vars(r)
//	eventId, ok := vars["id"]
//	if !ok {
//		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter id"}
//	}
//
//	eventService := service.NewEventService(d.DB)
//	resp, err := eventService.GetEventDetailById(eventId, r)
//	if err != nil {
//		return http.StatusInternalServerError, nil, err
//	}
//
//	return http.StatusOK, resp, nil
//}
