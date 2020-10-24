package app

import (
	"github.com/gin-gonic/gin"

	"github.com/emadghaffari/grpc_rest_items_service/databases/elasticsearch"
	"github.com/emadghaffari/res_errors/logger"
)

var (
	router = gin.Default()
)

// StartApplication func
func StartApplication() {
	elasticsearch.Init()
	mapURL()
	logger.Info("about to start application")
	router.Run(":8000")
}
