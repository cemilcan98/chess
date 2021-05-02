package controller

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterHandlers(instance *echo.Echo, database *mongo.Database) {

	resource := NewResource(database)

	instance.POST("/games", resource.postGame)
	instance.GET("/", resource.helloApi)
	instance.GET("/games/id/:id", resource.getGame)
	instance.GET("/games/user/:user", resource.getGameByUsername)

}
