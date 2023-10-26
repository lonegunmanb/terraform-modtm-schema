package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const modtmTelemetry = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Resource identifier",
        "description_kind": "markdown",
        "type": "string"
      },
      "nonce": {
        "computed": true,
        "description": "A nonce that work with tags-generation tools like [BridgeCrew Yor](https://yor.io/)",
        "description_kind": "markdown",
        "optional": true,
        "type": "number"
      },
      "tags": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "` + "`" + `modtm_telemetry` + "`" + ` resource gathers and sends telemetry data to a specified endpoint. The aim is to provide visibility into the lifecycle of your Terraform modules - whether they are being created, updated, or deleted.",
    "description_kind": "markdown"
  },
  "version": 0
}`

func ModtmTelemetrySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(modtmTelemetry), &result)
	return &result
}
