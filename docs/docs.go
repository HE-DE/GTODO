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
        "/getdoing": {
            "get": {
                "tags": [
                    "消息模块"
                ],
                "summary": "获取用户的所有待办事项",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/getmsg": {
            "get": {
                "tags": [
                    "消息模块"
                ],
                "summary": "获取用户的所有待办事项",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query"
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
        "/getuserlist": {
            "get": {
                "tags": [
                    "用户模块"
                ],
                "summary": "获取用户的全部信息",
                "responses": {}
            }
        },
        "/getusername": {
            "get": {
                "tags": [
                    "用户模块"
                ],
                "summary": "获取用户名的全部信息",
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
        "/updatedoing": {
            "get": {
                "tags": [
                    "消息模块"
                ],
                "summary": "更新用户的所有待办事项的doing时间",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/updatemsg": {
            "get": {
                "tags": [
                    "消息模块"
                ],
                "summary": "更新消息的状态",
                "parameters": [
                    {
                        "type": "string",
                        "description": "InfoId",
                        "name": "InfoId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "UserId",
                        "name": "UserId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Status",
                        "name": "Status",
                        "in": "query"
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
