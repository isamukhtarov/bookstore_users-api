package users

import (
	"fmt"
	"github.com/isamukhtarov/bookstore_users-api/utils/errors"
)


var UserDB = make(map[int64]*User)

// Get user (using map)
func (user *User) Get() *errors.RestErr{
	result := UserDB[user.Id]
	if result == nil{
		return errors.NewNotFoundError(fmt.Sprintf("User %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

// Save User (using map)
func (user *User) Save() *errors.RestErr {
	current := UserDB[user.Id]
	if current != nil {
		if current.Email == user.Email{
			return errors.NewBadRequestError(fmt.Sprintf("User with %s is already exists", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("User with id %d is already exists", user.Id))
	}
	UserDB[user.Id] = user
	return nil
}