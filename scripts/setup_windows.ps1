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