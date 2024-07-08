package main

import (
	"log"
	"os"

	"github.com/33mestre/smarters/database"
	_ "github.com/33mestre/smarters/docs"
	"github.com/33mestre/smarters/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	// Carrega as variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Println("Nenhum arquivo .env encontrado")
	}

	// Configura o modo do Gin
	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)

	// Conecta ao banco de dados
	database.ConnectDatabase()

	// Configura o roteador Gin
	r := gin.Default()

	// Rota para a documentação Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Registra as rotas da API
	routes.RegisterRoutes(r)

	// Define a porta do servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Exibe um cabeçalho no log
	separator := "______________________________________________________________________________"
	log.Println(separator)
	log.Println("    _____ __         __                    ______                          _")
	log.Println("   / ___// /_  ___  / /________  ____     / ____/__  ______________ ______(_)")
	log.Println("   \\__ \\/ __ \\/ _ \\/ / ___/ __ \\/ __ \\   / /_  / _ \\/ ___/ ___/ __ `/ ___/ /")
	log.Println("  ___/ / / / /  __/ (__  ) /_/ / / / /  / __/ /  __/ /  / /  / /_/ / /  / /")
	log.Println(" /____/_/ /_/\\___/_/____/\\____/_/ /_/  /_/    \\___/_/  /_/   \\__,_/_/  /_/")
	log.Println(separator)
	log.Println(separator)
	log.Println("Nome do Projeto: Smarters")
	log.Println("Missão do Projeto: Receber e tratar mensagens enviadas pelo Messenger da Meta e enviar respostas em texto ou listas de botões.")
	log.Println(separator)

	// Inicia o servidor
	log.Printf("Servidor está executando na porta %s\n", port)
	r.Run(":" + port)
}
