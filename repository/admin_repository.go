package repository

import (
	"spe/models"
)

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

func (AdminRepository) Save() {
	// Todo
}

func (AdminRepository) FindAll() {
	// Todo
}

func (AdminRepository) FindByUserID(uid int64) (*models.Admin, bool) {
	for _, a := range admins {
		if a.UserID == uid {
			return a, true
		}
	}
	return nil, false
}

func (AdminRepository) Update() {
	// Todo
}

func (AdminRepository) Delete() {
	// Todo
}
