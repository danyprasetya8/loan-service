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
        "/api/v1/auth/mock-login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Mock login using email and user role",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.MockLogin"
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
        "/api/v1/borrower": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Borrower"
                ],
                "summary": "Get list of borrowers",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "default is 1",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "default is 10",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.GetBorrower"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Borrower can only be created by user with fieldOfficer role",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Borrower"
                ],
                "summary": "Create borrower",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateBorrower"
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
        "/api/v1/borrower/{id}": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Borrower"
                ],
                "summary": "Delete borrower by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Borrower ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "boolean"
                        }
                    }
                }
            }
        },
        "/api/v1/loan": {
            "post": {
                "description": "User with role fieldOfficer can propose a loan for a borrower",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Loan"
                ],
                "summary": "Propose loan by field officer",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ProposeLoan"
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
        "/api/v1/loan/{id}/_approve": {
            "post": {
                "description": "User with role internal can approve a loan. fieldOfficerId is user's email, proofOfPicture is file id got from upload response",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Loan"
                ],
                "summary": "Approve loan by internal",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Loan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ApproveLoan"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "boolean"
                        }
                    }
                }
            }
        },
        "/api/v1/loan/{id}/_disburse": {
            "post": {
                "description": "User with role internal can approve a loan. fieldOfficerId is user's email, borrowerAgreementLetter is file id got from upload response",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Loan"
                ],
                "summary": "Invest loan by internal",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Loan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.DisburseLoan"
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
        "/api/v1/loan/{id}/_invest": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Loan"
                ],
                "summary": "Invest loan by investor",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Loan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.InvestLoan"
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
        "/api/v1/loan/{id}/borrower-agreement-letter": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Loan"
                ],
                "summary": "Upload borrower agreement letter",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Loan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "PDF to upload",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.UploadBorrowerLetter"
                        }
                    }
                }
            }
        },
        "/api/v1/loan/{id}/proof": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Loan"
                ],
                "summary": "Upload loan proof of picture",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Loan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Picture to upload",
                        "name": "image",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.UploadLoanProofOfPicture"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.ApproveLoan": {
            "type": "object",
            "properties": {
                "fieldOfficerId": {
                    "type": "string"
                },
                "proofOfPicture": {
                    "type": "string"
                }
            }
        },
        "request.CreateBorrower": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "request.DisburseLoan": {
            "type": "object",
            "properties": {
                "borrowerAgreementLetter": {
                    "type": "string"
                },
                "fieldOfficerId": {
                    "type": "string"
                }
            }
        },
        "request.InvestLoan": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                }
            }
        },
        "request.MockLogin": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "request.ProposeLoan": {
            "type": "object",
            "properties": {
                "borrowerId": {
                    "type": "string"
                },
                "principalAmount": {
                    "type": "integer"
                },
                "rate": {
                    "type": "number"
                },
                "roi": {
                    "type": "number"
                }
            }
        },
        "response.GetBorrower": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "createdBy": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "response.UploadBorrowerLetter": {
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
                }
            }
        },
        "response.UploadLoanProofOfPicture": {
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
