package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"planets.api/services"
)

func GetPlanetByName(sr services.PlanetInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		planetName := ctx.Param("name")

		res, err := sr.GetByName(planetName)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, res)
	}
}
