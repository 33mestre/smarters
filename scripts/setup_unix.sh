#!/bin/bash

# Nome do Projeto
PROJECT_NAME="Smarters"

# Missão do Projeto
PROJECT_MISSION="Receber e tratar mensagens enviadas pelo Messenger da Meta e enviar respostas em texto ou listas de botões."

# Linguagem do Projeto
PROJECT_LANGUAGE="Go"

# O que está sendo instalado
INSTALLING="Go, SQLite, golangci-lint e dependências do projeto"

# Separador
SEPARATOR="______________________________________________________________________________"

# Cabeçalho
echo "$SEPARATOR"
echo "    _____ __         __                    ______                          _"
echo "   / ___// /_  ___  / /________  ____     / ____/__  ______________ ______(_)"
echo "   \__ \/ __ \/ _ \/ / ___/ __ \/ __ \   / /_  / _ \/ ___/ ___/ __ \`/ ___/ /"
echo "  ___/ / / / /  __/ (__  ) /_/ / / / /  / __/ /  __/ /  / /  / /_/ / /  / /"
echo " /____/_/ /_/\___/_/____/\____/_/ /_/  /_/    \___/_/  /_/   \__,_/_/  /_/"
echo "$SEPARATOR"
echo "$SEPARATOR"
echo "Nome do Projeto: $PROJECT_NAME"
echo "Missão do Projeto: $PROJECT_MISSION"
echo "Linguagem do Projeto: $PROJECT_LANGUAGE"
echo "Instalando: $INSTALLING"
echo "$SEPARATOR"

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

# Função para instalar o golangci-lint
install_golangci_lint() {
    echo "Instalando golangci-lint..."
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.45.2
    echo "export PATH=\$PATH:$(go env GOPATH)/bin" >> ~/.profile
    source ~/.profile
}

# Função para instalar o SQLite no Unix
install_sqlite() {
    echo "Instalando SQLite..."
    if [ "$OS" == "Darwin" ]; then
        brew install sqlite3
        echo "$SEPARATOR"
    else
        sudo apt-get update
        sudo apt-get install sqlite3 -y
        echo "$SEPARATOR"
    fi
}

# Verifica e instala o Go
if ! command -v go &> /dev/null
then
    install_go
    echo "$SEPARATOR"
else
    echo "Go já está instalado."
    echo "$SEPARATOR"
fi

# Verifica e instala o SQLite
if ! command -v sqlite3 &> /dev/null
then
    install_sqlite
    echo "$SEPARATOR"
else
    echo "SQLite já está instalado."
    echo "$SEPARATOR"
fi

# Verifica e instala o golangci-lint
if ! command -v golangci-lint &> /dev/null
then
    install_golangci_lint
    echo "$SEPARATOR"
else
    echo "golangci-lint já está instalado."
    echo "$SEPARATOR"
fi

# Instala as dependências do Go
echo "Instalando dependências do Go..."
go mod tidy
echo "$SEPARATOR"

# Inicializa o banco de dados
echo "Inicializando o banco de dados..."
go run main.go &
echo "$SEPARATOR"

echo "Configuração concluída."
echo "$SEPARATOR"