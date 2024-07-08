// Package handlers contains the HTTP request handlers for the Smarters project.
package handlers

import (
	"log"
	"net/http"

	"github.com/33mestre/smarters/database"
	"github.com/33mestre/smarters/models"
	"github.com/gin-gonic/gin"
)

// HandleMessages processes incoming messages from Messenger.
// @Summary Processes incoming messages from Messenger
// @Description Receives messages from Messenger and responds with text or buttons
// @Accept json
// @Produce json
// @Param message body models.MessageReceived true "Received message"
// @Success 200 {object} models.MessageSent
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 404 {object} models.ErrorResponse "Not Found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /api/v1/webhook [post]
func HandleMessages(c *gin.Context) {
	var msgReceived models.MessageReceived
	if err := c.BindJSON(&msgReceived); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	log.Printf("Received message: %+v", msgReceived)
	// Save the received message to the database
	database.DB.Create(&msgReceived)

	// Respond to the user
	responseMsg := createResponse(msgReceived)
	sendMessage(responseMsg)
	c.JSON(http.StatusOK, responseMsg)
}

// forwardMessage sends the received message to another endpoint.
// This function is a stub and should be implemented as needed.
func forwardMessage(msg models.MessageReceived) {
	// Implement sending the message to another endpoint
}

// createResponse creates a response for the received message.
// Returns a MessageSent struct to be sent back to the user.
func createResponse(msg models.MessageReceived) models.MessageSent {
	userId := msg.Messaging[0].Sender.ID
	if msg.Messaging[0].Message != nil {
		text := "pong"
		return models.MessageSent{
			Recipient: models.User{ID: userId},
			Message:   models.Content{Text: &text},
		}
	} else if msg.Messaging[0].Postback != nil {
		return models.MessageSent{
			Recipient: models.User{ID: userId},
			Message: models.Content{
				Attachment: &models.Attachment{
					Type: "template",
					Payload: models.Payload{
						TemplateType: "button",
						Text:         "O que você gostaria de fazer a seguir?",
						Buttons: []models.Button{
							{Type: "postback", Title: "Começar", Payload: "START_PAYLOAD"},
							{Type: "postback", Title: "Ajuda", Payload: "HELP_PAYLOAD"},
						},
					},
				},
			},
		}
	}
	return models.MessageSent{}
}

// sendMessage sends the response message to the user.
// This function is a stub and should be implemented as needed.
func sendMessage(msg models.MessageSent) {
	// Implement sending the response message to the user
	log.Printf("Sending message: %+v", msg)
}
