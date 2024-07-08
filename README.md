# Smarters

Este projeto é um sistema em Go para receber e tratar mensagens enviadas pelo Messenger da Meta e enviar mensagens de resposta em texto ou listas de botões clicáveis.

## Estrutura do Código

- `main.go`: Arquivo principal contendo a lógica de recebimento e envio de mensagens.
- `go.mod`: Arquivo de gerenciamento de dependências do Go.

## Executando Localmente com Docker

Certifique-se de ter o Docker instalado em sua máquina. Para construir e executar o projeto usando Docker, siga os passos abaixo:

1. Construa a imagem Docker:
   ```sh
   docker build -t smarters .
   docker run -p 8080:8080 smarters
   ````

O servidor estará disponível em http://localhost:8080.

## Testando

Para rodar os testes, utilize o comando:

```sh
go test ./...
```




