package web

import (
	"fmt"
	"log"
	"net/http"

	"gorestledger/web/controllers"

	"github.com/gorilla/mux"
)

func Serve(app *controllers.Application) {

	r := mux.NewRouter()

	r.HandleFunc("/api/get_users", app.GetAllUsersDataHandler()).Methods("GET")
	r.HandleFunc("/api/get_user", app.GetUserDataByEmailHandler()).Methods("GET")

	r.HandleFunc("/api/user_login", app.LoginHandler).Methods("POST")
	r.HandleFunc("/api/user_register", app.RegisterHandler).Methods("POST")
	r.HandleFunc("/api/change_password", app.ChangePwdHandler()).Methods("POST")

	r.HandleFunc("/api/update_user", app.UpdateUserHandler()).Methods("PUT")

	r.HandleFunc("/api/delete_user", app.DeleteUserHandler()).Methods("DELETE")

	fmt.Println("Listening (http://localhost:4000) ...")
	log.Fatal(http.ListenAndServe(":4000", r))
}
