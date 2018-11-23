package models

// User is a struct to hold a user
type User struct {
	ID        int    `json:"userID"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
