// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-06-14 05:48:39.93579 +0700 WIB m=+0.156313891

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
        "/api/auth/forgot": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Forgot Password",
                "parameters": [
                    {
                        "description": "req param #changes are possible to adjust the form of the registration form from frontend",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.ForgotForm"
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
        },
        "/api/auth/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "req param #changes are possible to adjust the form of the registration form from frontend",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.LoginForm"
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
        },
        "/api/auth/register": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Add Client",
                "parameters": [
                    {
                        "description": "req param #changes are possible to adjust the form of the registration form from frontend",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.RegisterForm"
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
        },
        "/api/auth/reset": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Reset Password",
                "parameters": [
                    {
                        "description": "req param #changes are possible to adjust the form of the registration form from frontend",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.ResetPasswd"
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
        },
        "/api/auth/verify": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Verify / Aktivasi User",
                "parameters": [
                    {
                        "description": "req param #changes are possible to adjust the form of the registration form from frontend",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.ResetPasswd"
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
        },
        "/api/fileupload": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Upload file",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "FileUpload"
                ],
                "summary": "File Upload",
                "parameters": [
                    {
                        "type": "file",
                        "description": "account image",
                        "name": "upload_file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "path images",
                        "name": "path",
                        "in": "formData",
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
        "/api/menu": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Menu"
                ],
                "summary": "GetList SaMenu",
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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Menu"
                ],
                "summary": "Add Menu",
                "parameters": [
                    {
                        "description": "req param #changes are possible to adjust the form of the registration form from frontend",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.AddMenuForm"
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
        },
        "/api/menu/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Menu"
                ],
                "summary": "GetById SaMenu",
                "parameters": [
                    {
                        "type": "string",
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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Menu"
                ],
                "summary": "Update Menu",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "req param #changes are possible to adjust the form of the registration form from frontend",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.EditMenuForm"
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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Menu"
                ],
                "summary": "Delete menu",
                "parameters": [
                    {
                        "type": "string",
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
        "/api/role": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Role"
                ],
                "summary": "GetList SaRole",
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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Role"
                ],
                "summary": "Add Role",
                "parameters": [
                    {
                        "description": "req param #changes are possible to adjust the form of the registration form from frontend",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.AddRoleForm"
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
        },
        "/api/role/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Role"
                ],
                "summary": "GetById SaRole",
                "parameters": [
                    {
                        "type": "string",
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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Role"
                ],
                "summary": "Update Role",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "req param #changes are possible to adjust the form of the registration form from frontend",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.EditRoleForm"
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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Role"
                ],
                "summary": "Delete role",
                "parameters": [
                    {
                        "type": "string",
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
        "/api/user": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Add User",
                "parameters": [
                    {
                        "description": "req param #changes are possible to adjust the form of the registration form from frontend",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.AddUserForm"
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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "GetList SaUser",
                "parameters": [
                    {
                        "type": "string",
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
        "/api/user/permission": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get string Array Permission User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ClientID",
                        "name": "client_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "UserID",
                        "name": "user_id",
                        "in": "query"
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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "GetById SaUser",
                "parameters": [
                    {
                        "type": "string",
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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "req param #changes are possible to adjust the form of the registration form from frontend",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.EditUserForm"
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
                "message": {
                    "description": "Code int         ` + "`" + `json:\"code\"` + "`" + `",
                    "type": "string"
                }
            }
        },
        "models.AddMenuForm": {
            "type": "object",
            "properties": {
                "icon_class": {
                    "type": "string"
                },
                "level": {
                    "type": "integer"
                },
                "menu_url": {
                    "type": "string"
                },
                "order_seq": {
                    "type": "integer"
                },
                "parent_menu_id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.AddRoleForm": {
            "type": "object",
            "properties": {
                "created_by": {
                    "type": "string"
                },
                "descs": {
                    "type": "string"
                },
                "num": {
                    "type": "number"
                }
            }
        },
        "models.AddUserForm": {
            "type": "object",
            "properties": {
                "client_id": {
                    "type": "string"
                },
                "company_id": {
                    "type": "integer"
                },
                "created_by": {
                    "type": "string"
                },
                "data_permission": {
                    "type": "string"
                },
                "email_addr": {
                    "type": "string"
                },
                "file_id": {
                    "type": "string"
                },
                "handphone_no": {
                    "type": "string"
                },
                "level_no": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "role_id": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                },
                "user_status": {
                    "type": "integer"
                }
            }
        },
        "models.EditMenuForm": {
            "type": "object",
            "properties": {
                "icon_class": {
                    "type": "string"
                },
                "level": {
                    "type": "integer"
                },
                "menu_url": {
                    "type": "string"
                },
                "order_seq": {
                    "type": "integer"
                },
                "parent_menu_id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.EditRoleForm": {
            "type": "object",
            "properties": {
                "Updated_by": {
                    "type": "string"
                },
                "descs": {
                    "type": "string"
                }
            }
        },
        "models.EditUserForm": {
            "type": "object",
            "properties": {
                "client_id": {
                    "type": "string"
                },
                "company_id": {
                    "type": "integer"
                },
                "data_permission": {
                    "type": "string"
                },
                "email_addr": {
                    "type": "string"
                },
                "file_id": {
                    "type": "string"
                },
                "handphone_no": {
                    "type": "string"
                },
                "level_no": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "role_id": {
                    "type": "string"
                },
                "updated_by": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                },
                "user_status": {
                    "type": "integer"
                }
            }
        },
        "models.ForgotForm": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "models.LoginForm": {
            "type": "object",
            "properties": {
                "p": {
                    "type": "string"
                },
                "u": {
                    "type": "string"
                }
            }
        },
        "models.RegisterForm": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "client_name": {
                    "type": "string"
                },
                "client_type": {
                    "type": "string"
                },
                "contact_person": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "email_addr": {
                    "type": "string"
                },
                "expiry_date": {
                    "type": "string"
                },
                "joining_date": {
                    "type": "string"
                },
                "post_cd": {
                    "type": "string"
                },
                "start_billing_date": {
                    "type": "string"
                },
                "telephone_no": {
                    "type": "string"
                }
            }
        },
        "models.ResetPasswd": {
            "type": "object",
            "properties": {
                "cp": {
                    "type": "string"
                },
                "p": {
                    "type": "string"
                },
                "token_email": {
                    "type": "string"
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
