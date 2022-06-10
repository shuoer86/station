// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Thick client HTTP server API.",
    "title": "thick-client",
    "version": "0.0.0"
  },
  "paths": {
    "/cmd/callSmartContract": {
      "post": {
        "produces": [
          "application/json"
        ],
        "operationId": "cmdCallSC",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "operation_context": {
                  "description": "Context attributes of the operation to send to a node.",
                  "type": "object",
                  "properties": {
                    "expiry": {
                      "description": "Set the expiry duration (in ?) of the transaction.",
                      "type": "integer"
                    },
                    "fee": {
                      "description": "Set the fee amount (in massa) that will be given to the block creator.",
                      "type": "number"
                    },
                    "gaz": {
                      "description": "Gaz attibutes. Gaz is a virtual resource consumed by node while running smart contract.",
                      "type": "object",
                      "properties": {
                        "limit": {
                          "description": "Maximum number of gaz unit that a node will be able consume.",
                          "type": "integer"
                        },
                        "price": {
                          "description": "Price of a gaz unit.",
                          "type": "number"
                        }
                      }
                    },
                    "originator": {
                      "description": "Identifier of the originator of the transaction. This identifier must be knowned by the thick client.",
                      "type": "string"
                    }
                  }
                },
                "smart_contract_context": {
                  "description": "Smart contract attributes of the operation to send to a node.",
                  "type": "object",
                  "properties": {
                    "function": {
                      "description": "Function attibutes to call",
                      "type": "object",
                      "properties": {
                        "at": {
                          "description": "Address of the smart contract exposing the function",
                          "type": "string"
                        },
                        "name": {
                          "description": "Name of the function to call",
                          "type": "string"
                        },
                        "params": {
                          "description": "Parameters to pass to the function",
                          "type": "string"
                        }
                      }
                    }
                  }
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "operation id.",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/kpi": {
      "get": {
        "produces": [
          "application/json"
        ],
        "operationId": "kpi",
        "parameters": [
          {
            "uniqueItems": true,
            "type": "array",
            "items": {
              "enum": [
                "wallet",
                "node",
                "stacking",
                "blockchain"
              ],
              "type": "string"
            },
            "collectionFormat": "csv",
            "name": "scope",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "kpi message.",
            "schema": {
              "type": "object",
              "properties": {
                "node": {
                  "type": "array",
                  "items": {
                    "type": "object",
                    "properties": {
                      "cpu": {
                        "type": "array",
                        "items": {
                          "type": "number"
                        }
                      },
                      "memory": {
                        "type": "array",
                        "items": {
                          "type": "number"
                        }
                      },
                      "network": {
                        "type": "array",
                        "items": {
                          "type": "number"
                        }
                      },
                      "storage": {
                        "type": "array",
                        "items": {
                          "type": "number"
                        }
                      }
                    }
                  }
                },
                "stacking": {
                  "type": "array",
                  "items": {
                    "type": "object",
                    "properties": {
                      "address": {
                        "type": "string"
                      },
                      "gains": {
                        "type": "number"
                      },
                      "rolls": {
                        "type": "integer"
                      },
                      "slashing": {
                        "type": "integer"
                      }
                    }
                  }
                },
                "wallet": {
                  "type": "array",
                  "items": {
                    "type": "object",
                    "properties": {
                      "balance": {
                        "type": "number"
                      },
                      "coin": {
                        "type": "string"
                      }
                    }
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Thick client HTTP server API.",
    "title": "thick-client",
    "version": "0.0.0"
  },
  "paths": {
    "/cmd/callSmartContract": {
      "post": {
        "produces": [
          "application/json"
        ],
        "operationId": "cmdCallSC",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "operation_context": {
                  "description": "Context attributes of the operation to send to a node.",
                  "type": "object",
                  "properties": {
                    "expiry": {
                      "description": "Set the expiry duration (in ?) of the transaction.",
                      "type": "integer"
                    },
                    "fee": {
                      "description": "Set the fee amount (in massa) that will be given to the block creator.",
                      "type": "number"
                    },
                    "gaz": {
                      "description": "Gaz attibutes. Gaz is a virtual resource consumed by node while running smart contract.",
                      "type": "object",
                      "properties": {
                        "limit": {
                          "description": "Maximum number of gaz unit that a node will be able consume.",
                          "type": "integer"
                        },
                        "price": {
                          "description": "Price of a gaz unit.",
                          "type": "number"
                        }
                      }
                    },
                    "originator": {
                      "description": "Identifier of the originator of the transaction. This identifier must be knowned by the thick client.",
                      "type": "string"
                    }
                  }
                },
                "smart_contract_context": {
                  "description": "Smart contract attributes of the operation to send to a node.",
                  "type": "object",
                  "properties": {
                    "function": {
                      "description": "Function attibutes to call",
                      "type": "object",
                      "properties": {
                        "at": {
                          "description": "Address of the smart contract exposing the function",
                          "type": "string"
                        },
                        "name": {
                          "description": "Name of the function to call",
                          "type": "string"
                        },
                        "params": {
                          "description": "Parameters to pass to the function",
                          "type": "string"
                        }
                      }
                    }
                  }
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "operation id.",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/kpi": {
      "get": {
        "produces": [
          "application/json"
        ],
        "operationId": "kpi",
        "parameters": [
          {
            "minItems": 0,
            "uniqueItems": true,
            "type": "array",
            "items": {
              "enum": [
                "wallet",
                "node",
                "stacking",
                "blockchain"
              ],
              "type": "string"
            },
            "collectionFormat": "csv",
            "name": "scope",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "kpi message.",
            "schema": {
              "type": "object",
              "properties": {
                "node": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/NodeItems0"
                  }
                },
                "stacking": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/StackingItems0"
                  }
                },
                "wallet": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/WalletItems0"
                  }
                }
              }
            }
          }
        }
      }
    }
  },
  "definitions": {
    "CmdCallSCParamsBodyOperationContext": {
      "description": "Context attributes of the operation to send to a node.",
      "type": "object",
      "properties": {
        "expiry": {
          "description": "Set the expiry duration (in ?) of the transaction.",
          "type": "integer"
        },
        "fee": {
          "description": "Set the fee amount (in massa) that will be given to the block creator.",
          "type": "number"
        },
        "gaz": {
          "description": "Gaz attibutes. Gaz is a virtual resource consumed by node while running smart contract.",
          "type": "object",
          "properties": {
            "limit": {
              "description": "Maximum number of gaz unit that a node will be able consume.",
              "type": "integer"
            },
            "price": {
              "description": "Price of a gaz unit.",
              "type": "number"
            }
          }
        },
        "originator": {
          "description": "Identifier of the originator of the transaction. This identifier must be knowned by the thick client.",
          "type": "string"
        }
      }
    },
    "CmdCallSCParamsBodyOperationContextGaz": {
      "description": "Gaz attibutes. Gaz is a virtual resource consumed by node while running smart contract.",
      "type": "object",
      "properties": {
        "limit": {
          "description": "Maximum number of gaz unit that a node will be able consume.",
          "type": "integer"
        },
        "price": {
          "description": "Price of a gaz unit.",
          "type": "number"
        }
      }
    },
    "CmdCallSCParamsBodySmartContractContext": {
      "description": "Smart contract attributes of the operation to send to a node.",
      "type": "object",
      "properties": {
        "function": {
          "description": "Function attibutes to call",
          "type": "object",
          "properties": {
            "at": {
              "description": "Address of the smart contract exposing the function",
              "type": "string"
            },
            "name": {
              "description": "Name of the function to call",
              "type": "string"
            },
            "params": {
              "description": "Parameters to pass to the function",
              "type": "string"
            }
          }
        }
      }
    },
    "CmdCallSCParamsBodySmartContractContextFunction": {
      "description": "Function attibutes to call",
      "type": "object",
      "properties": {
        "at": {
          "description": "Address of the smart contract exposing the function",
          "type": "string"
        },
        "name": {
          "description": "Name of the function to call",
          "type": "string"
        },
        "params": {
          "description": "Parameters to pass to the function",
          "type": "string"
        }
      }
    },
    "NodeItems0": {
      "type": "object",
      "properties": {
        "cpu": {
          "type": "array",
          "items": {
            "type": "number"
          }
        },
        "memory": {
          "type": "array",
          "items": {
            "type": "number"
          }
        },
        "network": {
          "type": "array",
          "items": {
            "type": "number"
          }
        },
        "storage": {
          "type": "array",
          "items": {
            "type": "number"
          }
        }
      }
    },
    "StackingItems0": {
      "type": "object",
      "properties": {
        "address": {
          "type": "string"
        },
        "gains": {
          "type": "number"
        },
        "rolls": {
          "type": "integer"
        },
        "slashing": {
          "type": "integer"
        }
      }
    },
    "WalletItems0": {
      "type": "object",
      "properties": {
        "balance": {
          "type": "number"
        },
        "coin": {
          "type": "string"
        }
      }
    }
  }
}`))
}
