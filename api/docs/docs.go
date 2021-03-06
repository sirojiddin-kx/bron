// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/v1/client": {
            "post": {
                "description": "API for creating clients",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "client"
                ],
                "summary": "Create Company Client",
                "parameters": [
                    {
                        "description": "company-service",
                        "name": "employee",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ClientCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessModel"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            }
        },
        "/v1/clients": {
            "get": {
                "description": "API for get list of clients",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "client"
                ],
                "summary": "List Clients",
                "parameters": [
                    {
                        "type": "string",
                        "description": "company_id",
                        "name": "company_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CompanyServices"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            }
        },
        "/v1/company-service": {
            "get": {
                "description": "API for get list company_services",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "company-service"
                ],
                "summary": "List Company Services",
                "parameters": [
                    {
                        "type": "string",
                        "description": "company_id",
                        "name": "company_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CompanyServices"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            },
            "post": {
                "description": "API for create company-service",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "company-service"
                ],
                "summary": "Create Company Service",
                "parameters": [
                    {
                        "description": "company-service",
                        "name": "employee",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CompanyServiceCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessModel"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            }
        },
        "/v1/employee": {
            "get": {
                "description": "API for get list employeeslsls",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "employee"
                ],
                "summary": "List Employee",
                "parameters": [
                    {
                        "type": "string",
                        "description": "company_id",
                        "name": "company_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.EmployeeList"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            },
            "post": {
                "description": "API for creating client",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "employee"
                ],
                "summary": "Create Employee",
                "parameters": [
                    {
                        "description": "employee",
                        "name": "employee",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.EmployeeCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessModel"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            }
        },
        "/v1/employee/{employee_id}": {
            "get": {
                "description": "API for getting single Employee",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "employee"
                ],
                "summary": "Get Employee By ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "company_id",
                        "name": "company_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "guid",
                        "name": "employee_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.EmployeeList"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            }
        },
        "/v1/order": {
            "post": {
                "description": "API for creating orders",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "Create Order",
                "parameters": [
                    {
                        "description": "order",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateOrder"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessModel"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            }
        },
        "/v1/orders": {
            "get": {
                "description": "API for get list orders",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "List Orders",
                "parameters": [
                    {
                        "type": "string",
                        "description": "company_id",
                        "name": "company_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ListOrders"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ClientCreate": {
            "type": "object",
            "properties": {
                "company_id": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "models.CompanyService": {
            "type": "object",
            "properties": {
                "company_id": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer"
                },
                "guid": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.CompanyServiceCreate": {
            "type": "object",
            "properties": {
                "company_id": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer"
                },
                "price": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.CompanyServices": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "services": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.CompanyService"
                    }
                }
            }
        },
        "models.CreateOrder": {
            "type": "object",
            "properties": {
                "client_id": {
                    "type": "string"
                },
                "company_id": {
                    "type": "string"
                },
                "employee_id": {
                    "type": "string"
                },
                "order_date": {
                    "type": "string"
                },
                "services": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "start_time": {
                    "type": "string"
                }
            }
        },
        "models.Employee": {
            "type": "object",
            "properties": {
                "company_id": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "graphStartTime": {
                    "type": "string"
                },
                "graph_end_time": {
                    "type": "string"
                },
                "guid": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "position": {
                    "type": "string"
                }
            }
        },
        "models.EmployeeCreate": {
            "type": "object",
            "properties": {
                "company_id": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "graphStartTime": {
                    "type": "string"
                },
                "graph_end_time": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "position": {
                    "type": "string"
                }
            }
        },
        "models.EmployeeList": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "employees": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Employee"
                    }
                }
            }
        },
        "models.ListOrders": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "orders": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Order"
                    }
                }
            }
        },
        "models.Order": {
            "type": "object",
            "properties": {
                "client_contact": {
                    "type": "string"
                },
                "client_fullname": {
                    "type": "string"
                },
                "employee_fullname": {
                    "type": "string"
                },
                "end_time": {
                    "type": "string"
                },
                "order_date": {
                    "type": "string"
                },
                "order_id": {
                    "type": "string"
                },
                "services": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "start_time": {
                    "type": "string"
                },
                "total_price": {
                    "type": "string"
                }
            }
        },
        "models.ResponseError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.SuccessModel": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
