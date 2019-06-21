package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"unicode"
)

func (app *Application) RegisterHandler(w http.ResponseWriter, r *http.Request) {

	var userdata ModelUserData

	_ = json.NewDecoder(r.Body).Decode(&userdata)

	nameValue := userdata.Name
	emailValue := userdata.Email
	companyValue := userdata.Company
	occupationValue := userdata.Occupation
	salaryValue := userdata.Salary
	userType := userdata.UserType

	passwordValue := hash(userdata.Password)

	verifyErr := verifyPassword(userdata.Password)

	if verifyErr != nil && len(verifyErr.Error()) > 0 {

		respondJSON(w, map[string]string{"error": verifyErr.Error(), "message": "Password must contain at least one number and one uppercase and lowercase letter, and at least 8 or more characters"})
	}

	fmt.Println("Sign up --->  nameValue = " + nameValue + " , emailValue = " + emailValue + " , companyValue = " + companyValue + " , occupationValue = " + occupationValue + " , salaryValue = " + salaryValue + " , userType = " + userType)

	fabricUser, err := app.Fabric.RegisterUser(nameValue, emailValue, companyValue, occupationValue, salaryValue, passwordValue, userType)

	if err != nil {
		respondJSON(w, map[string]string{"error": "Unable to Register : " + err.Error()})
	} else {

		token := app.processAuthentication(w, emailValue)

		if len(token) > 0 {

			UserData, err := fabricUser.GetUserFromLedger(emailValue)

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

func verifyPassword(password string) error {
	var uppercasePresent bool
	var lowercasePresent bool
	var numberPresent bool
	var specialCharPresent bool
	const minPassLength = 8
	const maxPassLength = 64
	var passLen int
	var errorString string

	for _, ch := range password {
		switch {
		case unicode.IsNumber(ch):
			numberPresent = true
			passLen++
		case unicode.IsUpper(ch):
			uppercasePresent = true
			passLen++
		case unicode.IsLower(ch):
			lowercasePresent = true
			passLen++
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			specialCharPresent = true
			passLen++
		case ch == ' ':
			passLen++
		}
	}
	appendError := func(err string) {
		if len(strings.TrimSpace(errorString)) != 0 {
			errorString += ", " + err
		} else {
			errorString = err
		}
	}
	if !lowercasePresent {
		appendError("lowercase letter missing")
	}
	if !uppercasePresent {
		appendError("uppercase letter missing")
	}
	if !numberPresent {
		appendError("atleast one numeric character required")
	}
	if !specialCharPresent {
		appendError("special character missing")
	}
	if !(minPassLength <= passLen && passLen <= maxPassLength) {
		appendError(fmt.Sprintf("password length must be between %d to %d characters long", minPassLength, maxPassLength))
	}

	if len(errorString) != 0 {
		return fmt.Errorf(errorString)
	}
	return nil
}
