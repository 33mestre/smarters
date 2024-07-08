package models

// Messaging representa os detalhes da mensagem
// @description Estrutura que contém os detalhes da mensagem
type Messaging struct {
	Sender    User      `json:"sender"`                            // Informações do remetente
	Recipient User      `json:"recipient"`                         // Informações do destinatário
	Timestamp int64     `json:"timestamp" example:"1458692752478"` // Timestamp da mensagem
	Message   *Message  `json:"message,omitempty"`                 // Detalhes da mensagem de texto
	Postback  *Postback `json:"postback,omitempty"`                // Detalhes do postback
}
