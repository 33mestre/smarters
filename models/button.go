package models

// Button representa um botão no payload
// @description Estrutura que representa um botão no payload
type Button struct {
	Type    string `json:"type" example:"postback"`         // Tipo do botão
	Title   string `json:"title" example:"Começar"`         // Título do botão
	Payload string `json:"payload" example:"START_PAYLOAD"` // Payload do botão
}
