package api

func init() {
	Swagger.Add("secrets", `{
  "swagger": "2.0",
  "info": {
    "title": "api/external/secrets/secrets.proto",
    "version": "version not set"
  },
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/secrets": {
      "post": {
        "operationId": "Create",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.secrets.Id"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.secrets.Secret"
            }
          }
        ],
        "tags": [
          "SecretsService"
        ]
      }
    },
    "/secrets/id/{id}": {
      "get": {
        "operationId": "Read",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.secrets.Secret"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "SecretsService"
        ]
      },
      "delete": {
        "operationId": "Delete",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.secrets.DeleteResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "SecretsService"
        ]
      },
      "patch": {
        "operationId": "Update",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.secrets.UpdateResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.secrets.Secret"
            }
          }
        ],
        "tags": [
          "SecretsService"
        ]
      }
    },
    "/secrets/search": {
      "post": {
        "operationId": "List",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.secrets.Secrets"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.secrets.Query"
            }
          }
        ],
        "tags": [
          "SecretsService"
        ]
      }
    }
  },
  "definitions": {
    "chef.automate.api.secrets.DeleteResponse": {
      "type": "object"
    },
    "chef.automate.api.secrets.Filter": {
      "type": "object",
      "properties": {
        "key": {
          "type": "string"
        },
        "exclude": {
          "type": "boolean",
          "format": "boolean"
        },
        "values": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "chef.automate.api.secrets.Id": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.secrets.Kv": {
      "type": "object",
      "properties": {
        "key": {
          "type": "string"
        },
        "value": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.secrets.Query": {
      "type": "object",
      "properties": {
        "filters": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.secrets.Filter"
          }
        },
        "order": {
          "$ref": "#/definitions/chef.automate.api.secrets.Query.OrderType"
        },
        "sort": {
          "type": "string"
        },
        "page": {
          "type": "integer",
          "format": "int32"
        },
        "per_page": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "chef.automate.api.secrets.Query.OrderType": {
      "type": "string",
      "enum": [
        "ASC",
        "DESC"
      ],
      "default": "ASC"
    },
    "chef.automate.api.secrets.Secret": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "type": {
          "type": "string"
        },
        "last_modified": {
          "type": "string",
          "format": "date-time"
        },
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.secrets.Kv"
          }
        },
        "data": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.secrets.Kv"
          }
        }
      }
    },
    "chef.automate.api.secrets.Secrets": {
      "type": "object",
      "properties": {
        "secrets": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.secrets.Secret"
          }
        },
        "total": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "chef.automate.api.secrets.UpdateResponse": {
      "type": "object"
    }
  }
}
`)
}
