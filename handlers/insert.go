package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"planets.api/models"
	"planets.api/services"
)

func InsertPlanet(sr services.PlanetInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		planet := models.Planets{}

		err := ctx.BindJSON(&planet)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		res, err := sr.Insert(planet)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, res)
	}
}
