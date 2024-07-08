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

## Importando Requisições no Postman

Você pode importar o arquivo postman_collection.json incluído neste repositório para testar o servidor usando o Postman.

Abra o Postman.
Clique em Import no canto superior esquerdo.
Selecione Import File e escolha o arquivo postman_collection.json.
Clique em Import.


## Explicações sobre o Docker
O Dockerfile incluído neste repositório configura um ambiente Docker para executar o servidor Go. Ele copia os arquivos de código para o container, instala as dependências e expõe a porta 8080 para acesso externo.


## Documentação Swagger
A documentação Swagger para esta API está disponível em http://localhost:8080/swagger/index.html quando o servidor está em execução.



