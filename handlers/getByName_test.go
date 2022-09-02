package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"planets.api/handlers"
	"planets.api/models"
	"planets.api/services"
)

func TestGetPlanetByName(t *testing.T) {
	client := &services.MockPlanetClient{}

	tests := map[string]struct {
		name         string
		expectedCode int
	}{
		"should return 200": {
			name:         t.Name(),
			expectedCode: 200,
		},
		"should return 404": {
			name:         "",
			expectedCode: 404,
		},
	}

	for namme, test := range tests {
		t.Run(namme, func(t *testing.T) {
			if test.expectedCode == 200 {
				client.On("GetByName", test.name).Return(models.Planets{}, nil)
			}
			req, _ := http.NewRequest("GET", "/planets/get/"+test.name, nil)
			rec := httptest.NewRecorder()

			r := gin.Default()
			r.GET("/planets/get/:name", handlers.GetPlanetByName(client))
			r.ServeHTTP(rec, req)

			if test.expectedCode == 200 {
				client.AssertExpectations(t)
			} else {
				client.AssertNotCalled(t, "GetByName")
			}
		})
	}
}
