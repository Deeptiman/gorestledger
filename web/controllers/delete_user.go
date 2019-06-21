package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *Application) DeleteUserHandler() func(http.ResponseWriter, *http.Request) {

	return app.isAuthorized(func(w http.ResponseWriter, r *http.Request) {

		var userdata ModelUserData

		_ = json.NewDecoder(r.Body).Decode(&userdata)

		email := userdata.Email

		fmt.Println("DeleteUserHandler : Email = " + email)

		fabricUser, err := app.Fabric.SessionUser()

		if err != nil {
			respondJSON(w, map[string]string{"error": "Error Session User  " + err.Error()})
		} else {

			if len(fabricUser.Username) > 0 {

				err := fabricUser.RemoveUser(email)

				if err != nil {
					fmt.Println("DeleteUserHandler : RemoveUser = Error : " + err.Error())
					respondJSON(w, map[string]string{"error": "Error Session User  " + err.Error()})
				} else {
					fmt.Println("Success RemoveUser ")
					respondJSON(w, map[string]string{"success": "Succesfully delete the user with email - " + email})
				}

			} else {
				respondJSON(w, map[string]string{"error": "No Session is available "})
			}
		}
	})
}
