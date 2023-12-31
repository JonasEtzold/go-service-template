// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Jonas Etzold",
            "email": "jonas.e@affindi.com"
        },
        "license": {
            "name": "WTFPL",
            "url": "http://www.wtfpl.net/txt/copying/"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/example": {
            "get": {
                "security": [
                    {
                        "Authorization Token": []
                    }
                ],
                "description": "Get Example based on example name or text provided as body JSON",
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves example based on query",
                "parameters": [
                    {
                        "description": "Name",
                        "name": "name",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Text",
                        "name": "text",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/example.Example"
                                }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/http_err.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http_err.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http_err.HTTPError"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Authorization Token": []
                    }
                ],
                "description": "Updates an existing example with a name and a descriptive text based on the given ID.",
                "produces": [
                    "application/json"
                ],
                "summary": "Updates an example",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Example ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Name",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Text",
                        "name": "text",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/example.Example"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/http_err.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http_err.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http_err.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Authorization Token": []
                    }
                ],
                "description": "Create an example with a name and a descriptive text.",
                "produces": [
                    "application/json"
                ],
                "summary": "Creates an example",
                "parameters": [
                    {
                        "description": "Name",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Text",
                        "name": "text",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/example.Example"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http_err.HTTPError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/http_err.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http_err.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Authorization Token": []
                    }
                ],
                "description": "Deletes an existing example based on the given ID.",
                "produces": [
                    "application/json"
                ],
                "summary": "Deletes an example",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Example ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/http_err.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http_err.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http_err.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/example/{id}": {
            "get": {
                "security": [
                    {
                        "Authorization Token": []
                    }
                ],
                "description": "get Example by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves example based on given ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Example ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/example.Example"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/http_err.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http_err.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http_err.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/login": {
            "post": {
                "security": [
                    {
                        "Authorization Token": []
                    }
                ],
                "description": "provides an authorization token for any user with valid master password. Only dummy, don't use in production.",
                "produces": [
                    "application/json"
                ],
                "summary": "Logins an example API user",
                "parameters": [
                    {
                        "description": "Username",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Password",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
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
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/http_err.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http_err.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "example.Example": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "http_err.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Go Service Template",
	Description: "A template for generating a Go backend API service",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
