#!/bin/sh

# Instalar dependências Go
echo "Instalando dependências Go..."
go mod tidy

# Instalar Swag CLI
echo "Instalando Swag CLI..."
go install github.com/swaggo/swag/cmd/swag@latest

# Gerar documentação Swagger
echo "Gerando documentação Swagger..."
swag init

# Conectar ao banco de dados e rodar migrações
echo "Conectando ao banco de dados e rodando migrações..."
go run main.go
