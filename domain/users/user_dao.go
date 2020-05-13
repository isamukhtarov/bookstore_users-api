package users

import (
	"fmt"
	"github.com/isamukhtarov/bookstore_users-api/datasources/mysql/users_db"
	"github.com/isamukhtarov/bookstore_users-api/utils/date_utils"
	"github.com/isamukhtarov/bookstore_users-api/utils/errors"
	"strings"
)

const (
	indexUniqueEmail = "users.email_UNIQUE"
	errorNoRows      = "no rows in result set"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
)

//var UserDB = make(map[int64]*User)

// Get user (using map)
func (user *User) Get() *errors.RestErr{
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil{
		if strings.Contains(err.Error(), errorNoRows){
			return errors.NewNotFoundError(fmt.Sprintf("User with id %d not found", user.Id))
		}
		return errors.NewInternalServerError(fmt.Sprintf("Error when trying to get user with id %d: %s", user.Id, err.Error()))
	}
	return nil
}

// Save User into Database
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.DateCreated = date_utils.GetNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail){
			return errors.NewBadRequestError(fmt.Sprintf("Email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("Error while to save user: %s", err.Error()))
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error when trying to save user: %s", err.Error()))
	}
	user.Id = userId
	return nil
}