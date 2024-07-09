// Package models contém os modelos de dados para o projeto Smarters.
package models

// Messaging representa os detalhes da mensagem
// @description Estrutura que contém os detalhes da mensagem
type Messaging struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	SenderID          uint      `json:"sender_id"`                         // Chave estrangeira para User
	Sender            User      `json:"sender"`                            // Informações do remetente
	RecipientID       uint      `json:"recipient_id"`                      // Chave estrangeira para User
	Recipient         User      `json:"recipient"`                         // Informações do destinatário
	Timestamp         int64     `json:"timestamp" example:"1458692752478"` // Timestamp da mensagem
	MessageID         uint      `json:"message_id"`                        // Chave estrangeira para Message
	Message           *Message  `json:"message,omitempty"`                 // Detalhes da mensagem de texto
	PostbackID        uint      `json:"postback_id"`                       // Chave estrangeira para Postback
	Postback          *Postback `json:"postback,omitempty"`                // Detalhes do postback
	MessageReceivedID uint      `json:"message_received_id"`               // Chave estrangeira para MessageReceived
}
