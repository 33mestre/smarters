// Package database inicializa a conexão com o banco de dados e realiza a migração dos modelos para o projeto Smarters.
package database

import (
	"log"

	"github.com/33mestre/smarters/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase inicializa a conexão com o banco de dados e realiza a migração dos modelos
func ConnectDatabase() {
	// Conecta ao banco de dados SQLite
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database! %v", err)
	}

	DB = database

	// Migra os modelos para criar as tabelas no banco de dados
	if err := database.AutoMigrate(
		&models.MessageReceived{},
		&models.Message{},
		&models.User{},
		&models.Postback{},
		&models.MessageSent{},
		&models.Content{},
		&models.Attachment{},
		&models.Payload{},
		&models.Button{},
		&models.LogMessage{}, // Adiciona a migração para LogMessage
		&models.Messaging{},  // Adiciona a migração para Messaging
	); err != nil {
		log.Fatalf("Failed to migrate database! %v", err)
	}
}
