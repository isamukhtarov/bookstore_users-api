package users

import (
	"fmt"
	"github.com/isamukhtarov/bookstore_users-api/datasources/mysql/users_db"
	"github.com/isamukhtarov/bookstore_users-api/logger"
	"github.com/isamukhtarov/bookstore_users-api/utils/errors"
)

// Constants of MySql database queries
const (
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created, password, status) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	queryUpdateUser  = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?"
	queryDeleteUser  = "DELETE FROM users WHERE id=?"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
)

//var UserDB = make(map[int64]*User)

// Get user (using map)
func (user *User) Get() *errors.RestErr{
	// Prepare query for get user
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	// Get result by user Id
	result := stmt.QueryRow(user.Id)
	// Handle errors if exists else scan got user credentials
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil{
		logger.Error("Error when trying to get user by id", getErr)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

// Save User into Database
func (user *User)  Save() *errors.RestErr {
	//Prepare mysql query
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("Error when trying to prepare save user", err)
		return errors.NewInternalServerError("database error")
	}
	// Close db
	defer stmt.Close()
	// Insert data into mysql database
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Password, user.Status)
	// Handle errors while insert
	if saveErr != nil {
		logger.Error("Error when trying to save user", err)
		return errors.NewInternalServerError("database error")
	}
	// Get saved user id
	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("Error when trying to get saved user id", err)
		return errors.NewInternalServerError("database error")
	}
	user.Id = userId
	return nil
}

// Update user data
func (user *User) Update() *errors.RestErr{
	// Prepare user update query
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("Error when trying to prepare update user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	// Execute user update query form database
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		logger.Error("Error when trying to update user", err)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

// Delete user
func (user *User) Delete() *errors.RestErr {
	// Prepare user delete query
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("Error when trying to prepare delete user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	// Execute user delete query form database
	if _, err = stmt.Exec(user.Id); err != nil{
		logger.Error("Error when trying to delete user", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}


// Find User By status
func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	// Prepare search query
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("Error when trying to prepare query to get users by status param statement", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	// Create user query by user status
	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	// Create empty User slice
	results := make([]User, 0)

	// Loop found rows finding in database (if exists)
	for rows.Next(){
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("Error when trying to scan row into user struct", err)
			return nil, errors.NewInternalServerError("database error")
		}
		// Append to result slice finding user
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("User by given status %s not found", status))
	}
	return results, nil
}