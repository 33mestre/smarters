package models

// Attachment representa um anexo na mensagem
// @description Estrutura que representa um anexo na mensagem
type Attachment struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	Type      string  `json:"type" example:"template"` // Tipo de anexo
	PayloadID uint    `json:"payload_id"`              // Chave estrangeira para Payload
	Payload   Payload `json:"payload"`                 // Payload do anexo
}
