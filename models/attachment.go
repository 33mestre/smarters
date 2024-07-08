package models

// Attachment representa um anexo na mensagem
// @description Estrutura que representa um anexo na mensagem
type Attachment struct {
	Type    string  `json:"type" example:"template"` // Tipo de anexo
	Payload Payload `json:"payload"`                 // Payload do anexo
}
