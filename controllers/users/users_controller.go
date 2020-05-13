package users

import (
	"github.com/gin-gonic/gin"
	"github.com/isamukhtarov/bookstore_users-api/domain/users"
	"github.com/isamukhtarov/bookstore_users-api/services"
	"github.com/isamukhtarov/bookstore_users-api/utils/errors"
	"net/http"
	"strconv"
)

var(
	counter int
)

// Controller for create new User
func CreateUser(c *gin.Context)  {
	var user = users.User{}
	//fmt.Println(user)
	//bytes, err := ioutil.ReadAll(c.Request.Body)
	//if err != nil {
	//	//TODO: Handle error
	//	return
	//}
	//if err := json.Unmarshal(bytes,&user); err != nil {
	//	fmt.Println(err.Error())
	//	//TODO: Handle json error
	//	return
	//}

	// This part of code is may be replace instead commented codes
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid Json Body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

// Controller for Get a user by given id
func GetUser(c *gin.Context)  {
	userId, userErr := strconv.ParseInt(c.Param("user_id"),10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("User id should be a number")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
	}
	c.JSON(http.StatusOK, user)
}
