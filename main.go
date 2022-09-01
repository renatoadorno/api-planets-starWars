package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"planets.api/config"
	"planets.api/database"
	"planets.api/handlers"
	"planets.api/services"
)

func main() {
	conf := config.GetConfig()
	ctx := context.TODO()

	db := database.ConnectDB(ctx, conf.Mongo)
	collection := db.Collection(conf.Mongo.Collection)
	client := services.NewClient(collection, ctx)

	r := gin.Default()
	planets := r.Group("/planets")
	planets.Use(Authorization(conf.Token))

  {
		planets.POST("/", handlers.InsertPlanet(client))
		planets.GET("/:id", handlers.GetPlanet(client))
		planets.GET("/:name", handlers.GetPlanetByName(client))
	}

	r.Run(":6000")
}

func Authorization(token string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		if token != auth {
			ctx.AbortWithStatusJSON(401, gin.H{"message": "Invalid authorization token"})
			return
		}
		ctx.Next()
	}
}
