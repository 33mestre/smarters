// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/webhook": {
            "post": {
                "description": "Recebe mensagens do Messenger e responde com texto ou botões",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Processa mensagens recebidas do Messenger",
                "parameters": [
                    {
                        "description": "Mensagem recebida",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.MessageReceived"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.MessageSent"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Attachment": {
            "description": "Estrutura que representa um anexo na mensagem",
            "type": "object",
            "properties": {
                "payload": {
                    "description": "Payload do anexo",
                    "allOf": [
                        {
                            "$ref": "#/definitions/main.Payload"
                        }
                    ]
                },
                "type": {
                    "description": "Tipo de anexo",
                    "type": "string",
                    "example": "template"
                }
            }
        },
        "main.Button": {
            "description": "Estrutura que representa um botão no payload",
            "type": "object",
            "properties": {
                "payload": {
                    "description": "Payload do botão",
                    "type": "string",
                    "example": "START_PAYLOAD"
                },
                "title": {
                    "description": "Título do botão",
                    "type": "string",
                    "example": "Começar"
                },
                "type": {
                    "description": "Tipo do botão",
                    "type": "string",
                    "example": "postback"
                }
            }
        },
        "main.Content": {
            "description": "Estrutura que representa o conteúdo da mensagem enviada",
            "type": "object",
            "properties": {
                "attachment": {
                    "description": "Anexo da mensagem",
                    "allOf": [
                        {
                            "$ref": "#/definitions/main.Attachment"
                        }
                    ]
                },
                "text": {
                    "description": "Texto da mensagem",
                    "type": "string"
                }
            }
        },
        "main.ErrorResponse": {
            "description": "Estrutura para resposta de erro",
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "main.Message": {
            "description": "Estrutura que representa os detalhes de uma mensagem de texto",
            "type": "object",
            "properties": {
                "mid": {
                    "description": "ID da mensagem",
                    "type": "string",
                    "example": "mid.1457764197618:41d102a3e1ae206a38"
                },
                "text": {
                    "description": "Texto da mensagem",
                    "type": "string",
                    "example": "ping"
                }
            }
        },
        "main.MessageReceived": {
            "description": "Estrutura que representa uma mensagem recebida do Messenger.",
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID da página",
                    "type": "string",
                    "example": "PAGE_ID"
                },
                "messaging": {
                    "description": "Array de mensagens",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.Messaging"
                    }
                },
                "time": {
                    "description": "Timestamp da mensagem",
                    "type": "integer",
                    "example": 1458692752478
                }
            }
        },
        "main.MessageSent": {
            "description": "Estrutura que representa uma mensagem enviada ao Messenger",
            "type": "object",
            "properties": {
                "message": {
                    "description": "Conteúdo da mensagem",
                    "allOf": [
                        {
                            "$ref": "#/definitions/main.Content"
                        }
                    ]
                },
                "recipient": {
                    "description": "Informações do destinatário",
                    "allOf": [
                        {
                            "$ref": "#/definitions/main.User"
                        }
                    ]
                }
            }
        },
        "main.Messaging": {
            "description": "Estrutura que contém os detalhes da mensagem",
            "type": "object",
            "properties": {
                "message": {
                    "description": "Detalhes da mensagem de texto",
                    "allOf": [
                        {
                            "$ref": "#/definitions/main.Message"
                        }
                    ]
                },
                "postback": {
                    "description": "Detalhes do postback",
                    "allOf": [
                        {
                            "$ref": "#/definitions/main.Postback"
                        }
                    ]
                },
                "recipient": {
                    "description": "Informações do destinatário",
                    "allOf": [
                        {
                            "$ref": "#/definitions/main.User"
                        }
                    ]
                },
                "sender": {
                    "description": "Informações do remetente",
                    "allOf": [
                        {
                            "$ref": "#/definitions/main.User"
                        }
                    ]
                },
                "timestamp": {
                    "description": "Timestamp da mensagem",
                    "type": "integer",
                    "example": 1458692752478
                }
            }
        },
        "main.Payload": {
            "description": "Estrutura que representa o payload de um anexo",
            "type": "object",
            "properties": {
                "buttons": {
                    "description": "Botões do template",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.Button"
                    }
                },
                "template_type": {
                    "description": "Tipo de template",
                    "type": "string",
                    "example": "button"
                },
                "text": {
                    "description": "Texto do template",
                    "type": "string",
                    "example": "O que você gostaria de fazer a seguir?"
                }
            }
        },
        "main.Postback": {
            "description": "Estrutura que representa os detalhes de um postback",
            "type": "object",
            "properties": {
                "mid": {
                    "description": "ID do postback",
                    "type": "string",
                    "example": "mid.1457764197618:41d102a3e1ae206a38"
                },
                "payload": {
                    "description": "Payload do postback",
                    "type": "string",
                    "example": "START_PAYLOAD"
                }
            }
        },
        "main.User": {
            "description": "Estrutura que representa um usuário",
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID do usuário",
                    "type": "string",
                    "example": "USER_ID"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Smarters API",
	Description:      "Estrutura que representa um botão no payload",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
