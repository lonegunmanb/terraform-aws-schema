package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsS3BucketReplicationConfiguration = `{
  "block": {
    "attributes": {
      "bucket": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "role": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "rule": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "delete_marker_replication": [
                "list",
                [
                  "object",
                  {
                    "status": "string"
                  }
                ]
              ],
              "destination": [
                "list",
                [
                  "object",
                  {
                    "access_control_translation": [
                      "list",
                      [
                        "object",
                        {
                          "owner": "string"
                        }
                      ]
                    ],
                    "account": "string",
                    "bucket": "string",
                    "encryption_configuration": [
                      "list",
                      [
                        "object",
                        {
                          "replica_kms_key_id": "string"
                        }
                      ]
                    ],
                    "metrics": [
                      "list",
                      [
                        "object",
                        {
                          "event_threshold": [
                            "list",
                            [
                              "object",
                              {
                                "minutes": "number"
                              }
                            ]
                          ],
                          "status": "string"
                        }
                      ]
                    ],
                    "replication_time": [
                      "list",
                      [
                        "object",
                        {
                          "status": "string",
                          "time": [
                            "list",
                            [
                              "object",
                              {
                                "minutes": "number"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "storage_class": "string"
                  }
                ]
              ],
              "existing_object_replication": [
                "list",
                [
                  "object",
                  {
                    "status": "string"
                  }
                ]
              ],
              "filter": [
                "list",
                [
                  "object",
                  {
                    "and": [
                      "list",
                      [
                        "object",
                        {
                          "prefix": "string",
                          "tag": [
                            "list",
                            [
                              "object",
                              {
                                "key": "string",
                                "value": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "prefix": "string",
                    "tag": [
                      "list",
                      [
                        "object",
                        {
                          "key": "string",
                          "value": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "id": "string",
              "prefix": "string",
              "priority": "number",
              "source_selection_criteria": [
                "list",
                [
                  "object",
                  {
                    "replica_modifications": [
                      "list",
                      [
                        "object",
                        {
                          "status": "string"
                        }
                      ]
                    ],
                    "sse_kms_encrypted_objects": [
                      "list",
                      [
                        "object",
                        {
                          "status": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "status": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsS3BucketReplicationConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsS3BucketReplicationConfiguration), &result)
	return &result
}
