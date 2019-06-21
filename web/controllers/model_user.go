package controllers

type ModelUserData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	OldPassword string `json:"oldPassword"`
	Company     string `json:"company"`
	Occupation  string `json:"occupation"`
	Salary      string `json:"salary"`
	UserType    string `json:"userType"`
}
