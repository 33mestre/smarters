package models

// MessageSent representa a mensagem enviada ao Messenger
// @description Estrutura que representa uma mensagem enviada ao Messenger
type MessageSent struct {
	Recipient User    `json:"recipient"` // Informações do destinatário
	Message   Content `json:"message"`   // Conteúdo da mensagem
}
