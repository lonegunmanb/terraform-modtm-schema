package generated

import (
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-modtm-schema/generated/data"
	"github.com/lonegunmanb/terraform-modtm-schema/generated/resource"
)

var Resources map[string]*tfjson.Schema
var DataSources map[string]*tfjson.Schema

func init() {
	data.PlaceHolder()
	resource.PlaceHolder()
	resources := make(map[string]*tfjson.Schema, 0)
	dataSources := make(map[string]*tfjson.Schema, 0)  
	resources["modtm_telemetry"] = resource.ModtmTelemetrySchema()  
	dataSources["modtm_module_source"] = data.ModtmModuleSourceSchema()  
	Resources = resources
	DataSources = dataSources
}