package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"planets.api/services"
)

func GetPlanet(sr services.PlanetInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		res, err := sr.Get(id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, res)
	}
}
