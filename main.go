package main

import (
	"log"
	"net/http"

	_ "github.com/33mestre/smarters/docs"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Smarters API
// @version 1.0
// @description API para receber e tratar mensagens enviadas pelo Messenger da Meta.

// @host localhost:8080
// @BasePath /

// MessageReceived representa a mensagem recebida do Messenger
type MessageReceived struct {
	ID        string      `json:"id"`
	Time      int64       `json:"time"`
	Messaging []Messaging `json:"messaging"`
}

type Messaging struct {
	Sender    User      `json:"sender"`
	Recipient User      `json:"recipient"`
	Timestamp int64     `json:"timestamp"`
	Message   *Message  `json:"message,omitempty"`
	Postback  *Postback `json:"postback,omitempty"`
}

type User struct {
	ID string `json:"id"`
}

type Message struct {
	MID  string `json:"mid"`
	Text string `json:"text"`
}

type Postback struct {
	MID     string `json:"mid"`
	Payload string `json:"payload"`
}

// MessageSent representa a mensagem enviada ao Messenger
type MessageSent struct {
	Recipient User    `json:"recipient"`
	Message   Content `json:"message"`
}

type Content struct {
	Text       *string     `json:"text,omitempty"`
	Attachment *Attachment `json:"attachment,omitempty"`
}

type Attachment struct {
	Type    string  `json:"type"`
	Payload Payload `json:"payload"`
}

type Payload struct {
	TemplateType string   `json:"template_type"`
	Text         string   `json:"text"`
	Buttons      []Button `json:"buttons"`
}

type Button struct {
	Type    string `json:"type"`
	Title   string `json:"title"`
	Payload string `json:"payload"`
}

// handleMessages processa as mensagens recebidas
// @Summary Processa mensagens recebidas do Messenger
// @Description Recebe mensagens do Messenger e responde com texto ou botões
// @Accept  json
// @Produce  json
// @Param   message body MessageReceived true "Mensagem recebida"
// @Success 200 {object} MessageSent
// @Router /webhook [post]
func handleMessages(c *gin.Context) {
	var msgReceived MessageReceived
	if err := c.BindJSON(&msgReceived); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	// Implementar o envio da mensagem para outro endpoint
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
