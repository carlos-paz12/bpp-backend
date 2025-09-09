package models

import "golang.org/x/crypto/bcrypt"

type Bolsista struct {
    Username string
    Password string
}

func (u *Bolsista) SetPassword(password string) error {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    u.Password = string(hash)
    return nil
}

func (u *Bolsista) CheckPassword(password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
    return err == nil
}
