package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *Application) UpdateUserHandler() func(http.ResponseWriter, *http.Request) {

	return app.isAuthorized(func(w http.ResponseWriter, r *http.Request) {

		var userdata ModelUserData

		_ = json.NewDecoder(r.Body).Decode(&userdata)

		userId := userdata.ID
		userType := userdata.UserType
		name := userdata.Name
		company := userdata.Company
		occupation := userdata.Occupation
		salary := userdata.Salary

		fmt.Println("UpdateUser == " + userId + " -- " + userType + " -- " + name + " -- " + company + " -- " + occupation + " -- " + salary)

		fabricUser, err := app.Fabric.SessionUser()

		if err != nil {
			respondJSON(w, map[string]string{"error": "Error Session User  " + err.Error()})
		} else {

			email := fabricUser.Username

			err = fabricUser.UpdateUserData(userId, name, email, company, occupation, salary, userType)

			if err != nil {

				respondJSON(w, map[string]string{"error": "Error Update User Data = " + err.Error()})

			} else {

				UserData, err := fabricUser.GetUserFromLedger(email)

				if err != nil {

					respondJSON(w, map[string]string{"error": "No User Data Found"})

				} else {

					respondJSON(w, map[string]string{
						"id":         UserData.ID,
						"name":       UserData.Name,
						"email":      UserData.Email,
						"company":    UserData.Company,
						"occupation": UserData.Occupation,
						"salary":     UserData.Salary,
						"userType":   UserData.UserType})
				}
			}
		}
	})
}
