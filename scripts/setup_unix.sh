#!/bin/bash

# Nome do Projeto
PROJECT_NAME="Smarters"

# Missão do Projeto
PROJECT_MISSION="Receber e tratar mensagens enviadas pelo Messenger da Meta e enviar respostas em texto ou listas de botões."

# Linguagem do Projeto
PROJECT_LANGUAGE="Go"

# O que está sendo instalado
INSTALLING="Go, SQLite e dependências do projeto"

# Cabeçalho
echo "===================================="
echo "Nome do Projeto: $PROJECT_NAME"
echo "Missão do Projeto: $PROJECT_MISSION"
echo "Linguagem do Projeto: $PROJECT_LANGUAGE"
echo "Instalando: $INSTALLING"
echo "===================================="

# Detecta o sistema operacional
OS=$(uname -s)
echo "Sistema operacional detectado: $OS"

# Função para instalar o Go no Unix
install_go() {
    echo "Instalando Go..."
    if [ "$OS" == "Darwin" ]; then
        curl -OL https://golang.org/dl/go1.20.5.darwin-arm64.tar.gz
        sudo tar -C /usr/local -xzf go1.20.5.darwin-arm64.tar.gz
        echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.zshrc
        source ~/.zshrc
    else
        curl -OL https://golang.org/dl/go1.20.5.linux-amd64.tar.gz
        sudo tar -C /usr/local -xzf go1.20.5.linux-amd64.tar.gz
        echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.profile
        source ~/.profile
    fi
}

# Função para instalar o SQLite no Unix
install_sqlite() {
    echo "Instalando SQLite..."
    if [ "$OS" == "Darwin" ]; then
        brew install sqlite3
    else
        sudo apt-get update
        sudo apt-get install sqlite3 -y
    fi
}

# Verifica e instala o Go
if ! command -v go &> /dev/null
then
    install_go
else
    echo "Go já está instalado."
fi

# Verifica e instala o SQLite
if ! command -v sqlite3 &> /dev/null
then
    install_sqlite
else
    echo "SQLite já está instalado."
fi

# Instala as dependências do Go
echo "Instalando dependências do Go..."
go mod tidy

# Inicializa o banco de dados
echo "Inicializando o banco de dados..."
go run main.go &

echo "Configuração concluída."
