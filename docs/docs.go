// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/api/v1/userExample": {
            "post": {
                "description": "提交信息创建userExample",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "userExample"
                ],
                "summary": "创建userExample",
                "parameters": [
                    {
                        "description": "userExample信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.CreateUserExampleRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/userExample/{id}": {
            "get": {
                "description": "根据id获取userExample详情",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "userExample"
                ],
                "summary": "获取userExample详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Result"
                        }
                    }
                }
            },
            "put": {
                "description": "根据id更新userExample信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "userExample"
                ],
                "summary": "更新userExample信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "userExample信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.UpdateUserExampleByIDRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Result"
                        }
                    }
                }
            },
            "delete": {
                "description": "根据id删除userExample",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "userExample"
                ],
                "summary": "删除userExample",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/userExamples": {
            "post": {
                "description": "使用post请求获取userExample列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "userExample"
                ],
                "summary": "获取userExample列表",
                "parameters": [
                    {
                        "description": "查询条件",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.Params"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Result"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "check health",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "system"
                ],
                "summary": "check health",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlerfunc.checkHealthResponse"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "system"
                ],
                "summary": "ping",
                "responses": {}
            }
        }
    },
    "definitions": {
        "handler.Column": {
            "type": "object",
            "properties": {
                "exp": {
                    "description": "表达式，值为空时默认为=，有=、!=、\u003e、\u003e=、\u003c、\u003c=、like七种类型",
                    "type": "string"
                },
                "logic": {
                    "description": "逻辑类型，值为空时默认为and，有\u0026(and)、||(or)两种类型",
                    "type": "string"
                },
                "name": {
                    "description": "列名",
                    "type": "string"
                },
                "value": {
                    "description": "列值"
                }
            }
        },
        "handler.CreateUserExampleRequest": {
            "type": "object",
            "properties": {
                "age": {
                    "description": "年龄",
                    "type": "integer"
                },
                "avatar": {
                    "description": "头像",
                    "type": "string",
                    "minLength": 5
                },
                "email": {
                    "description": "邮件",
                    "type": "string"
                },
                "gender": {
                    "description": "性别，1:男，2:女",
                    "type": "integer",
                    "maximum": 2,
                    "minimum": 0
                },
                "name": {
                    "description": "名称",
                    "type": "string",
                    "minLength": 2
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "phone": {
                    "description": "手机号码，必须在前加'+86'",
                    "type": "string"
                }
            }
        },
        "handler.Params": {
            "type": "object",
            "properties": {
                "columns": {
                    "description": "列查询条件",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handler.Column"
                    }
                },
                "page": {
                    "description": "页码",
                    "type": "integer",
                    "minimum": 0
                },
                "size": {
                    "description": "每页行数",
                    "type": "integer"
                },
                "sort": {
                    "description": "排序字段，默认值为-id，字段前面有-号表示倒序，否则升序，多个字段用逗号分隔",
                    "type": "string"
                }
            }
        },
        "handler.Result": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "返回码",
                    "type": "integer"
                },
                "data": {
                    "description": "返回数据"
                },
                "msg": {
                    "description": "返回信息说明",
                    "type": "string"
                }
            }
        },
        "handler.UpdateUserExampleByIDRequest": {
            "type": "object",
            "properties": {
                "age": {
                    "description": "年龄",
                    "type": "integer"
                },
                "avatar": {
                    "description": "头像",
                    "type": "string"
                },
                "email": {
                    "description": "邮件",
                    "type": "string"
                },
                "gender": {
                    "description": "性别，1:男，2:女",
                    "type": "integer"
                },
                "id": {
                    "description": "id",
                    "type": "integer"
                },
                "name": {
                    "description": "名称",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "phone": {
                    "description": "手机号码，必须在前加'+86'",
                    "type": "string"
                }
            }
        },
        "handlerfunc.checkHealthResponse": {
            "type": "object",
            "properties": {
                "hostname": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v0.0.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{"http", "https"},
	Title:            "sponge api docs",
	Description:      "http server api docs",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
