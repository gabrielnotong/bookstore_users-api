package app

import (
	"github.com/gabrielnotong/bookstore_users-api/controllers/ping"
	"github.com/gabrielnotong/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", handler_ping.Pong)

	router.GET("/users/:id", handler_users.Search)
	router.POST("/users", handler_users.Create)
	router.PUT("/users/:id", handler_users.Update)
}
