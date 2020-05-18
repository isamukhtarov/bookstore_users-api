package errors

import (
	"net/http"
)

// Struct for Error which we will user for handle application errors
type RestErr struct {
	Message string `json:"message"`
	Status int	   `json:"code"`
	Error string   `json:"error"`
}

//func NewError(message string) error {
//	return errors.New(message)
//}

// Custom Bad request error
func NewBadRequestError(message string)*RestErr{
	return &RestErr{
		Message: message,
		Status: http.StatusBadRequest,
		Error: "Bad Request",
	}
}

// Custom Not found request error
func NewNotFoundError(message string)*RestErr{
	return &RestErr{
		Message: message,
		Status: http.StatusNotFound,
		Error: "Not Found",
	}
}

// Custom Internal server error
func NewInternalServerError(message string)*RestErr  {
	return &RestErr{
		Message: message,
		Status: http.StatusInternalServerError,
		Error: "Internal Error",
	}
}
