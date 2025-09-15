package services

import (
	"spe/models"
)

type AdminService struct{}

func (AdminService) Create(newAdmin models.Admin) (models.Admin, error) {
	// Todo
	return models.Admin{}, nil
}

func (AdminService) FindAll() ([]models.Admin, error) {
	// Todo
	return []models.Admin{}, nil
}

func (AdminService) Update() (*models.Admin, error) {
	// Todo
	return &models.Admin{}, nil
}

func (AdminService) Delete() error {
	// Todo
	return nil
}
