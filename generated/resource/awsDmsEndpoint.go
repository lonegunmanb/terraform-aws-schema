package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDmsEndpoint = `{
  "block": {
    "attributes": {
      "certificate_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "database_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "endpoint_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "endpoint_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "engine_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "extra_connection_attributes": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kms_key_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "password": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "pause_replication_tasks": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "port": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "secrets_manager_access_role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "secrets_manager_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "server_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_access_role": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ssl_mode": {
        "computed": true,
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
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "username": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "elasticsearch_settings": {
        "block": {
          "attributes": {
            "endpoint_uri": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "error_retry_duration": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "full_load_error_percentage": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "service_access_role_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "use_new_mapping_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "kafka_settings": {
        "block": {
          "attributes": {
            "broker": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "include_control_details": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_null_and_empty": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_partition_value": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_table_alter_operations": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_transaction_details": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "message_format": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "message_max_bytes": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "no_hex_prefix": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "partition_include_schema_table": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "sasl_mechanism": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sasl_password": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "sasl_username": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "security_protocol": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ssl_ca_certificate_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ssl_client_certificate_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ssl_client_key_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ssl_client_key_password": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "topic": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "kinesis_settings": {
        "block": {
          "attributes": {
            "include_control_details": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_null_and_empty": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_partition_value": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_table_alter_operations": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_transaction_details": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "message_format": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "partition_include_schema_table": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "service_access_role_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "stream_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_large_integer_value": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "mongodb_settings": {
        "block": {
          "attributes": {
            "auth_mechanism": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "auth_source": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "auth_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "docs_to_investigate": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "extract_doc_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "nesting_level": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_update_lookup": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "mysql_settings": {
        "block": {
          "attributes": {
            "after_connect_script": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "authentication_method": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "clean_source_metadata_on_mismatch": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "events_poll_interval": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "execute_timeout": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_file_size": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "parallel_load_threads": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "server_timezone": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "service_access_role_arn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "target_db_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "oracle_settings": {
        "block": {
          "attributes": {
            "access_alternate_directly": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "add_supplemental_logging": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "additional_archived_log_dest_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "allow_selected_nested_tables": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "archived_log_dest_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "archived_logs_only": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "asm_password": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "asm_server": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "asm_user": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "authentication_method": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "char_length_semantics": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "convert_timestamp_with_zone_to_utc": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "direct_path_no_log": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "direct_path_parallel_load": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_homogenous_tablespace": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "extra_archived_log_dest_ids": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "number"
              ]
            },
            "fail_task_on_lob_truncation": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "number_datatype_scale": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "open_transaction_window": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "oracle_path_prefix": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "parallel_asm_read_threads": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "read_ahead_blocks": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "read_table_space_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "replace_path_prefix": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "retry_interval": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
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
            "security_db_encryption": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "security_db_encryption_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "spatial_data_option_to_geo_json_function_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "standby_delay_time": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "trim_space_in_char": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "use_alternate_folder_for_online": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "use_bfile": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "use_direct_path_full_load": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "use_logminer_reader": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "use_path_prefix": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "postgres_settings": {
        "block": {
          "attributes": {
            "after_connect_script": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "authentication_method": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "babelfish_database_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "capture_ddls": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "database_mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ddl_artifacts_schema": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "execute_timeout": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "fail_tasks_on_lob_truncation": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "heartbeat_enable": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "heartbeat_frequency": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "heartbeat_schema": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "map_boolean_as_boolean": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "map_jsonb_as_clob": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "map_long_varchar_as": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_file_size": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "plugin_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "service_access_role_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "slot_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "redis_settings": {
        "block": {
          "attributes": {
            "auth_password": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "auth_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "auth_user_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "port": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "server_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "ssl_ca_certificate_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ssl_security_protocol": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "redshift_settings": {
        "block": {
          "attributes": {
            "bucket_folder": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "bucket_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "encryption_mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "server_side_encryption_kms_key_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "service_access_role_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsDmsEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDmsEndpoint), &result)
	return &result
}
