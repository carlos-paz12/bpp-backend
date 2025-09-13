package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(hash)

	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}
