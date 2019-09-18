// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-09-18 13:34:09.321106544 +0100 BST m=+0.021854677

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is a temporarily centralized directory/PKI/metrics API to allow us to get the other Nym node types running. Its functionality will eventually be folded into other parts of Nym.",
        "title": "Nym Directory API",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "https://github.com/nymtech/directory-server/license"
        },
        "version": "1.0"
    },
    "paths": {
        "/api/healthcheck": {
            "get": {
                "description": "Returns a 200 if the directory server is available.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "healthcheck"
                ],
                "summary": "Lets the directory server tell the world it's alive",
                "operationId": "healthCheck",
                "responses": {
                    "200": {}
                }
            }
        },
        "/api/metrics/mixes": {
            "get": {
                "description": "You'd never want to run this in production, but for demo and debug purposes it gives us the ability to generate useful visualisations of network traffic.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "metrics"
                ],
                "summary": "Lists mixnode activity in the past 1 second",
                "operationId": "listMixMetrics",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.MixMetric"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "You'd never want to run this in production, but for demo and debug purposes it gives us the ability to generate useful visualisations of network traffic.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "metrics"
                ],
                "summary": "Create a metric detailing how many messages a given mixnode sent and received",
                "operationId": "createMixMetric",
                "parameters": [
                    {
                        "description": "object",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.MixMetric"
                        }
                    }
                ],
                "responses": {
                    "201": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/api/presence/coconodes": {
            "post": {
                "description": "Nym Coconut nodes can ping this method to let the directory server know they're up. We can then use this info to create topologies of the overall Nym network.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "presence"
                ],
                "summary": "Lets a coconut node tell the directory server it's alive",
                "operationId": "addCocoNode",
                "parameters": [
                    {
                        "description": "object",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.HostInfo"
                        }
                    }
                ],
                "responses": {
                    "201": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/api/presence/mixnodes": {
            "post": {
                "description": "Nym mixnodes can ping this method to let the directory server know they're up. We can then use this info to create topologies of the overall Nym network.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "presence"
                ],
                "summary": "Lets mixnode a node tell the directory server it's alive",
                "operationId": "addMixNode",
                "parameters": [
                    {
                        "description": "object",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.MixHostInfo"
                        }
                    }
                ],
                "responses": {
                    "201": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/api/presence/providers": {
            "post": {
                "description": "Nym mix providers can ping this method to let the directory server know they're up. We can then use this info to create topologies of the overall Nym network.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "presence"
                ],
                "summary": "Lets a mixnode tell the directory server it's alive",
                "operationId": "addMixProvider",
                "parameters": [
                    {
                        "description": "object",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.MixProviderHostInfo"
                        }
                    }
                ],
                "responses": {
                    "201": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/api/presence/topology": {
            "get": {
                "description": "Nym nodes periodically ping the directory server to register that they're alive. This method provides a list of nodes which have been most recently seen.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "presence"
                ],
                "summary": "Lists which Nym mixnodes, providers, and coconodes are alive",
                "operationId": "topology",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Topology"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Error": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "models.HostInfo": {
            "type": "object",
            "required": [
                "host",
                "pubKey"
            ],
            "properties": {
                "host": {
                    "type": "string"
                },
                "pubKey": {
                    "type": "string"
                }
            }
        },
        "models.MixHostInfo": {
            "type": "object",
            "required": [
                "pubKey",
                "layer",
                "host"
            ],
            "properties": {
                "host": {
                    "type": "string"
                },
                "layer": {
                    "type": "integer"
                },
                "pubKey": {
                    "type": "string"
                }
            }
        },
        "models.MixMetric": {
            "type": "object",
            "required": [
                "pubKey"
            ],
            "properties": {
                "pubKey": {
                    "type": "string"
                },
                "received": {
                    "type": "integer"
                },
                "sent": {
                    "type": "object",
                    "required": [
                        "sent"
                    ]
                }
            }
        },
        "models.MixProviderHostInfo": {
            "type": "object",
            "required": [
                "host",
                "pubKey"
            ],
            "properties": {
                "host": {
                    "type": "string"
                },
                "pubKey": {
                    "type": "string"
                },
                "registeredClients": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.RegisteredClient"
                    }
                }
            }
        },
        "models.RegisteredClient": {
            "type": "object",
            "required": [
                "pubKey"
            ],
            "properties": {
                "pubKey": {
                    "type": "string"
                }
            }
        },
        "models.Topology": {
            "type": "object",
            "properties": {
                "cocoNodes": {
                    "type": "object"
                },
                "mixNodes": {
                    "type": "object"
                },
                "mixProviderNodes": {
                    "type": "object"
                }
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}
