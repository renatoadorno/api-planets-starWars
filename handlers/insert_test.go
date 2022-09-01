package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"planets.api/handlers"
	"planets.api/models"
	"planets.api/services"
)

func TestInsertPlanet(t *testing.T) {
	client := &services.MockPlanetClient{}
	tests := map[string]struct {
		payload      string
		expectedCode int
	}{
		"should return 201": {
			payload:      `{"name":"Dagobah","climate":"murky", "terrain":"swamp, jungle"}`,
			expectedCode: 201,
		},
		"should return 400": {
			payload:      "invalid json string",
			expectedCode: 400,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			client.On("Insert", mock.Anything).Return(models.Planets{}, nil)
			req, _ := http.NewRequest("POST", "/planets", strings.NewReader(test.payload))
			rec := httptest.NewRecorder()

			r := gin.Default()
			r.POST("/planets", handlers.InsertPlanet(client))
			r.ServeHTTP(rec, req)

			if test.expectedCode == 201 {
				client.AssertExpectations(t)
			} else {
				client.AssertNotCalled(t, "Insert")
			}
		})
	}
}
