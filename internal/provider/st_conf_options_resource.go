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
}

func NewSTConfigurationOptionsModelResource() resource.Resource {
	return NewSTResource(&stConfigurationOptionsModel{}, "st_conf_option", "", "-", "/api/v2.0/configurations/options/{name}").OnlyReplace()
}

func init() {
	registerResource(NewSTConfigurationOptionsModelResource)
}
