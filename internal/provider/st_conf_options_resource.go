package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

/*
{
  "isEnabled": false,
  "globalArchivingPolicy": "disabled",
  "archiveFolder": "string",
  "encryptionCertificate": "empty",
  "deleteFilesOlderThan": 1,
  "deleteFilesOlderThanUnit": "days",
  "maximumFileSizeAllowedToArchive": 0
}
*/

type stConfigurationOptionsModel struct {
	// Id          types.String `tfsdk:"id" helper:",computed,state"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`
	Name        types.String `tfsdk:"name" helper:",required"`
	Value       types.String `tfsdk:"value" helper:",required,fieldMapOnRead:values.0"`
	ReadOnly    types.Bool   `tfsdk:"read_only" helper:"readOnly,computed"`
	Composite   types.Bool   `tfsdk:"composite" helper:",computed"`
	Disabled    types.Bool   `tfsdk:"disabled" helper:",computed"`
	Encrypted   types.Bool   `tfsdk:"encrypted" helper:",computed"`
	Specified   types.Bool   `tfsdk:"specified" helper:",computed"`
	Description types.String `tfsdk:"description" helper:",computed"`
	IsLocal     types.Bool   `tfsdk:"is_local" helper:"isLocal,computed"`
}

func NewSTConfigurationOptionsModelResource() resource.Resource {
	return NewSTResource(&stConfigurationOptionsModel{}, "st_conf_option", "", "-", "/api/v2.0/configurations/options/{name}").OnlyReplace()
}

func init() {
	registerResource(NewSTConfigurationOptionsModelResource)
}
