package repository

import "spe/models"

var admins = map[int64]*models.Admin{}

func init() {
	leonardo := &models.Admin{
		ID:     1,
		UserID: 1,
	}
	admins[leonardo.ID] = leonardo

	cadu := &models.Admin{
		ID:     1,
		UserID: 5,
	}
	admins[cadu.ID] = cadu
}

type AdminRepository struct{}

func (AdminRepository) Create() {
	// Todo
}

func (AdminRepository) FindAll() {
	// Todo
}

func (AdminRepository) FindByUserId(userID int64) (models.Admin, error) {
	// Todo
	return models.Admin{}, nil
}

func (AdminRepository) Update() {
	// Todo
}

func (AdminRepository) Delete() {
	// Todo
}
