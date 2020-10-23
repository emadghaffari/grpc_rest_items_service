package app

import "github.com/emadghaffari/grpc_rest_items_service/controllers/ping"








func mapURL()  {
	// healthcheck for rest service
	router.GET("/ping", ping.Ping)
}