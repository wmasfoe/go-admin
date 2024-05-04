// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/user": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "获取所有用户列表",
                "responses": {
                    "200": {
                        "description": "code\", \"msg\", \"data\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "新建用户接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "Name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "Password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\", \"msg\", \"data\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/user/:id": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户详情",
                "responses": {
                    "200": {
                        "description": "code\", \"msg\", \"data\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "删除用户",
                "responses": {
                    "200": {
                        "description": "code\", \"msg\", \"data\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
