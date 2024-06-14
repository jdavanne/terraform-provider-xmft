package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

/*
{
	"name": "string",
	"type": "Basic",
	"notes": "string",
	"managedByCG": false,
	"additionalAttributes": {
	  "additionalProp1": "string",
	  "additionalProp2": "string",
	  "additionalProp3": "string"
	},
	"businessUnits": [
	  "string"
	]
  }
*/

type stBasicApplication struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	Type          types.String   `tfsdk:"type" helper:",default:Basic"`
	Notes         types.String   `tfsdk:"notes"`
	BusinessUnits []types.String `tfsdk:"business_units" helper:"businessUnits,default:"`
	ManagedByCG   types.Bool     `tfsdk:"managed_by_cg" helper:"managedByCG,default:false"`

	AdditionalAttributes types.Map `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`
}

func NewSTBasicApplicationResource() resource.Resource {
	return NewSTResource(&stBasicApplication{}, "st_basic_application", "", "/api/v2.0/applications", "/api/v2.0/applications/{name}")
}

func init() {
	registerResource(NewSTBasicApplicationResource)
}
