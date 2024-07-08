package models

// Postback representa os detalhes de um postback
// @description Estrutura que representa os detalhes de um postback
type Postback struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	MID     string `json:"mid" example:"mid.1457764197618:41d102a3e1ae206a38"` // ID do postback
	Payload string `json:"payload" example:"START_PAYLOAD"`                    // Payload do postback
}
