// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Caio Everest",
            "email": "caioeverest@edu.unirio.br"
        },
        "license": {
            "name": "MIT",
            "url": "https://github.com/caioeverest/fed-its/license"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/call": {
            "post": {
                "description": "Request a method from a provider or a group of providers and return the first response received",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orquestrator"
                ],
                "summary": "Request a method",
                "parameters": [
                    {
                        "description": "Payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.CallRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/itserrors.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/itserrors.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/itserrors.Error"
                        }
                    }
                }
            }
        },
        "/method": {
            "get": {
                "description": "List methods",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "method"
                ],
                "summary": "List methods",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Method"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/itserrors.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/itserrors.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/itserrors.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new method",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "method"
                ],
                "summary": "Create a new method",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Method"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Method"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/itserrors.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/itserrors.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/itserrors.Error"
                        }
                    }
                }
            }
        },
        "/method/{method}": {
            "get": {
                "description": "Get method",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "method"
                ],
                "summary": "Get method",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Method",
                        "name": "method",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Method"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/itserrors.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/itserrors.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/itserrors.Error"
                        }
                    }
                }
            }
        },
        "/provider": {
            "post": {
                "description": "Create a new provider",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "provider"
                ],
                "summary": "Create a new provider",
                "parameters": [
                    {
                        "description": "Provider",
                        "name": "provider",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Provider"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Provider"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/itserrors.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/itserrors.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/itserrors.Error"
                        }
                    }
                }
            }
        },
        "/provider/list/{method}": {
            "get": {
                "description": "List providers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "provider"
                ],
                "summary": "List providers",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Provider method",
                        "name": "method",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Provider"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/itserrors.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/itserrors.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/itserrors.Error"
                        }
                    }
                }
            }
        },
        "/provider/{slug}": {
            "get": {
                "description": "Get a provider",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "provider"
                ],
                "summary": "Get a provider",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Provider slug",
                        "name": "slug",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Provider"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/itserrors.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/itserrors.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/itserrors.Error"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a provider",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "provider"
                ],
                "summary": "Delete a provider",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Provider slug",
                        "name": "slug",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Provider"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/itserrors.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/itserrors.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/itserrors.Error"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update a provider",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "provider"
                ],
                "summary": "Update a provider",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Provider slug",
                        "name": "slug",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Provider",
                        "name": "provider",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Provider"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Signature",
                        "name": "X-Signature",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Provider"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/itserrors.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/itserrors.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/itserrors.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.CallRequest": {
            "type": "object",
            "properties": {
                "method": {
                    "type": "string"
                },
                "params": {
                    "type": "array",
                    "items": {}
                }
            }
        },
        "handler.Provider": {
            "type": "object"
        },
        "itserrors.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "http_status": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "model.Method": {
            "type": "object"
        },
        "model.Provider": {
            "type": "object",
            "required": [
                "name",
                "secret",
                "slug",
                "webhook"
            ],
            "properties": {
                "contact": {
                    "type": "string",
                    "example": "some@email.com"
                },
                "name": {
                    "type": "string",
                    "example": "Example LTDA"
                },
                "secret": {
                    "type": "string"
                },
                "slug": {
                    "type": "string",
                    "example": "provider-slug"
                },
                "webhook": {
                    "type": "string",
                    "example": "https://provider.com/webhook"
                }
            }
        },
        "model.ResultStructure": {
            "type": "object",
            "additionalProperties": {}
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "FED ITS API",
	Description:      "This is a conceptual API that manages providers, users and methods for the FED ITS PoC.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
