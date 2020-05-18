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

func testServiceInterface()  {

}

func getUserId(userIdParam string) (int64, *errors.RestErr){
	userId, userErr := strconv.ParseInt(userIdParam,10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("User id should be a number")
	}

	return userId, nil
}

// Controller for create new User
func Create(c *gin.Context)  {
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

	result, saveErr := services.UserService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusOK, result.Marshal(c.GetHeader("X-Public") == "true"))
}

// Controller for Get a user by given id
func Get(c *gin.Context)  {
	// Check user id param parse as integer
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	user, getErr := services.UserService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
	}
	c.JSON(http.StatusOK, user.Marshal(c.GetHeader("X-Public") == "true"))
}

// Controller for update User data
func Update(c *gin.Context){
	// Check user id param parse as integer
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User
	// Handle request data, give it to User model in json format and handle errors if exists
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid Json Body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	// Variable for patch request method(if true)
	isPartial := c.Request.Method == http.MethodPatch


	result, err := services.UserService.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result.Marshal(c.GetHeader("X-Public") == "true"))
}

// Delete User Controller
func Delete(c *gin.Context)  {
	// Check user id param parse as integer
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.UserService.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "User successfully deleted"})
}

// Find by status controller function
func Search(c *gin.Context){
	status := c.Query("status")
	foundUsers, err := services.UserService.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, foundUsers.Marshal(c.GetHeader("X-Public") == "true"))
}