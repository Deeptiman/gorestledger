package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *Application) LoginHandler(w http.ResponseWriter, r *http.Request) {

	var userdata ModelUserData

	_ = json.NewDecoder(r.Body).Decode(&userdata)

	email := userdata.Email
	password := hash(userdata.Password)

	fmt.Println("Sign In --->  emailValue = " + email)

	fabricUser, err := app.Fabric.LoginUser(email, password)

	if err != nil {
		respondJSON(w, map[string]string{"error": "Unable to Login : " + err.Error()})
	} else {

		fmt.Println("Logged In User : " + fabricUser.Username)

		token := app.processAuthentication(w, email)

		if len(token) > 0 {

			UserData, err := fabricUser.GetUserFromLedger(email)

			if err != nil {
				respondJSON(w, map[string]string{"error": "No User Data Found"})
			} else {
				respondJSON(w, map[string]string{

					"token":      token,
					"id":         UserData.ID,
					"name":       UserData.Name,
					"email":      UserData.Email,
					"company":    UserData.Company,
					"occupation": UserData.Occupation,
					"salary":     UserData.Salary,
					"userType":   UserData.UserType,
				})
			}

		} else {
			respondJSON(w, map[string]string{"error": "Failed to generate token"})
		}
	}
}
