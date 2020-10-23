package app

import (
	"github.com/gin-gonic/gin"

	"github.com/emadghaffari/res_errors/logger"
)

var (
	router = gin.Default()
)

// StartApplication func
func StartApplication() {
	mapURL()
	logger.Info("about to start application")
	router.Run(":8080")
}
