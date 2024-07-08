package models

// Content representa o conteúdo da mensagem enviada
// @description Estrutura que representa o conteúdo da mensagem enviada
type Content struct {
	ID           uint        `gorm:"primaryKey" json:"id"`
	Text         *string     `json:"text,omitempty"`       // Texto da mensagem
	AttachmentID uint        `json:"attachment_id"`        // Chave estrangeira para Attachment
	Attachment   *Attachment `json:"attachment,omitempty"` // Anexo da mensagem
}
