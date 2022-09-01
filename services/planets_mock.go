package services

import (
	"fmt"

	"github.com/stretchr/testify/mock"
  "planets.api/models"
)

type MockPlanetClient struct {
	mock.Mock
}

func (m *MockPlanetClient) Insert(planet models.Planets) (models.Planets, error) {
	args := m.Called(planet)
	return args.Get(0).(models.Planets), args.Error(1)
}

func (m *MockPlanetClient) Get(id string) (models.Planets, error) {
	fmt.Println("call get mock function")
	args := m.Called(id)
	return args.Get(0).(models.Planets), args.Error(1)
}

func (m *MockPlanetClient) GetByName(name string) (models.Planets, error) {
	fmt.Println("call GetByName mock function")
	args := m.Called(name)
	return args.Get(0).(models.Planets), args.Error(1)
}
