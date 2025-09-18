package repository

import (
	"spe/models"
)

var (
	userIDs int64 = 1
	users         = map[int64]*models.User{}
)

func init() {
	leonardo := &models.User{
		ID:       userIDs,
		Name:     "Leonardo",
		Username: "leonardo",
		Email:    "leonardo@dimap.ufrn.br",
	}
	_ = leonardo.SetPassword("A1B2C3")
	users[leonardo.ID] = leonardo
	userIDs++

	carlos := &models.User{
		ID:       userIDs,
		Name:     "Carlos",
		Username: "carlos",
		Email:    "carlos@dimap.ufrn.br",
	}
	_ = carlos.SetPassword("12345")
	users[carlos.ID] = carlos
	userIDs++

	cicero := &models.User{
		ID:       userIDs,
		Name:     "Cicero",
		Username: "cicero",
		Email:    "cicero@dimap.ufrn.br",
	}
	_ = cicero.SetPassword("mafiacomunista")
	users[cicero.ID] = cicero
	userIDs++

	leandro := &models.User{
		ID:       userIDs,
		Name:     "Leandro",
		Username: "leandro",
		Email:    "leandro@dimap.ufrn.br",
	}
	_ = leandro.SetPassword("qwerty@12345")
	users[leandro.ID] = leandro
	userIDs++

	cadu := &models.User{
		ID:       userIDs,
		Name:     "Carlos Eduardo",
		Username: "cadu",
		Email:    "cadu@dimap.ufrn.br",
	}
	_ = cadu.SetPassword("A1B2C3")
	users[cadu.ID] = cadu
	userIDs++
}

type UserRepository struct{}

func (UserRepository) Save(u *models.User) *models.User {
	u.ID = userIDs
	userIDs++
	users[u.ID] = u
	return u
}

func (UserRepository) FindAll() []*models.User {
	list := make([]*models.User, 0, len(users))
	for _, u := range users {
		list = append(list, u)
	}
	return list
}

func (UserRepository) FindByID(uid int64) (*models.User, bool) {
	if user, ok := users[uid]; ok {
		return user, true
	}
	return nil, false
}

func (UserRepository) FindByUsername(u string) (*models.User, bool) {
	for _, user := range users {
		if user.Username == u {
			return user, true
		}
	}
	return nil, false
}

func (UserRepository) Update(u *models.User) bool {
	if _, ok := users[u.ID]; !ok {
		return false
	}
	users[u.ID] = u
	return true
}

func (UserRepository) Delete(uid int64) bool {
	if _, ok := users[uid]; !ok {
		return false
	}
	delete(users, uid)
	return true
}
