// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms",
        "contact": {
            "name": "Gadam Jumayev",
            "url": "https://github.com/jumayevgadam",
            "email": "hypergadam@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/author/register": {
            "post": {
                "description": "create author with properties.",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AUTHORS"
                ],
                "summary": "CREATE-AUTHOR.",
                "operationId": "create-author",
                "parameters": [
                    {
                        "type": "string",
                        "name": "avatar",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "biography",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "maxLength": 20,
                        "minLength": 6,
                        "type": "string",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "int"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "localhost:4000",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "BOOK-APP-WITH-REFRESH-TOKEN api documentation",
	Description:      "book app with refresh token.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
