package models

// MessageReceived representa a mensagem recebida do Messenger
// @description Estrutura que representa uma mensagem recebida do Messenger.
type MessageReceived struct {
	ID        string      `gorm:"primaryKey" json:"id" example:"PAGE_ID"` // ID da página
	Time      int64       `json:"time" example:"1458692752478"`           // Timestamp da mensagem
	Messaging []Messaging `json:"messaging"`                              // Array de mensagens
}
