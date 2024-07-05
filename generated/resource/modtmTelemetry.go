package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const modtmTelemetry = `{
  "block": {
    "attributes": {
      "endpoint": {
        "description": "Telemetry endpoint to send data to, will override provider's default ` + "`" + `endpoint` + "`" + ` setting.\nYou can set ` + "`" + `endpoint` + "`" + ` in this resource, when there's no explicit ` + "`" + `setting` + "`" + ` in the provider block, it will override provider's default ` + "`" + `endpoint` + "`" + `.\n\n|Explicit ` + "`" + `endpoint` + "`" + ` in ` + "`" + `provider` + "`" + ` block | ` + "`" + `MODTM_ENDPOINT` + "`" + ` environment variable set | Explicit ` + "`" + `endpoint` + "`" + ` in resource block | Telemetry endpoint |\n|--|--|--|--|\n| ✓ | ✓ | ✓ | Explicit ` + "`" + `endpoint` + "`" + ` in ` + "`" + `provider` + "`" + ` block | \n| ✓ | ✓ | × | Explicit ` + "`" + `endpoint` + "`" + ` in ` + "`" + `provider` + "`" + ` block | \n| ✓ | × | ✓ | Explicit ` + "`" + `endpoint` + "`" + ` in ` + "`" + `provider` + "`" + ` block | \n| ✓ | × | × | Explicit ` + "`" + `endpoint` + "`" + ` in ` + "`" + `provider` + "`" + ` block | \n| × | ✓ | ✓ | ` + "`" + `MODTM_ENDPOINT` + "`" + ` environment variable | \n| × | ✓ | × | ` + "`" + `MODTM_ENDPOINT` + "`" + ` environment variable | \n| × | × | ✓ | Explicit ` + "`" + `endpoint` + "`" + ` in resource block | \n| × | × | × | Default Microsoft telemetry service endpoint | \n",
        "description_kind": "markdown",
        "optional": true,
        "type": "string"
      },
      "ephemeral_number": {
        "computed": true,
        "deprecated": true,
        "description": "An ephemeral number that works with tags-generation tools like [BridgeCrew Yor](https://yor.io/)",
        "description_kind": "markdown",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description": "Resource identifier",
        "description_kind": "markdown",
        "type": "string"
      },
      "nonce": {
        "computed": true,
        "deprecated": true,
        "description": "A nonce that works with tags-generation tools like [BridgeCrew Yor](https://yor.io/)",
        "description_kind": "markdown",
        "optional": true,
        "type": "number"
      },
      "tags": {
        "description": "Tags to be sent to telemetry endpoint. The following tags are reserved and cannot be used: ` + "`" + `event` + "`" + `. When specififying ` + "`" + `module_path` + "`" + `, the ` + "`" + `source` + "`" + ` and ` + "`" + `version` + "`" + ` tags will be automatically added to the tags sent to the telemetry endpoint.",
        "description_kind": "markdown",
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
