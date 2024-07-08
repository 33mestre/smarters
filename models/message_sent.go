package models

// MessageSent representa a mensagem enviada ao Messenger
// @description Estrutura que representa uma mensagem enviada ao Messenger
type MessageSent struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	RecipientID uint    `json:"recipient_id"` // Chave estrangeira para User
	Recipient   User    `json:"recipient"`    // Informações do destinatário
	MessageID   uint    `json:"message_id"`   // Chave estrangeira para Content
	Message     Content `json:"message"`      // Conteúdo da mensagem
}
