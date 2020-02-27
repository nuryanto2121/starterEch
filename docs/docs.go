// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-02-27 17:43:55.4678415 +0700 +07 m=+0.066002901

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Nuryanto",
            "url": "https://www.linkedin.com/in/nuryanto-1b2721156/",
            "email": "nuryantofattih@gmail.com"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/user": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "GetList SaUser",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "PerPage",
                        "name": "perpage",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Search",
                        "name": "search",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "InitSearch",
                        "name": "initsearch",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "SortField",
                        "name": "sortfield",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseModelList"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Add User",
                "parameters": [
                    {
                        "description": "req param #changes are possible to adjust the form of the registration form from forntend",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/contsauser.AddUserForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseModel"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "GetList SaUser",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseModel"
                        }
                    }
                }
            }
        },
        "/api/user/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "GetById SaUser",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseModel"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update User",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "req param #changes are possible to adjust the form of the registration form from forntend",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/contsauser.EditUserForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseModel"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app.ResponseModel": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "msg": {
                    "description": "Code int         ` + "`" + `json:\"code\"` + "`" + `",
                    "type": "string"
                }
            }
        },
        "contsauser.AddUserForm": {
            "type": "object",
            "properties": {
                "company_id": {
                    "type": "integer"
                },
                "created_by": {
                    "type": "string"
                },
                "email_addr": {
                    "type": "string"
                },
                "group_id": {
                    "type": "integer"
                },
                "handphone_no": {
                    "type": "string"
                },
                "level_no": {
                    "type": "integer"
                },
                "passwd": {
                    "type": "string"
                },
                "picture_url": {
                    "type": "string"
                },
                "project_id": {
                    "type": "integer"
                },
                "user_name": {
                    "type": "string"
                },
                "user_status": {
                    "type": "integer"
                }
            }
        },
        "contsauser.EditUserForm": {
            "type": "object",
            "properties": {
                "Updated_by": {
                    "type": "string"
                },
                "company_id": {
                    "type": "integer"
                },
                "email_addr": {
                    "type": "string"
                },
                "group_id": {
                    "type": "integer"
                },
                "handphone_no": {
                    "type": "string"
                },
                "level_no": {
                    "type": "integer"
                },
                "passwd": {
                    "type": "string"
                },
                "picture_url": {
                    "type": "string"
                },
                "project_id": {
                    "type": "integer"
                },
                "user_name": {
                    "type": "string"
                },
                "user_status": {
                    "type": "integer"
                }
            }
        },
        "models.ResponseModelList": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "last_page": {
                    "type": "integer"
                },
                "msg": {
                    "type": "string"
                },
                "page": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
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
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Starter",
	Description: "Backend REST API for golang starter",
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
	swag.Register(swag.Name, &s{})
}
