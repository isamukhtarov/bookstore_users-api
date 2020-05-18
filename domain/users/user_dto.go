package users

import (
	"github.com/isamukhtarov/bookstore_users-api/utils/errors"
	"regexp"
	"strings"
)

const StatusActive = "active"

// Main struct of User model
type User struct {
	Id 			int64  `json:"id"`
	FirstName 	string `json:"first_name"`
	LastName 	string `json:"last_name"`
	Email 		string `json:"email"`
	DateCreated string `json:"date_created"`
	Status 		string `json:"status"`
	Password    string `json:"password"`
}
//Create type of user slice
type Users []User

// Regular expression for email validation
var mailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

//Function for email format validation
func ValidateEmail(email string) bool {
	if !mailRegexp.MatchString(email){
		return false
	}
	return true
}

// Validate function for user instance before save
func (user *User) Validate() *errors.RestErr {
	// Delete spaces at first_name, last_name and email before saving
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	if user.Email == ""{
		return errors.NewBadRequestError("Email addres is required")
	}
	if !ValidateEmail(user.Email){
		return errors.NewBadRequestError("Wrong email format")
	}
	if strings.TrimSpace(user.Password)== "" || len(strings.TrimSpace(user.Password)) < 8{
		return errors.NewBadRequestError("Password is required and password length must be higher than 8 characters")
	}


	return nil
}