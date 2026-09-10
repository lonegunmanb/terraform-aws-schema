package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDmsDataProvider = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "engine": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "tags_all": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "virtual": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "settings": {
        "block": {
          "block_types": {
            "doc_db_settings": {
              "block": {
                "attributes": {
                  "certificate_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "database_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "port": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "server_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ssl_mode": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "ibm_db2_luw_settings": {
              "block": {
                "attributes": {
                  "certificate_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "database_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "encryption_algorithm": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "port": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "s3_access_role_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "security_mechanism": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "server_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ssl_mode": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "ibm_db2_zos_settings": {
              "block": {
                "attributes": {
                  "certificate_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "database_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "port": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "s3_access_role_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "server_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ssl_mode": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "maria_db_settings": {
              "block": {
                "attributes": {
                  "certificate_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "port": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "s3_access_role_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "server_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ssl_mode": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "microsoft_sql_server_settings": {
              "block": {
                "attributes": {
                  "certificate_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "database_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "port": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "s3_access_role_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "server_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ssl_mode": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "mongo_db_settings": {
              "block": {
                "attributes": {
                  "auth_mechanism": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "auth_source": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "auth_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "certificate_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "database_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "port": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "server_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ssl_mode": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "mysql_settings": {
              "block": {
                "attributes": {
                  "certificate_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "port": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "s3_access_role_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "server_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ssl_mode": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "oracle_settings": {
              "block": {
                "attributes": {
                  "asm_server": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "certificate_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "database_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "port": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "s3_access_role_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "secrets_manager_oracle_asm_access_role_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "secrets_manager_oracle_asm_secret_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "secrets_manager_security_db_encryption_access_role_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "secrets_manager_security_db_encryption_secret_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "server_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ssl_mode": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "postgresql_settings": {
              "block": {
                "attributes": {
                  "certificate_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "database_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "port": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "s3_access_role_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "server_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ssl_mode": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "redshift_settings": {
              "block": {
                "attributes": {
                  "database_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "port": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "s3_access_role_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "server_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "sybase_ase_settings": {
              "block": {
                "attributes": {
                  "certificate_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "database_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "encrypt_password": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "port": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "server_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ssl_mode": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsDmsDataProviderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDmsDataProvider), &result)
	return &result
}
