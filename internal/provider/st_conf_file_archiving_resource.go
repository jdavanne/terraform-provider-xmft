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

type stConfigurationFileArchivingModel struct {
	// Id          types.String `tfsdk:"id" helper:",computed,state"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`
	Name        types.String `tfsdk:"name" helper:",default:sentinel,noread,nowrite"`

	// Enabled                         types.Bool   `tfsdk:"enabled" tsdk:"isEnabled:true"`
	GlobalArchivingPolicy           types.String `tfsdk:"global_archiving_policy" helper:"globalArchivingPolicy,default:disabled"`
	ArchiveFolder                   types.String `tfsdk:"archive_folder" helper:"archiveFolder"`
	EncryptionCertificate           types.String `tfsdk:"encryption_certificate" helper:"encryptionCertificate,default:"`
	DeleteFilesOlderThan            types.Int64  `tfsdk:"delete_files_older_than" helper:"deleteFilesOlderThan,default:1"`
	DeleteFilesOlderThanUnit        types.String `tfsdk:"delete_files_older_than_unit" helper:"deleteFilesOlderThanUnit,default:days"`
	MaximumFileSizeAllowedToArchive types.Int64  `tfsdk:"maximum_file_size_allowed_to_archive" helper:"maximumFileSizeAllowedToArchive,default:0"`
}

func NewSTFileArchivingModelResource() resource.Resource {
	return NewSTResource(&stConfigurationFileArchivingModel{}, "st_file_archiving", "", "/api/v2.0/configurations/fileArchiving", "/api/v2.0/configurations/fileArchiving").OnlyReplace()
}

func init() {
	registerResource(NewSTFileArchivingModelResource)
}
