package app

import (
	"github.com/emadghaffari/grpc_rest_items_service/app/http/middleware"
	"github.com/emadghaffari/grpc_rest_items_service/controllers/item"
	"github.com/emadghaffari/grpc_rest_items_service/controllers/ping"
)



func mapURL()  {
	// healthcheck for rest service
	router.GET("/ping", ping.Ping)

	authorized := router.Group("/items")
	authorized.Use(middleware.AccessTokenMiddleware.CheckMiddleware)
	authorized.POST("/",item.Create)
	authorized.DELETE("/:id",item.Delete)

	router.GET("/items/:id",item.Get)
	router.POST("/items/search",item.Search)
}