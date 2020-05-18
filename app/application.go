package app

import (
	"github.com/gin-gonic/gin"
	"github.com/isamukhtarov/bookstore_users-api/logger"
)

// Call gin framework router
var(
	router = gin.Default()
)

// function witch calls all application routers and run the application in current host
func StartApplication()  {
	mapUrls()
	logger.Info("about start the application...")
	router.Run(":8080")
}
