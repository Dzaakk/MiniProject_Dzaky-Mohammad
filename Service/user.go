package service

import "MiniProject/models"

type PersonService interface {
	Add(person models.Users) (models.Users, error)
	Get() ([]models.Users, error)
}
