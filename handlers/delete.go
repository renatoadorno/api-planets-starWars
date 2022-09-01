package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"planets.api/services"
)

func Delete(sr services.PlanetInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		planetName := ctx.Param("name")

		err := sr.Delete(planetName)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "planet deleted successfully"})
	}
}
