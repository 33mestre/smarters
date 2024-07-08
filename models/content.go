package models

// Content representa o conteúdo da mensagem enviada
// @description Estrutura que representa o conteúdo da mensagem enviada
type Content struct {
	Text       *string     `json:"text,omitempty"`       // Texto da mensagem
	Attachment *Attachment `json:"attachment,omitempty"` // Anexo da mensagem
}
