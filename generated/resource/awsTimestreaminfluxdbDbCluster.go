package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsTimestreaminfluxdbDbCluster = `{
  "block": {
    "attributes": {
      "allocated_storage": {
        "description": "The amount of storage to allocate for your DB storage type in GiB (gibibytes).",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bucket": {
        "description": "The name of the initial InfluxDB bucket. All InfluxDB data is stored in a bucket. \n\t\t\t\t\tA bucket combines the concept of a database and a retention period (the duration of time \n\t\t\t\t\tthat each data point persists). A bucket belongs to an organization.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_instance_type": {
        "description": "The Timestream for InfluxDB DB instance type to run InfluxDB on.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_parameter_group_identifier": {
        "description": "The ID of the DB parameter group to assign to your DB cluster. \n\t\t\t\t\tDB parameter groups specify how the database is configured. For example, DB parameter groups \n\t\t\t\t\tcan specify the limit for query concurrency.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_storage_type": {
        "computed": true,
        "description": "The Timestream for InfluxDB DB storage type to read and write InfluxDB data. \n\t\t\t\t\tYou can choose between 3 different types of provisioned Influx IOPS included storage according \n\t\t\t\t\tto your workloads requirements: Influx IO Included 3000 IOPS, Influx IO Included 12000 IOPS, \n\t\t\t\t\tInflux IO Included 16000 IOPS.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deployment_type": {
        "computed": true,
        "description": "Specifies the type of cluster to create.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "endpoint": {
        "computed": true,
        "description": "The endpoint used to connect to InfluxDB. The default InfluxDB port is 8086.",
        "description_kind": "plain",
        "type": "string"
      },
      "failover_mode": {
        "computed": true,
        "description": "Specifies the behavior of failure recovery when the primary node of the cluster\n\t\t\t\t\tfails.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "influx_auth_parameters_secret_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the AWS Secrets Manager secret containing the \n\t\t\t\t\tinitial InfluxDB authorization parameters. The secret value is a JSON formatted \n\t\t\t\t\tkey-value pair holding InfluxDB authorization values: organization, bucket, \n\t\t\t\t\tusername, and password.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name that uniquely identifies the DB cluster when interacting with the \n\t\t\t\t\tAmazon Timestream for InfluxDB API and CLI commands. This name will also be a \n\t\t\t\t\tprefix included in the endpoint. DB cluster names must be unique per customer \n\t\t\t\t\tand per region.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_type": {
        "computed": true,
        "description": "Specifies whether the networkType of the Timestream for InfluxDB cluster is \n\t\t\t\t\tIPV4, which can communicate over IPv4 protocol only, or DUAL, which can communicate \n\t\t\t\t\tover both IPv4 and IPv6 protocols.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "organization": {
        "description": "The name of the initial organization for the initial admin user in InfluxDB. An \n\t\t\t\t\tInfluxDB organization is a workspace for a group of users.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "password": {
        "description": "The password of the initial admin user created in InfluxDB. This password will \n\t\t\t\t\tallow you to access the InfluxDB UI to perform various administrative tasks and \n\t\t\t\t\talso use the InfluxDB CLI to create an operator token. These attributes will be \n\t\t\t\t\tstored in a Secret created in AWS SecretManager in your account.",
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "port": {
        "computed": true,
        "description": "The port number on which InfluxDB accepts connections.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "publicly_accessible": {
        "computed": true,
        "description": "Configures the Timestream for InfluxDB cluster with a public IP to facilitate access.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "reader_endpoint": {
        "computed": true,
        "description": "The endpoint used to connect to the Timestream for InfluxDB cluster for \n\t\t\t\t\tread-only operations.",
        "description_kind": "plain",
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
      "username": {
        "description": "The username of the initial admin user created in InfluxDB. \n\t\t\t\t\tMust start with a letter and can't end with a hyphen or contain two \n\t\t\t\t\tconsecutive hyphens. For example, my-user1. This username will allow \n\t\t\t\t\tyou to access the InfluxDB UI to perform various administrative tasks \n\t\t\t\t\tand also use the InfluxDB CLI to create an operator token. These \n\t\t\t\t\tattributes will be stored in a Secret created in Amazon Secrets \n\t\t\t\t\tManager in your account.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vpc_security_group_ids": {
        "description": "A list of VPC security group IDs to associate with the Timestream for InfluxDB cluster.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "vpc_subnet_ids": {
        "description": "A list of VPC subnet IDs to associate with the DB cluster. Provide at least \n\t\t\t\t\ttwo VPC subnet IDs in different availability zones when deploying with a Multi-AZ standby.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "block_types": {
      "log_delivery_configuration": {
        "block": {
          "block_types": {
            "s3_configuration": {
              "block": {
                "attributes": {
                  "bucket_name": {
                    "description": "The name of the S3 bucket to deliver logs to.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "enabled": {
                    "description": "Indicates whether log delivery to the S3 bucket is enabled.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "Configuration for S3 bucket log delivery.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Configuration for sending InfluxDB engine logs to a specified S3 bucket.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours).",
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

func AwsTimestreaminfluxdbDbClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsTimestreaminfluxdbDbCluster), &result)
	return &result
}
