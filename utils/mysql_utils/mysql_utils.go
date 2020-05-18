package mysql_utils

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/isamukhtarov/bookstore_users-api/utils/errors"
	"strings"
)

const errorNoRows  = "no rows in result set"


// func for parse sql error
func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	fmt.Println(sqlErr, ok)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows){
			return errors.NewNotFoundError("No record matching by given id")
		}
		return errors.NewInternalServerError("Error parsing mysql database response")
	}

	// Check mysql error number
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("Given email already exists")
	}

	return errors.NewInternalServerError("Error processing request")
}