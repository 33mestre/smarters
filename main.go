package main

import (
	"github.com/33mestre/smarters/database"
	_ "github.com/33mestre/smarters/docs"
	"github.com/33mestre/smarters/routes"
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

func main() {
	// Conecta ao banco de dados
	database.ConnectDatabase()

	// Configura o roteador Gin
	r := gin.Default()

	// Rota para a documentação Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Registra as rotas da API
	routes.RegisterRoutes(r)

	// Roda o servidor
	r.Run(":8080")
}
