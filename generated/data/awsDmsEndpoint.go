package data

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
        "type": "string"
      },
      "database_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "elasticsearch_settings": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "endpoint_uri": "string",
              "error_retry_duration": "number",
              "full_load_error_percentage": "number",
              "service_access_role_arn": "string"
            }
          ]
        ]
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
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "engine_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "extra_connection_attributes": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kafka_settings": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "broker": "string",
              "include_control_details": "bool",
              "include_null_and_empty": "bool",
              "include_partition_value": "bool",
              "include_table_alter_operations": "bool",
              "include_transaction_details": "bool",
              "message_format": "string",
              "message_max_bytes": "number",
              "no_hex_prefix": "bool",
              "partition_include_schema_table": "bool",
              "sasl_mechanism": "string",
              "sasl_password": "string",
              "sasl_username": "string",
              "security_protocol": "string",
              "ssl_ca_certificate_arn": "string",
              "ssl_client_certificate_arn": "string",
              "ssl_client_key_arn": "string",
              "ssl_client_key_password": "string",
              "topic": "string"
            }
          ]
        ]
      },
      "kinesis_settings": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "include_control_details": "bool",
              "include_null_and_empty": "bool",
              "include_partition_value": "bool",
              "include_table_alter_operations": "bool",
              "include_transaction_details": "bool",
              "message_format": "string",
              "partition_include_schema_table": "bool",
              "service_access_role_arn": "string",
              "stream_arn": "string",
              "use_large_integer_value": "bool"
            }
          ]
        ]
      },
      "kms_key_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "mongodb_settings": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "auth_mechanism": "string",
              "auth_source": "string",
              "auth_type": "string",
              "docs_to_investigate": "string",
              "extract_doc_id": "string",
              "nesting_level": "string"
            }
          ]
        ]
      },
      "password": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "port": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "postgres_settings": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "after_connect_script": "string",
              "authentication_method": "string",
              "babelfish_database_name": "string",
              "capture_ddls": "bool",
              "database_mode": "string",
              "ddl_artifacts_schema": "string",
              "execute_timeout": "number",
              "fail_tasks_on_lob_truncation": "bool",
              "heartbeat_enable": "bool",
              "heartbeat_frequency": "number",
              "heartbeat_schema": "string",
              "map_boolean_as_boolean": "bool",
              "map_jsonb_as_clob": "bool",
              "map_long_varchar_as": "string",
              "max_file_size": "number",
              "plugin_name": "string",
              "service_access_role_arn": "string",
              "slot_name": "string"
            }
          ]
        ]
      },
      "redis_settings": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "auth_password": "string",
              "auth_type": "string",
              "auth_user_name": "string",
              "port": "number",
              "server_name": "string",
              "ssl_ca_certificate_arn": "string",
              "ssl_security_protocol": "string"
            }
          ]
        ]
      },
      "redshift_settings": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "bucket_folder": "string",
              "bucket_name": "string",
              "encryption_mode": "string",
              "server_side_encryption_kms_key_id": "string",
              "service_access_role_arn": "string"
            }
          ]
        ]
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "s3_settings": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "add_column_name": "bool",
              "bucket_folder": "string",
              "bucket_name": "string",
              "canned_acl_for_objects": "string",
              "cdc_inserts_and_updates": "bool",
              "cdc_inserts_only": "bool",
              "cdc_max_batch_interval": "number",
              "cdc_min_file_size": "number",
              "cdc_path": "string",
              "compression_type": "string",
              "csv_delimiter": "string",
              "csv_no_sup_value": "string",
              "csv_null_value": "string",
              "csv_row_delimiter": "string",
              "data_format": "string",
              "data_page_size": "number",
              "date_partition_delimiter": "string",
              "date_partition_enabled": "bool",
              "date_partition_sequence": "string",
              "dict_page_size_limit": "number",
              "enable_statistics": "bool",
              "encoding_type": "string",
              "encryption_mode": "string",
              "external_table_definition": "string",
              "glue_catalog_generation": "bool",
              "ignore_header_rows": "number",
              "ignore_headers_row": "number",
              "include_op_for_full_load": "bool",
              "max_file_size": "number",
              "parquet_timestamp_in_millisecond": "bool",
              "parquet_version": "string",
              "preserve_transactions": "bool",
              "rfc_4180": "bool",
              "row_group_length": "number",
              "server_side_encryption_kms_key_id": "string",
              "service_access_role_arn": "string",
              "timestamp_column_name": "string",
              "use_csv_no_sup_value": "bool",
              "use_task_start_time_for_full_load_timestamp": "bool"
            }
          ]
        ]
      },
      "secrets_manager_access_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secrets_manager_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "server_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_access_role": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ssl_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "username": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
