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
            "name": "John Avitable",
            "url": "https://github.com/javitab"
        },
        "license": {
            "name": "GPL-3.0",
            "url": "https://github.com/javitab/go-web/blob/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/server_events": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Gets all server events from database matching filter criteria",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api"
                ],
                "summary": "Get Logged Server Events",
                "parameters": [
                    {
                        "type": "string",
                        "description": "search filter",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "EventType to filter for",
                        "name": "EventType",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "ServerRunID to filter for",
                        "name": "ServerRunID",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.GetServerEventsResponse"
                        }
                    }
                }
            }
        },
        "/auth/create_user": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a new user given a CreateUserInput object",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user/group security"
                ],
                "summary": "Create a new user via API",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.CreateUserInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.CreateUserResponse"
                        }
                    }
                }
            }
        },
        "/auth/generate_api_key": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Given a username, gives details about user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user/group security"
                ],
                "summary": "Generate API Key",
                "parameters": [
                    {
                        "type": "string",
                        "description": "desc for key usage",
                        "name": "description",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.GenerateAPIKeyResponse"
                        }
                    }
                }
            }
        },
        "/auth/generate_jwt": {
            "post": {
                "description": "Authenticate a user given an APIKey. Returns a JWT token upon successful authentication",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "login"
                ],
                "summary": "Get JWT from API Key",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.GetJWTFromAPIKeyInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/group": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Given a group id, gives details about group",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user/group security"
                ],
                "summary": "Get group info",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "groupid to lookup",
                        "name": "group_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.UserInfo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Authenticate a user given a LoginUserInput object. Returns a JWT token upon successful authentication",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "login"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.LoginUserInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/sec_point": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Returns information about a single security point including groups with FK relationships",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user/group security"
                ],
                "summary": "Get information about a given security point",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "SPID to search",
                        "name": "spid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.SecPointInfo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/update_user": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Given a username, will make given updates",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "user/group security"
                ],
                "summary": "Update User Record",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username to update",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "enum": [
                            "delete_user",
                            "undelete_user",
                            "add_group",
                            "remove_group",
                            "add_user_sec_point",
                            "remove_user_sec_point"
                        ],
                        "type": "string",
                        "description": "action to perform",
                        "name": "action",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "reason for update (incident #, etc.)",
                        "name": "reason",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "value to set",
                        "name": "value",
                        "in": "query",
                        "required": true
                    },
                    {
                        "enum": [
                            "UserAddSecPoints",
                            "UserDelSecPoints",
                            "UserOvrSecPoints"
                        ],
                        "type": "string",
                        "description": "field to append user-level security point to",
                        "name": "sec_point_field",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/user": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Given a username, gives details about user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user/group security"
                ],
                "summary": "Get user info",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username to lookup",
                        "name": "username",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.UserInfo"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.GetServerEventsResponse": {
            "type": "object",
            "properties": {
                "events": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/database.ServerEvent"
                    }
                },
                "limit": {
                    "type": "integer"
                }
            }
        },
        "auth.CreateUserInput": {
            "type": "object",
            "required": [
                "email",
                "first_name",
                "last_name",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "auth.CreateUserResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "enum": [
                        "User created",
                        "User already exists",
                        "Error creating user"
                    ]
                }
            }
        },
        "auth.EvalSP": {
            "type": "object",
            "properties": {
                "source": {
                    "type": "string"
                },
                "sp": {
                    "$ref": "#/definitions/database.SecPoint"
                }
            }
        },
        "auth.GenerateAPIKeyResponse": {
            "type": "object",
            "properties": {
                "apikey": {
                    "type": "string"
                },
                "message": {
                    "type": "string",
                    "enum": [
                        "API Key generated"
                    ]
                },
                "user": {
                    "type": "string"
                }
            }
        },
        "auth.GetJWTFromAPIKeyInput": {
            "type": "object",
            "required": [
                "key"
            ],
            "properties": {
                "key": {
                    "type": "string"
                }
            }
        },
        "auth.GroupInfo": {
            "type": "object",
            "properties": {
                "db": {
                    "$ref": "#/definitions/database.Group"
                }
            }
        },
        "auth.LoginUserInput": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "auth.SecPointInfo": {
            "type": "object",
            "properties": {
                "db": {
                    "$ref": "#/definitions/database.SecPoint"
                },
                "referencingGroups": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/auth.GroupInfo"
                    }
                }
            }
        },
        "auth.UserInfo": {
            "type": "object",
            "properties": {
                "db": {
                    "$ref": "#/definitions/database.User"
                },
                "isActiveUser": {
                    "type": "boolean"
                },
                "isLDAPUser": {
                    "type": "boolean"
                },
                "ldapgroups": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "securityPoints": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/auth.EvalSP"
                    }
                }
            }
        },
        "database.Group": {
            "type": "object",
            "properties": {
                "addSecPoints": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/database.SecPoint"
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "delSecPoints": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/database.SecPoint"
                    }
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "desc": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "ldap_group": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "ovrSecPoints": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/database.SecPoint"
                    }
                },
                "priority": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "database.SecPoint": {
            "type": "object",
            "properties": {
                "ID": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "desc": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "sp_group": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "database.ServerEvent": {
            "type": "object",
            "properties": {
                "archived": {
                    "type": "boolean"
                },
                "createdAt": {
                    "type": "string"
                },
                "dateTime": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "details": {
                    "type": "string"
                },
                "eventType": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "serverRunID": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "uuid_ID": {
                    "type": "string"
                }
            }
        },
        "database.User": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "groups": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/database.Group"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "isLDAPUser": {
                    "type": "boolean",
                    "default": false
                },
                "lastName": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userAddSecPoints": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/database.SecPoint"
                    }
                },
                "userDelSecPoints": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/database.SecPoint"
                    }
                },
                "userOvrSecPoints": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/database.SecPoint"
                    }
                },
                "username": {
                    "type": "string"
                },
                "uuid_ID": {
                    "type": "string"
                }
            }
        },
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "description": "JWT can be obtained from ` + "`" + `login` + "`" + ` or ` + "`" + `generate_jwt` + "`" + ` endpoints. Be sure to include ` + "`" + `Bearer` + "`" + ` before the JWT",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Go Web API Documentation",
	Description:      "This is a sample Go-based CRUD Application Template",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
