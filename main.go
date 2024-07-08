package main

import (
    "encoding/json"
    "log"
    "net/http"
)

// Estruturas para mensagens recebidas
type MessageReceived struct {
    ID        string      `json:"id"`
    Time      int64       `json:"time"`
    Messaging []Messaging `json:"messaging"`
}

type Messaging struct {
    Sender    User     `json:"sender"`
    Recipient User     `json:"recipient"`
    Timestamp int64    `json:"timestamp"`
    Message   *Message `json:"message,omitempty"`
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

// Estruturas para mensagens enviadas
type MessageSent struct {
    Recipient User    `json:"recipient"`
    Message   Content `json:"message"`
}

type Content struct {
    Text      *string     `json:"text,omitempty"`
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

func handleMessages(w http.ResponseWriter, r *http.Request) {
    var msgReceived MessageReceived
    if err := json.NewDecoder(r.Body).Decode(&msgReceived); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    log.Printf("Received message: %+v", msgReceived)
    // Enviar mensagem para outro endpoint
    forwardMessage(msgReceived)

    // Responder ao usuário
    responseMsg := createResponse(msgReceived)
    sendMessage(responseMsg)
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
    http.HandleFunc("/webhook", handleMessages)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
