package service

import (
	"MiniProject/models"

	"gorm.io/gorm"
)

type DBPersonService struct {
	db *gorm.DB
}

func (ps DBPersonService) Add(person models.Users) (models.Users, error) {
	tx := ps.db.Save(&person)
	err := tx.Error
	return person, err
}

func (ps DBPersonService) Get() ([]models.Users, error) {
	users := []models.Users{}
	tx := ps.db.Find(&users)
	err := tx.Error
	return users, err
}

func NewDBPersonService(db *gorm.DB) DBPersonService {
	return DBPersonService{
		db: db,
	}
}
