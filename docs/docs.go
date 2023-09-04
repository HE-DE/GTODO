// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/addmsga": {
            "get": {
                "tags": [
                    "消息模块"
                ],
                "summary": "添加消息(指派)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "UserId",
                        "name": "UserId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Content",
                        "name": "Content",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "AdminId",
                        "name": "AdminId",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/addmsgu": {
            "get": {
                "tags": [
                    "消息模块"
                ],
                "summary": "添加消息(未指派)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "UserId",
                        "name": "UserId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Content",
                        "name": "Content",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/deleteuser": {
            "post": {
                "tags": [
                    "用户模块"
                ],
                "summary": "删除用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "formData"
                    }
                ],
                "responses": {}
            }
        },
        "/getuser": {
            "get": {
                "tags": [
                    "用户模块"
                ],
                "summary": "获取用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/login": {
            "post": {
                "tags": [
                    "用户模块"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "formData"
                    }
                ],
                "responses": {}
            }
        },
        "/register": {
            "post": {
                "tags": [
                    "用户模块"
                ],
                "summary": "注册用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "phone",
                        "name": "phone",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "sex",
                        "name": "sex",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "indentity",
                        "name": "indentity",
                        "in": "formData"
                    }
                ],
                "responses": {}
            }
        },
        "/updateuser": {
            "post": {
                "tags": [
                    "用户模块"
                ],
                "summary": "更新用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "phone",
                        "name": "phone",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "sex",
                        "name": "sex",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "formData"
                    }
                ],
                "responses": {}
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
