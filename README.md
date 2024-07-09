# Smarters

Este projeto é um sistema em Go para receber e tratar mensagens enviadas pelo Messenger da Meta e enviar mensagens de resposta em texto ou listas de botões clicáveis.

## Estrutura do Código

```txt
/smarters
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── handlers
│   └── messages.go
├── routes
│   └── api.go
├── scripts
│   └── setup_unix.sh
│   └── setup_windows.ps1
├── database
│   └── database.go
├── models
│   ├── attachment.go
│   ├── button.go
│   ├── content.go
│   ├── doc.go
│   ├── error_response.go
│   ├── message.go
│   ├── message_received.go
│   ├── message_sent.go
│   ├── messaging.go
│   ├── payload.go
│   └── user.go
├── main.go
├── main_test.go
├── go.mod
└── go.sum
```

- `main.go`: Arquivo principal contendo a lógica de recebimento e envio de mensagens.
- `go.mod`: Arquivo de gerenciamento de dependências do Go.
- `main_test.go`: Testes unitários para o servidor.

## Dependências

As principais bibliotecas utilizadas neste projeto são:

- [Gin](https://github.com/gin-gonic/gin): Framework web em Go.
- [GORM](https://gorm.io): ORM para Go.
- [GoDotEnv](https://github.com/joho/godotenv): Carrega variáveis de ambiente a partir de um arquivo `.env`.
- [Swag](https://github.com/swaggo/swag): Gera documentação Swagger a partir de comentários no código.
- [Gin-Swagger](https://github.com/swaggo/gin-swagger): Integração do Swagger com o Gin.

## Instalação

### Clonando o Repositório

Primeiro, clone o repositório:

```sh
git clone git@github.com:33mestre/smarters.git cd smarters
```

### Executando Scripts de Configuração

#### Unix/Linux

Para sistemas Unix/Linux, use o script `setup_unix.sh` para instalar as dependências e configurar o ambiente:

```sh
bash scripts/setup_unix.sh
```

#### Windows

Para sistemas Windows, use o script `setup_windows.bat` para instalar as dependências e configurar o ambiente:

```sh
scripts\setup_windows.bat
```

## Executando Localmente com Docker

Certifique-se de ter o Docker instalado em sua máquina. Para construir e executar o projeto usando Docker, siga os passos abaixo:

1. Construa a imagem Docker:
    
    ```sh
    docker build -t smarters .
   ``` 
    
2. Execute o container Docker:
    
    ```sh
    docker run -p 8080:8080 smarters
    ```

O servidor estará disponível em [http://localhost:8080](http://localhost:8080).

## Testando

Para rodar os testes, utilize o comando:

```sh
go test ./...
```

## Importando Requisições no Postman

Você pode importar o arquivo `postman_collection.json` incluído neste repositório para testar o servidor usando o Postman.

1. Abra o Postman.
2. Clique em Import no canto superior esquerdo.
3. Selecione Import File e escolha o arquivo `postman_collection.json`.
4. Clique em Import.

## Documentação Swagger

A documentação Swagger para esta API está disponível em [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) quando o servidor está em execução.

## CI/CD com GitHub Actions

Este projeto inclui um workflow de CI/CD configurado para ser executado no GitHub Actions. O arquivo de configuração do workflow é `.github/workflows/go.yml`.

O workflow realiza as seguintes etapas:

- Faz o checkout do código.
- Configura o Go.
- Instala as dependências.
- Instala o golangci-lint e gosec.
- Executa verificações de segurança.
- Executa os testes.
- Constrói a aplicação.

### Arquivo de Configuração do GitHub Actions

```yaml
name: Go CI

on:
  push:
    branches: [ "main" ]
  pull_request:
    branches: [ "main" ]

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
    - name: Check out code
      uses: actions/checkout@v2

    - name: Set up Go
      uses: actions/setup-go@v2
      with:
        go-version: 1.18

    - name: Install dependencies
      run: go mod tidy

    - name: Install golangci-lint
      run: |
        curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.45.2
        echo "$(go env GOPATH)/bin" >> $GITHUB_PATH

    - name: Install gosec
      run: |
        go install github.com/securego/gosec/v2/cmd/gosec@v2.15.0
        echo "$(go env GOPATH)/bin" >> $GITHUB_PATH

    - name: Run security checks
      run: gosec ./...

    - name: Run tests
      run: go test ./...

    - name: Build
      run: go build -v ./...
```

## Scripts de Configuração

### Unix/Linux (`scripts/setup_unix.sh`)

```sh
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
```

### Windows (`scripts/setup_windows.ps1`)

```sh
# Nome do Projeto
$PROJECT_NAME = "Smarters"

# Missão do Projeto
$PROJECT_MISSION = "Receber e tratar mensagens enviadas pelo Messenger da Meta e enviar respostas em texto ou listas de botões."

# Linguagem do Projeto
$PROJECT_LANGUAGE = "Go"

# Separador
SEPARATOR="______________________________________________________________________________"

# O que está sendo instalado
$INSTALLING = "Go, SQLite, golangci-lint e dependências do projeto"

# Cabeçalho
Write-Output "$SEPARATOR"
Write-Output "    _____ __         __                    ______                          _"
Write-Output "   / ___// /_  ___  / /________  ____     / ____/__  ______________ ______(_)"
Write-Output "   \__ \/ __ \/ _ \/ / ___/ __ \/ __ \   / /_  / _ \/ ___/ ___/ __ \`/ ___/ /"
Write-Output "  ___/ / / / /  __/ (__  ) /_/ / / / /  / __/ /  __/ /  / /  / /_/ / /  / /"
Write-Output " /____/_/ /_/\___/_/____/\____/_/ /_/  /_/    \___/_/  /_/   \__,_/_/  /_/"
Write-Output "$SEPARATOR"
Write-Output "$SEPARATOR"
Write-Output "Nome do Projeto: $PROJECT_NAME"
Write-Output "Missão do Projeto: $PROJECT_MISSION"
Write-Output "Linguagem do Projeto: $PROJECT_LANGUAGE"
Write-Output "Instalando: $INSTALLING"
Write-Output "$SEPARATOR"

# Detecta o sistema operacional
$OS = Get-WmiObject Win32_OperatingSystem | Select-Object -ExpandProperty Caption
Write-Output "Sistema operacional detectado: $OS"

# Função para instalar o Go no Windows
function Install-Go {
    Write-Output "Instalando Go..."
    Invoke-WebRequest -Uri https://golang.org/dl/go1.20.5.windows-amd64.msi -OutFile go.msi
    Start-Process msiexec.exe -ArgumentList '/i', 'go.msi', '/quiet', '/norestart' -Wait
    [System.Environment]::SetEnvironmentVariable('PATH', $env:PATH + ';C:\Go\bin', [System.EnvironmentVariableTarget]::Machine)
    $env:PATH += ';C:\Go\bin'
}

# Função para instalar o SQLite no Windows
function Install-SQLite {
    Write-Output "Instalando SQLite..."
    choco install sqlite
}

# Função para instalar o golangci-lint
function Install-GolangciLint {
    Write-Output "Instalando golangci-lint..."
    Invoke-WebRequest -Uri https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh -OutFile install-golangci-lint.ps1
    .\install-golangci-lint.ps1 -b $env:GOPATH\bin v1.45.2
    [System.Environment]::SetEnvironmentVariable('PATH', $env:PATH + ';' + $env:GOPATH + '\bin', [System.EnvironmentVariableTarget]::Machine)
    $env:PATH += ';' + $env:GOPATH + '\bin'
}

# Verifica e instala o Go
if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
    Install-Go
    Write-Output "$SEPARATOR"
} else {
    Write-Output "Go já está instalado."
    Write-Output "$SEPARATOR"
}

# Verifica e instala o SQLite
if (-not (Get-Command sqlite3 -ErrorAction SilentlyContinue)) {
    Install-SQLite
    Write-Output "$SEPARATOR"
} else {
    Write-Output "SQLite já está instalado."
    Write-Output "$SEPARATOR"
}

# Verifica e instala o golangci-lint
if (-not (Get-Command golangci-lint -ErrorAction SilentlyContinue)) {
    Install-GolangciLint
    Write-Output "$SEPARATOR"
} else {
    Write-Output "golangci-lint já está instalado."
    Write-Output "$SEPARATOR"
}

# Instala as dependências do Go
Write-Output "Instalando dependências do Go..."
go mod tidy
Write-Output "$SEPARATOR"

# Inicializa o banco de dados
Write-Output "Inicializando o banco de dados..."
Start-Process -NoNewWindow -FilePath "go" -ArgumentList "run main.go"
Write-Output "$SEPARATOR"

Write-Output "Configuração concluída."
Write-Output "$SEPARATOR"
```

## Executando

Para iniciar o servidor localmente, use:

```sh
go run main.go
```

## Gerando Documentação

Para gerar a documentação usando o Swag, use:

```sh
swag init
```

A documentação gerada estará disponível em [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) quando o servidor estiver em execução.

## Documentação GoDoc

Para visualizar a documentação gerada pelo GoDoc, execute o seguinte comando:

```sh
godoc -http=:6060
```

A documentação estará disponível em [http://localhost:6060/pkg/github.com/33mestre/smarters/](http://localhost:6060/pkg/github.com/33mestre/smarters/) enquanto o servidor estiver em execução.

## Visualizando os Dados Persistidos

Para verificar se os dados estão sendo corretamente persistidos no banco de dados SQLite, você pode usar o comando `sqlite3` no terminal. Siga os passos abaixo:

1. Abra o terminal.
    
2. Navegue até o diretório onde o arquivo do banco de dados está localizado:
    
    ```sh
    cd /path/to/your/project
    ```
    
3. Inicie o cliente SQLite:
    
    ```sh
    sqlite3 test.db
    ```
    
4. Liste as tabelas disponíveis no banco de dados:
    
    ```sh
    .tables
    ```
    
5. Visualize os dados de uma tabela específica. Por exemplo, para ver os dados da tabela `message_receiveds`:
    
    ```sh
    SELECT * FROM message_receiveds;
    ```
    
6. Para sair do cliente SQLite, digite:
    
    ```sh
    .exit
    ```



## Licença

Este projeto é licenciado sob os termos da licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## Contribuindo

Sinta-se à vontade para abrir issues e pull requests. Todas as contribuições são bem-vindas.