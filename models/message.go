package models

// Message representa os detalhes de uma mensagem de texto
// @description Estrutura que representa os detalhes de uma mensagem de texto
type Message struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	MID  string `json:"mid" example:"mid.1457764197618:41d102a3e1ae206a38"` // ID da mensagem
	Text string `json:"text" example:"ping"`                                // Texto da mensagem
}
