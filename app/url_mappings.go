package app

import (
	"github.com/isamukhtarov/bookstore_users-api/controllers/ping_controller"
	"github.com/isamukhtarov/bookstore_users-api/controllers/users"
)

func mapUrls()  {
	router.GET("/ping", ping_controller.Ping)
	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
}
