package repository

import "spe/models"

var (
	adminsIDs int64 = 1
	admins          = map[int64]*models.Admin{}
)

func init() {
	leonardo := &models.Admin{
		ID:     adminsIDs,
		UserID: 1,
	}
	admins[leonardo.ID] = leonardo
	adminsIDs++

	cadu := &models.Admin{
		ID:     adminsIDs,
		UserID: 5,
	}
	admins[cadu.ID] = cadu
	adminsIDs++
}

type AdminRepository struct{}

func (AdminRepository) Create() {
	// Todo
}

func (AdminRepository) FindAll() {
	// Todo
}

func (AdminRepository) FindByUserID(id int64) (models.Admin, error) {
	// Todo
	return models.Admin{}, nil
}

func (AdminRepository) Update() {
	// Todo
}

func (AdminRepository) Delete() {
	// Todo
}
