package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsNeptunegraphGraph = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection": {
        "computed": true,
        "description": "A value that indicates whether the graph has deletion protection enabled. The graph can't be deleted when deletion protection is enabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "graph_name": {
        "computed": true,
        "description": "The graph name. For example: my-graph-1.\n\t\t\t\t\t\t\t\tThe name must contain from 1 to 63 letters, numbers, or hyphens, \n\t\t\t\t\t\t\t\tand its first character must be a letter. It cannot end with a hyphen or contain two consecutive hyphens.\n\t\t\t\t\t\t\t\tIf you don't specify a graph name, a unique graph name is generated for you using the prefix graph-for, \n\t\t\t\t\t\t\t\tfollowed by a combination of Stack Name and a UUID.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "graph_name_prefix": {
        "description": "Allows user to specify name prefix and have remainder of name automatically generated.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_identifier": {
        "computed": true,
        "description": "Specifies a KMS key to use to encrypt data in the new graph.  Value must be ARN of KMS Key.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "provisioned_memory": {
        "description": "The provisioned memory-optimized Neptune Capacity Units (m-NCUs) to use for the graph.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "public_connectivity": {
        "computed": true,
        "description": "Specifies whether or not the graph can be reachable over the internet. \n\t\t\t\t\t\t\t\tAll access to graphs is IAM authenticated.\n\t\t\t\t\t\t\t\tWhen the graph is publicly available, its domain name system (DNS) endpoint resolves to \n\t\t\t\t\t\t\t\tthe public IP address from the internet. When the graph isn't publicly available, you need \n\t\t\t\t\t\t\t\tto create a PrivateGraphEndpoint in a given VPC to ensure the DNS name resolves to a private \n\t\t\t\t\t\t\t\tIP address that is reachable from the VPC.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "replica_count": {
        "computed": true,
        "description": "The number of replicas in other AZs.  Value must be between 0 and 2.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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
      }
    },
    "block_types": {
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
      },
      "vector_search_configuration": {
        "block": {
          "attributes": {
            "vector_search_dimension": {
              "description": "Specifies the number of dimensions for vector embeddings.  Value must be between 1 and 65,535.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description": "Vector search configuration for the Neptune Graph",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsNeptunegraphGraphSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsNeptunegraphGraph), &result)
	return &result
}
