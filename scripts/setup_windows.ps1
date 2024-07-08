# Nome do Projeto
$PROJECT_NAME = "Smarters"

# Missão do Projeto
$PROJECT_MISSION = "Receber e tratar mensagens enviadas pelo Messenger da Meta e enviar respostas em texto ou listas de botões."

# Linguagem do Projeto
$PROJECT_LANGUAGE = "Go"

# O que está sendo instalado
$INSTALLING = "Go, SQLite e dependências do projeto"

# Cabeçalho
Write-Output "===================================="
Write-Output "Nome do Projeto: $PROJECT_NAME"
Write-Output "Missão do Projeto: $PROJECT_MISSION"
Write-Output "Linguagem do Projeto: $PROJECT_LANGUAGE"
Write-Output "Instalando: $INSTALLING"
Write-Output "===================================="

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

# Verifica e instala o Go
if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
    Install-Go
} else {
    Write-Output "Go já está instalado."
}

# Verifica e instala o SQLite
if (-not (Get-Command sqlite3 -ErrorAction SilentlyContinue)) {
    Install-SQLite
} else {
    Write-Output "SQLite já está instalado."
}

# Instala as dependências do Go
Write-Output "Instalando dependências do Go..."
go mod tidy

# Inicializa o banco de dados
Write-Output "Inicializando o banco de dados..."
Start-Process -NoNewWindow -FilePath "go" -ArgumentList "run main.go"

Write-Output "Configuração concluída."
