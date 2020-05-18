package services

import (
	"github.com/isamukhtarov/bookstore_users-api/domain/users"
	"github.com/isamukhtarov/bookstore_users-api/utils/crypto_utils"
	"github.com/isamukhtarov/bookstore_users-api/utils/date_utils"
	"github.com/isamukhtarov/bookstore_users-api/utils/errors"
)

// Create variable type of below interface
var UserService userServiceInterface = &userService{}

// Struct which will implement all below interface methods
type userService struct {}

// Create interface for CRUD methods
type userServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestErr)
	CreateUser(users.User) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	Search(string) (users.Users, *errors.RestErr)
}

// Function for get user by id
func (u *userService) GetUser(userId int64) (*users.User, *errors.RestErr){
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

// Function for create an user
func (u *userService) CreateUser(user users.User) (*users.User, *errors.RestErr){
	if err := user.Validate(); err !=nil {
		return nil, err
	}

	user.DateCreated = date_utils.GetNowDbFormat()
	// Hashing of user password before save
	user.Password = crypto_utils.GetMd5(user.Password)
	user.Status = users.StatusActive
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}


// Function for update user
func (u *userService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	// Get user by id
	current, err := u.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	// check got data validation
	if err := user.Validate(); err != nil {
		return nil, err
	}

	// Update user partial or all data
	if !isPartial{
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}else{
		if user.FirstName != ""{
			current.FirstName = user.FirstName
		}
		if user.LastName != ""{
			current.LastName = user.LastName
		}
		if user.Email != ""{
			current.Email = user.Email
		}
	}

	if err := current.Update();err != nil {
		return nil, err
	}
	return current, nil
}

// Function for delete user
func (u *userService) DeleteUser(userId int64) *errors.RestErr {
	currentUser := &users.User{Id: userId}
	return currentUser.Delete()
}

// Find by status function
func (u *userService) Search(status string) (users.Users, *errors.RestErr){
	dao := &users.User{}
	return dao.FindByStatus(status)
}