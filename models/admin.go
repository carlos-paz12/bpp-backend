package models

type Admin struct {
	ID     int64 `json:"id"`      //!< ID do registro de admin.
	UserID int64 `json:"user_id"` //!< ID do usuário associado.
}
