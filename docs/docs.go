// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "swagger": "2.0",
    "info": {
      "version": "1.0.0",
      "title": "Swagger Petstore",
      "description": "A sample API that uses a petstore as an example to demonstrate features in the swagger-2.0 specification",
      "termsOfService": "http://swagger.io/terms/",
      "contact": {
        "name": "Swagger API Team"
      },
      "license": {
        "name": "MIT"
      }
    },
    "host": "localhost:5000",
    "basePath": "/apis",
    "schemes": [
      "http"
    ],
    "consumes": [
      "application/json"
    ],
    "produces": [
      "application/json"
    ],
    "paths": {
      "/pets": {
        "get": {
          "description": "Returns all pets from the system that the user has access to",
          "operationId": "findPets",
          "produces": [
            "application/json",
            "application/xml",
            "text/xml",
            "text/html"
          ],
          "parameters": [
            {
              "name": "tags",
              "in": "query",
              "description": "tags to filter by",
              "required": false,
              "type": "array",
              "items": {
                "type": "string"
              },
              "collectionFormat": "csv"
            },
            {
              "name": "limit",
              "in": "query",
              "description": "maximum number of results to return",
              "required": false,
              "type": "integer",
              "format": "int32"
            }
          ],
          "responses": {
            "200": {
              "description": "pet response",
              "schema": {
                "type": "array",
                "items": {
                  "$ref": "#/definitions/Pet"
                }
              }
            },
            "default": {
              "description": "unexpected error",
              "schema": {
                "$ref": "#/definitions/ErrorModel"
              }
            }
          }
        },
        "post": {
          "description": "Creates a new pet in the store.  Duplicates are allowed",
          "operationId": "addPet",
          "produces": [
            "application/json"
          ],
          "parameters": [
            {
              "name": "pet",
              "in": "body",
              "description": "Pet to add to the store",
              "required": true,
              "schema": {
                "$ref": "#/definitions/NewPet"
              }
            }
          ],
          "responses": {
            "200": {
              "description": "pet response",
              "schema": {
                "$ref": "#/definitions/Pet"
              }
            },
            "default": {
              "description": "unexpected error",
              "schema": {
                "$ref": "#/definitions/ErrorModel"
              }
            }
          }
        }
      },
      "/pets/{id}": {
        "get": {
          "description": "Returns a user based on a single ID, if the user does not have access to the pet",
          "operationId": "findPetById",
          "produces": [
            "application/json",
            "application/xml",
            "text/xml",
            "text/html"
          ],
          "parameters": [
            {
              "name": "id",
              "in": "path",
              "description": "ID of pet to fetch",
              "required": true,
              "type": "integer",
              "format": "int64"
            }
          ],
          "responses": {
            "200": {
              "description": "pet response",
              "schema": {
                "$ref": "#/definitions/Pet"
              }
            },
            "default": {
              "description": "unexpected error",
              "schema": {
                "$ref": "#/definitions/ErrorModel"
              }
            }
          }
        },
        "delete": {
          "description": "deletes a single pet based on the ID supplied",
          "operationId": "deletePet",
          "parameters": [
            {
              "name": "id",
              "in": "path",
              "description": "ID of pet to delete",
              "required": true,
              "type": "integer",
              "format": "int64"
            }
          ],
          "responses": {
            "204": {
              "description": "pet deleted"
            },
            "default": {
              "description": "unexpected error",
              "schema": {
                "$ref": "#/definitions/ErrorModel"
              }
            }
          }
        }
      },
      "/v1/user/login": {
        "post": {
          "description": "User login",
          "operationId": "userLogin",
          "produces": [
            "application/json",
          ],
          "parameters": [
            {
                "name": "pet",
                "in": "body",
                "description": "User login request",
                "required": true,
                "schema": {
                    "$ref": "#/definitions/NewUserLogin"
                }
            }
          ],
          "responses": {
            "200": {
              "description": "User login response",
              "schema": {
                "$ref": "#/definitions/Pet"
              }
            },
            "default": {
              "description": "unexpected error",
              "schema": {
                "$ref": "#/definitions/ErrorModel"
              }
            }
          }
        }
      }
    },
    "definitions": {
      "Pet": {
        "type": "object",
        "allOf": [
          {
            "$ref": "#/definitions/NewPet"
          },
          {
            "required": [
              "id"
            ],
            "properties": {
              "id": {
                "type": "integer",
                "format": "int64"
              }
            }
          }
        ]
      },
      "NewPet": {
        "type": "object",
        "required": [
          "name"
        ],
        "properties": {
          "name": {
            "type": "string"
          },
          "tag": {
            "type": "string"
          }
        }
      },
      "UserLogin": {
        "type": "object",
        "properties": {
          "code": {
            "type": "integer",
            "example": 200
          },
          "message": {
            "type": "string",
            "example": "Success"
          },
          "data": {
            "$ref": "#/definitions/UserLoginModel"
          }
        }
      },
      "UserLoginModel": {
        "type": "object",
        "properties": {
          "id": {
            "type": "integer",
            "example": 1
          },
          "username": {
            "type": "string",
            "example": "root"
          },
          "token": {
            "type": "string",
            "example": "token"
          }
        }
      },
      "NewUserLogin": {
        "type": "object",
        "required": [
          "username",
          "password"
        ],
        "properties": {
          "username": {
            "type": "string"
          },
          "password": {
            "type": "string"
          }
        }
      },
      "ErrorModel": {
        "type": "object",
        "required": [
          "code",
          "message"
        ],
        "properties": {
          "code": {
            "type": "integer",
            "format": "int32"
          },
          "message": {
            "type": "string"
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
