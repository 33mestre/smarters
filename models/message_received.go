// Package models contém os modelos de dados para o projeto Smarters.
package models

// MessageReceived representa a mensagem recebida do Messenger
// @description Estrutura que representa uma mensagem recebida do Messenger.
type MessageReceived struct {
	ID        uint        `gorm:"primaryKey;autoIncrement" json:"id" example:"1"` // ID da mensagem recebida
	PageID    string      `json:"page_id" example:"PAGE_ID"`                      // ID da página
	Time      int64       `json:"time" example:"1458692752478"`                   // Timestamp da mensagem
	Messaging []Messaging `json:"messaging"`                                      // Array de mensagens
}
