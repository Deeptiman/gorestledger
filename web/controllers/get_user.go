package controllers

import (
	"encoding/json"
	"net/http"
)

func (app *Application) GetUserDataByEmailHandler() func(http.ResponseWriter, *http.Request) {

	return app.isAuthorized(func(w http.ResponseWriter, r *http.Request) {

		var userdata ModelUserData

		_ = json.NewDecoder(r.Body).Decode(&userdata)

		email := userdata.Email

		fabricUser, err := app.Fabric.SessionUser()

		if err != nil {
			respondJSON(w, map[string]string{"error": "No Session Available"})
		} else {

			if fabricUser == nil {
				respondJSON(w, map[string]string{"error": "Session is Null"})
			} else {

				UserData, err := fabricUser.GetUserFromLedger(email)

				if err != nil {
					respondJSON(w, map[string]string{"error": "No User Data Found"})
				} else {
					respondJSON(w, UserData)
				}
			}
		}
	})
}

func (app *Application) GetAllUsersDataHandler() func(http.ResponseWriter, *http.Request) {

	return app.isAuthorized(func(w http.ResponseWriter, r *http.Request) {

		fabricUser, err := app.Fabric.SessionUser()

		if err != nil {
			respondJSON(w, map[string]string{"error": "No Session Available"})
		} else {

			if fabricUser == nil {
				respondJSON(w, map[string]string{"error": "Session is Null"})
			} else {

				allUsersData, err := fabricUser.GetAllUsersFromLedger()

				if err != nil {
					respondJSON(w, map[string]string{"error": "No User Data Found"})
				} else {
					respondJSON(w, allUsersData)
				}
			}
		}
	})
}
