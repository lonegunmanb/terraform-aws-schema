package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRekognitionStreamProcessor = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_id": {
        "description": "The identifier for your AWS Key Management Service key (AWS KMS key). You can supply the Amazon Resource Name (ARN) of your KMS key, the ID of your KMS key, an alias for your KMS key, or an alias ARN.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "An identifier you assign to the stream processor.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_arn": {
        "description": "The Amazon Resource Number (ARN) of the IAM role that allows access to the stream processor.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "stream_processor_arn": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
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
      }
    },
    "block_types": {
      "data_sharing_preference": {
        "block": {
          "attributes": {
            "opt_in": {
              "description": "Do you want to share data with Rekognition to improve model performance.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description": "Shows whether you are sharing data with Rekognition to improve model performance.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "input": {
        "block": {
          "block_types": {
            "kinesis_video_stream": {
              "block": {
                "attributes": {
                  "arn": {
                    "description": "ARN of the Kinesis video stream stream that streams the source video.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Kinesis video stream stream that provides the source streaming video for a Amazon Rekognition Video stream processor.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Information about the source streaming video.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "notification_channel": {
        "block": {
          "attributes": {
            "sns_topic_arn": {
              "description": "The Amazon Resource Number (ARN) of the Amazon Amazon Simple Notification Service topic to which Amazon Rekognition posts the completion status.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The Amazon Simple Notification Service topic to which Amazon Rekognition publishes the object detection results and completion status of a video analysis operation.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "output": {
        "block": {
          "block_types": {
            "kinesis_data_stream": {
              "block": {
                "attributes": {
                  "arn": {
                    "description": "ARN of the output Amazon Kinesis Data Streams stream.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The Amazon Kinesis Data Streams stream to which the Amazon Rekognition stream processor streams the analysis results.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "s3_destination": {
              "block": {
                "attributes": {
                  "bucket": {
                    "description": "The name of the Amazon S3 bucket you want to associate with the streaming video project.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "key_prefix": {
                    "description": "The prefix value of the location within the bucket that you want the information to be published to.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The Amazon S3 bucket location to which Amazon Rekognition publishes the detailed inference results of a video analysis operation.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Kinesis data stream stream or Amazon S3 bucket location to which Amazon Rekognition Video puts the analysis results.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "regions_of_interest": {
        "block": {
          "block_types": {
            "bounding_box": {
              "block": {
                "attributes": {
                  "height": {
                    "description": "Height of the bounding box as a ratio of the overall image height.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "left": {
                    "description": "Left coordinate of the bounding box as a ratio of overall image width.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "top": {
                    "description": "Top coordinate of the bounding box as a ratio of overall image height.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "width": {
                    "description": "Width of the bounding box as a ratio of the overall image width.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "The box representing a region of interest on screen.",
                "description_kind": "plain"
              },
              "nesting_mode": "single"
            },
            "polygon": {
              "block": {
                "attributes": {
                  "x": {
                    "description": "The value of the X coordinate for a point on a Polygon.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "y": {
                    "description": "The value of the Y coordinate for a point on a Polygon.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Specifies a shape made of 3 to 10 Point objects that define a region of interest.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "settings": {
        "block": {
          "block_types": {
            "connected_home": {
              "block": {
                "attributes": {
                  "labels": {
                    "description": "Specifies what you want to detect in the video, such as people, packages, or pets.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "min_confidence": {
                    "computed": true,
                    "description": "The minimum confidence required to label an object in the video.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Label detection settings to use on a streaming video.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "face_search": {
              "block": {
                "attributes": {
                  "collection_id": {
                    "description": "The ID of a collection that contains faces that you want to search for.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "face_match_threshold": {
                    "computed": true,
                    "description": "Minimum face match confidence score that must be met to return a result for a recognized face.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Face search settings to use on a streaming video.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Input parameters used in a streaming video analyzed by a stream processor.",
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

func AwsRekognitionStreamProcessorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRekognitionStreamProcessor), &result)
	return &result
}
