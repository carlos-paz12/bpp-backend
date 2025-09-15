package repository

import "spe/models"

var scholarships = map[int64]*models.Scholarship{}

func init() {
	carlos := &models.Scholarship{
		ID:            1,
		UserID:        2,
		Register:      "20240000000",
		HoursPerMonth: 80,
	}
	scholarships[carlos.ID] = carlos

	cicero := &models.Scholarship{
		ID:            2,
		UserID:        3,
		Register:      "20241111111",
		HoursPerMonth: 80,
	}
	scholarships[cicero.ID] = cicero

	leandro := &models.Scholarship{
		ID:            3,
		UserID:        4,
		Register:      "20242222222",
		HoursPerMonth: 80,
	}
	scholarships[leandro.ID] = leandro
}

type ScholarshipRepository struct{}

func (ScholarshipRepository) Create() {
	// Todo
}

func (ScholarshipRepository) FindAll() {
	// Todo
}

func (ScholarshipRepository) FindByUserId(id int64) (models.Scholarship, error) {
	// Todo
	return models.Scholarship{}, nil
}

func (ScholarshipRepository) Update() {
	// Todo
}

func (ScholarshipRepository) Delete() {
	// Todo
}
