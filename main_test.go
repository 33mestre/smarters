package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHandleMessages(t *testing.T) {
	// Configura o Gin para modo de teste
	gin.SetMode(gin.TestMode)

	// Cria o roteador Gin
	router := gin.Default()
	router.POST("/webhook", handleMessages)

	// Corpo da requisição de teste
	reqBody := []byte(`{"id": "PAGE_ID", "time": 1458692752478, "messaging": [{"sender": {"id": "USER_ID"}, "recipient": {"id": "PAGE_ID"}, "timestamp": 1762902671, "message": {"mid": "mid.1457764197618:41d102a3e1ae206a38", "text": "ping"}}]}`)
	req, err := http.NewRequest("POST", "/webhook", bytes.NewBuffer(reqBody))
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
	expected := `{"recipient":{"id":"USER_ID"},"message":{"text":"pong"}}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
