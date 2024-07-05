package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const modtmModuleSource = `{
  "block": {
    "attributes": {
      "module_path": {
        "description": "The path of the module that the telemetry resource is associated with. From this data the provider will attempt to read the ` + "`" + `$TF_DATA_DIR/modules/modules.json` + "`" + ` file and will send the module source and version to the telemetry endpoint.",
        "description_kind": "markdown",
        "required": true,
        "type": "string"
      },
      "module_source": {
        "computed": true,
        "description": "The source of the module that the telemetry resource is associated with",
        "description_kind": "markdown",
        "type": "string"
      },
      "module_version": {
        "computed": true,
        "description": "The version of the module that the telemetry resource is associated with",
        "description_kind": "markdown",
        "type": "string"
      }
    },
    "description": "` + "`" + `modtm_module_source` + "`" + ` data source is used to read the source and version that the current module is associated with. It tried to read ` + "`" + `modules.json` + "`" + ` file in ` + "`" + `.terraform/modules` + "`" + ` folder during the plan time.",
    "description_kind": "markdown"
  },
  "version": 0
}`

func ModtmModuleSourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(modtmModuleSource), &result)
	return &result
}
