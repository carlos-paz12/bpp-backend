package repository

import (
	"errors"
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

// UserRepository é um repositório de usuários em memória.
type UserRepository struct{}

// Create adiciona um novo usuário no repositório.
func (UserRepository) Create(user *models.User) error {
	if _, exists := users[user.ID]; exists {
		return errors.New("Usuário já existe.")
	}
	users[user.ID] = user
	return nil
}

// FindAll retorna todos os usuários do repositório.
func (UserRepository) FindAll() []*models.User {
	list := make([]*models.User, 0, len(users))
	for _, u := range users {
		list = append(list, u)
	}
	return list
}

// FindByID retorna um usuário pelo ID.
func (UserRepository) FindByID(id int64) (*models.User, error) {
	if user, ok := users[id]; ok {
		return user, nil
	}
	return nil, errors.New("Usuário não encontrado.")
}

// FindByUsername retorna um usuário pelo username.
func (UserRepository) FindByUsername(username string) (*models.User, error) {
	for _, user := range users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, errors.New("Usuário não encontrado.")
}

// Update altera os dados de um usuário existente.
func (UserRepository) Update(user *models.User) error {
	if _, ok := users[user.ID]; !ok {
		return errors.New("Usuário não encontrado.")
	}
	users[user.ID] = user
	return nil
}

// Delete remove um usuário pelo ID.
func (UserRepository) Delete(id int64) error {
	if _, ok := users[id]; !ok {
		return errors.New("Usuário não encontrado.")
	}
	delete(users, id)
	return nil
}
