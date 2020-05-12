package services

import (
	"github.com/isamukhtarov/bookstore_users-api/domain/users"
	"github.com/isamukhtarov/bookstore_users-api/utils/errors"
)

// Function for get user by id
func GetUser(userId int64) (*users.User, *errors.RestErr){
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

// Function for create an user
func CreateUser(user users.User) (*users.User, *errors.RestErr){
	if err := user.Validate(); err !=nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
