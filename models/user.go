package models

// User representa um usuário
// @description Estrutura que representa um usuário
type User struct {
	ID  uint   `gorm:"primaryKey" json:"id"`
	UID string `json:"uid" example:"USER_ID"` // ID do usuário
}
