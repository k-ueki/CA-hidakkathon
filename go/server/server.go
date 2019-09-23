package server

import (
	"hidakkathon/controller"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	// db     *sqlx.DB
	router *mux.Router
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init() {
	s.router = s.Route()
}

func (s *Server) Route() *mux.Router {
	var DSN string = "hidakkathon:hidakkathon@tcp(localhost:3306)/sugori_rendez_vous"
	//var DSN string = "root:@tcp(go_db_1:3306)/sugori_rendez_vous"
	r := mux.NewRouter()
	db, err := sqlx.Open("mysql", DSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	Controller := &controller.DBHandler{DB: db}

	r.Methods(http.MethodGet).Path("/hc").Handler(AppHandler{Controller.Test})

	r.Methods(http.MethodGet).Path("/top").Handler(AppHandler{Controller.Index})
	//途中
	//r.Methods(http.MethodGet).Path("/user").Handler(AppHandler{Controller.GetSelf})
	r.Methods(http.MethodGet).Path("/event/image/{event_id}").HandlerFunc(Controller.GetEventImage)
	r.Methods(http.MethodGet).Path("/admin/tag").Handler(AppHandler{Controller.GetTag})
	r.Methods(http.MethodGet).Path("/admin/target_user_type").Handler(AppHandler{Controller.GetTargetUserType})
	r.Methods(http.MethodGet).Path("/admin/users").Handler(AppHandler{Controller.GetAllUsers})
	//↓途中
	r.Methods(http.MethodGet).Path("/user/{user_id}").Handler(AppHandler{Controller.GetUser})

	r.Methods(http.MethodPost).Path("/auth/login").Handler(AppHandler{Controller.Login})
	r.Methods(http.MethodPost).Path("/auth/logout").Handler(AppHandler{Controller.Logout})
	r.Methods(http.MethodPost).Path("/admin/tag").Handler(AppHandler{Controller.CreateTag})
	r.Methods(http.MethodPost).Path("/admin/target_user_type").Handler(AppHandler{Controller.CreateTargetUserType})
	r.Methods(http.MethodPost).Path("/admin/user").Handler(AppHandler{Controller.CreateUser})

	r.Methods(http.MethodPut).Path("/admin/target_user_type").Handler(AppHandler{Controller.UpdateTargetUserType})
	r.Methods(http.MethodPut).Path("/admin/user").Handler(AppHandler{Controller.UpdateUser})
	//途中
	r.Methods(http.MethodPut).Path("/user").Handler(AppHandler{Controller.UpdateUserSelf})

	r.Methods(http.MethodDelete).Path("/event/{event_id}").Handler(AppHandler{Controller.DeleteEvent})
	r.Methods(http.MethodDelete).Path("/admin/user/{user_id}").Handler(AppHandler{Controller.DeleteUser})
	r.Methods(http.MethodDelete).Path("/admin/tag/{tag_id}").Handler(AppHandler{Controller.DeleteTag})
	r.Methods(http.MethodDelete).Path("/admin/target_user_type/{id}").Handler(AppHandler{Controller.DeleteTargetUserType})

//	r.Methods(http.MethodGet).Path("/event/{id}").Handler(AppHandler{Controller.GetEventDetail})

	//r.Methods(http.MethodGet).Path("/event/image/{event_id}").Handler(Controller.GetEventImage)


	http.ListenAndServe(":8080", r)

	return r
}
