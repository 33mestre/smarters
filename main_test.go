package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/33mestre/smarters/database"
	"github.com/33mestre/smarters/routes"
	"github.com/gin-gonic/gin"
)

func TestHandleMessages(t *testing.T) {
	// Configura o Gin para modo de teste
	gin.SetMode(gin.TestMode)

	// Inicializa o banco de dados
	database.ConnectDatabase()

	// Cria o roteador Gin
	router := gin.Default()
	routes.RegisterRoutes(router)

	// Corpo da requisição de teste
	reqBody := []byte(`{
		"id": "PAGE_ID", 
		"time": 1458692752478, 
		"messaging": [
			{
				"sender": {"id": 123}, 
				"recipient": {"id": 1111}, 
				"timestamp": 1762902671, 
				"message": {
					"mid": "mid.1457764197618:41d102a3e1ae206a38", 
					"text": "ping"
				}
			}
		]
	}`)
	req, err := http.NewRequest("POST", "/api/v1/webhook", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	// Cria um gravador de respostas HTTP
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Verifica o código de status da resposta
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Verifica o corpo da resposta
	expected := `{"id":0,"recipient_id":0,"recipient":{"id":123,"uid":""},"message_id":0,"message":{"id":0,"text":"pong","attachment_id":0}}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
