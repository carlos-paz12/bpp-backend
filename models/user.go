package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int64  `json:"id"`       //!< ID do registro de usuário.
	Name     string `json:"name"`     //!< Nome completo do usuário.
	Username string `json:"username"` //!< Nome de usuário.
	Password string `json:"-"`        //!< Senha do usuário (criptografada).
	Email    string `json:"email"`    //!< E-mail do usuário.
}

// SetPassword gera o hash da senha e salva no struct.
func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

// CheckPassword compara a senha fornecida com o hash armazenado.
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
