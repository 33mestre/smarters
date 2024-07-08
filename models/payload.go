package models

// Payload representa o payload de um anexo
// @description Estrutura que representa o payload de um anexo
type Payload struct {
	ID           uint     `gorm:"primaryKey" json:"id"`
	TemplateType string   `json:"template_type" example:"button"`                        // Tipo de template
	Text         string   `json:"text" example:"O que você gostaria de fazer a seguir?"` // Texto do template
	Buttons      []Button `json:"buttons"`                                               // Botões do template
}
