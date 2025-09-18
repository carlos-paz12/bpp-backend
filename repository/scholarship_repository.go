package repository

import (
	"spe/models"
)

var (
	scholarshipsIDs int64 = 1
	scholarships          = map[int64]*models.Scholarship{}
)

func init() {
	carlos := &models.Scholarship{
		ID:            scholarshipsIDs,
		UserID:        2,
		Register:      "20240000000",
		HoursPerMonth: 80,
	}
	scholarships[carlos.ID] = carlos
	scholarshipsIDs++

	cicero := &models.Scholarship{
		ID:            scholarshipsIDs,
		UserID:        3,
		Register:      "20241111111",
		HoursPerMonth: 80,
	}
	scholarships[cicero.ID] = cicero
	scholarshipsIDs++

	leandro := &models.Scholarship{
		ID:            scholarshipsIDs,
		UserID:        4,
		Register:      "20242222222",
		HoursPerMonth: 80,
	}
	scholarships[leandro.ID] = leandro
	scholarshipsIDs++
}

type ScholarshipRepository struct{}

func (ScholarshipRepository) Save() {
	// Todo
}

func (ScholarshipRepository) FindAll() {
	// Todo
}

func (ScholarshipRepository) FindByUserID(uid int64) (*models.Scholarship, bool) {
	for _, s := range scholarships {
		if s.UserID == uid {
			return s, true
		}
	}
	return nil, false
}

func (ScholarshipRepository) Update() {
	// Todo
}

func (ScholarshipRepository) Delete() {
	// Todo
}
