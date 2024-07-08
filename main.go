package main

import (
	"log"
	"net/http"

	_ "github.com/33mestre/smarters/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Smarters API
// @version 1.0
// @description API para receber e tratar mensagens enviadas pelo Messenger da Meta.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /

// Estrutura para resposta de erro
// @description Estrutura para resposta de erro
type ErrorResponse struct {
	Error string `json:"error"`
}

// MessageReceived representa a mensagem recebida do Messenger
// @description Estrutura que representa uma mensagem recebida do Messenger.
type MessageReceived struct {
	ID        string      `json:"id" example:"PAGE_ID"`         // ID da página
	Time      int64       `json:"time" example:"1458692752478"` // Timestamp da mensagem
	Messaging []Messaging `json:"messaging"`                    // Array de mensagens
}

// Messaging representa os detalhes da mensagem
// @description Estrutura que contém os detalhes da mensagem
type Messaging struct {
	Sender    User      `json:"sender"`                            // Informações do remetente
	Recipient User      `json:"recipient"`                         // Informações do destinatário
	Timestamp int64     `json:"timestamp" example:"1458692752478"` // Timestamp da mensagem
	Message   *Message  `json:"message,omitempty"`                 // Detalhes da mensagem de texto
	Postback  *Postback `json:"postback,omitempty"`                // Detalhes do postback
}

// User representa um usuário
// @description Estrutura que representa um usuário
type User struct {
	ID string `json:"id" example:"USER_ID"` // ID do usuário
}

// Message representa os detalhes de uma mensagem de texto
// @description Estrutura que representa os detalhes de uma mensagem de texto
type Message struct {
	MID  string `json:"mid" example:"mid.1457764197618:41d102a3e1ae206a38"` // ID da mensagem
	Text string `json:"text" example:"ping"`                                // Texto da mensagem
}

// Postback representa os detalhes de um postback
// @description Estrutura que representa os detalhes de um postback
type Postback struct {
	MID     string `json:"mid" example:"mid.1457764197618:41d102a3e1ae206a38"` // ID do postback
	Payload string `json:"payload" example:"START_PAYLOAD"`                    // Payload do postback
}

// MessageSent representa a mensagem enviada ao Messenger
// @description Estrutura que representa uma mensagem enviada ao Messenger
type MessageSent struct {
	Recipient User    `json:"recipient"` // Informações do destinatário
	Message   Content `json:"message"`   // Conteúdo da mensagem
}

// Content representa o conteúdo da mensagem enviada
// @description Estrutura que representa o conteúdo da mensagem enviada
type Content struct {
	Text       *string     `json:"text,omitempty"`       // Texto da mensagem
	Attachment *Attachment `json:"attachment,omitempty"` // Anexo da mensagem
}

// Attachment representa um anexo na mensagem
// @description Estrutura que representa um anexo na mensagem
type Attachment struct {
	Type    string  `json:"type" example:"template"` // Tipo de anexo
	Payload Payload `json:"payload"`                 // Payload do anexo
}

// Payload representa o payload de um anexo
// @description Estrutura que representa o payload de um anexo
type Payload struct {
	TemplateType string   `json:"template_type" example:"button"`                        // Tipo de template
	Text         string   `json:"text" example:"O que você gostaria de fazer a seguir?"` // Texto do template
	Buttons      []Button `json:"buttons"`                                               // Botões do template
}

// Button representa um botão no payload
// @description Estrutura que representa um botão no payload
type Button struct {
	Type    string `json:"type" example:"postback"`         // Tipo do botão
	Title   string `json:"title" example:"Começar"`         // Título do botão
	Payload string `json:"payload" example:"START_PAYLOAD"` // Payload do botão
}

// handleMessages processa as mensagens recebidas
// @Summary Processa mensagens recebidas do Messenger
// @Description Recebe mensagens do Messenger e responde com texto ou botões
// @Accept  json
// @Produce  json
// @Param   message body MessageReceived true "Mensagem recebida"
// @Success 200 {object} MessageSent
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /webhook [post]
func handleMessages(c *gin.Context) {
	var msgReceived MessageReceived
	if err := c.BindJSON(&msgReceived); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	log.Printf("Received message: %+v", msgReceived)
	// Enviar mensagem para outro endpoint
	forwardMessage(msgReceived)

	// Responder ao usuário
	responseMsg := createResponse(msgReceived)
	sendMessage(responseMsg)
	c.JSON(http.StatusOK, responseMsg)
}

func forwardMessage(msg MessageReceived) {
	// Implemente o envio da mensagem para outro endpoint
}

func createResponse(msg MessageReceived) MessageSent {
	userId := msg.Messaging[0].Sender.ID
	if msg.Messaging[0].Message != nil {
		text := "pong"
		return MessageSent{
			Recipient: User{ID: userId},
			Message:   Content{Text: &text},
		}
	} else if msg.Messaging[0].Postback != nil {
		return MessageSent{
			Recipient: User{ID: userId},
			Message: Content{
				Attachment: &Attachment{
					Type: "template",
					Payload: Payload{
						TemplateType: "button",
						Text:         "O que você gostaria de fazer a seguir?",
						Buttons: []Button{
							{Type: "postback", Title: "Começar", Payload: "START_PAYLOAD"},
							{Type: "postback", Title: "Ajuda", Payload: "HELP_PAYLOAD"},
						},
					},
				},
			},
		}
	}
	return MessageSent{}
}

func sendMessage(msg MessageSent) {
	// Implemente o envio da mensagem de resposta ao usuário
	log.Printf("Sending message: %+v", msg)
}

func main() {
	r := gin.Default()

	// Rota para a documentação Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/webhook", handleMessages)
	r.Run(":8080")
}
