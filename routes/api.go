// Package routes registra as rotas da API para o projeto Smarters.
package routes

import (
	"github.com/33mestre/smarters/handlers"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registra todas as rotas da API no roteador Gin.
// @param r *gin.Engine - O roteador Gin onde as rotas serão registradas.
func RegisterRoutes(r *gin.Engine) {
	// Cria um grupo de rotas com o prefixo /api/v1
	api := r.Group("/api/v1")
	{
		// Registra a rota para receber e processar mensagens do Messenger
		// @Summary Recebe e processa mensagens do Messenger
		// @Description Recebe mensagens enviadas pelo Messenger da Meta e responde com texto ou botões
		// @Accept json
		// @Produce json
		// @Param message body models.MessageReceived true "Mensagem recebida"
		// @Success 200 {object} models.MessageSent "Resposta enviada ao usuário"
		// @Failure 400 {object} models.ErrorResponse "Requisição inválida"
		// @Failure 404 {object} models.ErrorResponse "Rota não encontrada"
		// @Failure 500 {object} models.ErrorResponse "Erro interno do servidor"
		// @Router /api/v1/webhook [post]
		api.POST("/webhook", handlers.HandleMessages)
	}
}
