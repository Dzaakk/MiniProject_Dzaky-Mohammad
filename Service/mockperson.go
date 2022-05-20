package service

import (
	"MiniProject/models"
)

type MockPersonService struct {
	data []models.Users
}

func (ps *MockPersonService) Add(user models.Users) (models.Users, error) {
	ps.data = append(ps.data, user)
	return user, nil
}

func (ps *MockPersonService) Get() ([]models.Users, error) {
	return ps.data, nil
}

func NewMockPersonService() *MockPersonService {
	return &MockPersonService{
		data: []models.Users{},
	}
}
