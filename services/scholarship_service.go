package services

import (
	"spe/models"
)

type ScholarshipService struct{}

func (ScholarshipService) Create(newScholarship models.Scholarship) (models.Scholarship, error) {
	// Todo
	return models.Scholarship{}, nil
}

func (ScholarshipService) FindAll() ([]models.Scholarship, error) {
	// Todo
	return []models.Scholarship{}, nil
}

func (ScholarshipService) Update() (*models.Scholarship, error) {
	// Todo
	return &models.Scholarship{}, nil
}

func (ScholarshipService) Delete() error {
	// Todo
	return nil
}
