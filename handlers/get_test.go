package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"planets.api/handlers"
	"planets.api/models"
	"planets.api/services"
)

func TestGetPlanet(t *testing.T) {
	client := &services.MockPlanetClient{}
	id := primitive.NewObjectID().Hex()

	tests := map[string]struct {
		id           string
		expectedCode int
	}{
		"should return 200": {
			id:           id,
			expectedCode: 200,
		},
		"should return 404": {
			id:           "",
			expectedCode: 404,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if test.expectedCode == 200 {
				client.On("Get", test.id).Return(models.Planets{}, nil)
			}
			req, _ := http.NewRequest("GET", "/planets/"+test.id, nil)
			rec := httptest.NewRecorder()

			r := gin.Default()
			r.GET("/planets/:id", handlers.GetPlanet(client))
			r.ServeHTTP(rec, req)

			if test.expectedCode == 200 {
				client.AssertExpectations(t)
			} else {
				client.AssertNotCalled(t, "Get")
			}
		})
	}
}
