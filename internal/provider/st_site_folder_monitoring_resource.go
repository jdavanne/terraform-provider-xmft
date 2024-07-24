package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

/*
{
  "name": "string",
  "type": "folder",
  "protocol": "folder",
  "transferType": "internal",
  "maxConcurrentConnection": 0,
  "default": false,
  "accessLevel": "PRIVATE",
  "account": "string",
  "additionalAttributes": {
    "additionalProp1": "string",
    "additionalProp2": "string",
    "additionalProp3": "string"
  },
  "downloadFolderAdvancedExpressionEnabled": false,
  "downloadFolder": "string",
  "downloadPatternAdvancedExpressionEnabled": false,
  "downloadPatternType": "string",
  "downloadPattern": "string",
  "downloadPatternCaseSensitive": true,
  "downloadSubfolderMaxDepth": 1,
  "downloadSubfolderPatternType": "string",
  "downloadSubfolderPattern": "string",
  "downloadSubfolderPatternCaseSensitive": true,
  "uploadFolder": "string",
  "uploadFolderExpressionSupport": false,
  "uploadFolderAutocreate": false,
  "uploadFolderOverridable": false,
  "postTransmissionActions": {
    "deleteOnTempFailOut": false,
    "deleteOnSuccessIn": false,
    "deleteOnPermFailOut": false,
    "deleteOnPermFailIn": false,
    "doAsOut": "string",
    "doAsIn": "string",
    "moveOnTempFailOut": "string",
    "moveOnPermFailOut": "string",
    "moveOnSuccessOut": "string",
    "moveOnPermFailIn": "string",
    "moveOnSuccessIn": "string"
  }
}
*/

type stTransferSiteFolderMonitorModel struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	Type                    types.String `tfsdk:"type" helper:",default:folder"`
	Protocol                types.String `tfsdk:"protocol" helper:",default:folder"`
	TransferType            types.String `tfsdk:"transfer_type" helper:"transferType,default:internal"`
	MaxConcurrentConnection types.Int64  `tfsdk:"max_concurrent_connection" helper:"maxConcurrentConnection,default:0"`
	Default                 types.Bool   `tfsdk:"default" helper:",default:false"`
	AccessLevel             types.String `tfsdk:"access_level" helper:"accessLevel,default:PRIVATE"`
	Account                 types.String `tfsdk:"account" helper:",required"`

	DownloadFolderAdvancedExpressionEnabled  types.Bool   `tfsdk:"download_folder_advanced_expression_enabled" helper:"downloadFolderAdvancedExpressionEnabled,default:false"`
	DownloadFolder                           types.String `tfsdk:"download_folder" helper:"downloadFolder,default:"`
	DownloadPatternAdvancedExpressionEnabled types.Bool   `tfsdk:"download_pattern_advanced_expression_enabled" helper:"downloadPatternAdvancedExpressionEnabled,default:false"`
	DownloadPatternType                      types.String `tfsdk:"download_pattern_type" helper:"downloadPatternType,default:glob"`
	DownloadPattern                          types.String `tfsdk:"download_pattern" helper:"downloadPattern,default:*"`
	DownloadPatternCaseSensitive             types.Bool   `tfsdk:"download_pattern_case_sensitive" helper:"downloadPatternCaseSensitive,default:true"`
	DownloadSubfolderMaxDepth                types.Int64  `tfsdk:"download_subfolder_max_depth" helper:"downloadSubfolderMaxDepth,default:1"`
	DownloadSubfolderPatternType             types.String `tfsdk:"download_subfolder_pattern_type" helper:"downloadSubfolderPatternType,default:glob"`
	DownloadSubfolderPattern                 types.String `tfsdk:"download_subfolder_pattern" helper:"downloadSubfolderPattern,default:"`
	DownloadSubfolderPatternCaseSensitive    types.Bool   `tfsdk:"download_subfolder_pattern_case_sensitive" helper:"downloadSubfolderPatternCaseSensitive,default:true"`
	UploadFolder                             types.String `tfsdk:"upload_folder" helper:"uploadFolder,default:"`
	UploadFolderExpressionSupport            types.Bool   `tfsdk:"upload_folder_expression_support" helper:"uploadFolderExpressionSupport,default:false"`
	UploadFolderAutocreate                   types.Bool   `tfsdk:"upload_folder_autocreate" helper:"uploadFolderAutocreate,default:false"`
	UploadFolderOverridable                  types.Bool   `tfsdk:"upload_folder_overridable" helper:"uploadFolderOverridable,default:false"`

	/*PostTransmissionActions                  *struct {
		DeleteOnTempFailOut types.Bool   `tfsdk:"delete_on_temp_fail_out" helper:"deleteOnTempFailOut,default:false"`
		DeleteOnSuccessIn   types.Bool   `tfsdk:"delete_on_success_in" helper:"deleteOnSuccessIn,default:false"`
		DeleteOnPermFailOut types.Bool   `tfsdk:"delete_on_perm_fail_out" helper:"deleteOnPermFailOut,default:false"`
		DeleteOnPermFailIn  types.Bool   `tfsdk:"delete_on_perm_fail_in" helper:"deleteOnPermFailIn,default:false"`
		DoAsOut             types.String `tfsdk:"do_as_out" helper:"doAsOut,default"`
		DoAsIn              types.String `tfsdk:"do_as_in" helper:"doAsIn,default"`
		MoveOnTempFailOut   types.String `tfsdk:"move_on_temp_fail_out" helper:"moveOnTempFailOut,default"`
		MoveOnPermFailOut   types.String `tfsdk:"move_on_perm_fail_out" helper:"moveOnPermFailOut,default"`
		MoveOnSuccessOut    types.String `tfsdk:"move_on_success_out" helper:"moveOnSuccessOut,default"`
		MoveOnPermFailIn    types.String `tfsdk:"move_on_perm_fail_in" helper:"moveOnPermFailIn,default"`
		MoveOnSuccessIn     types.String `tfsdk:"move_on_success_in" helper:"moveOnSuccessIn,default"`
		DoMoveOverwriteIn   types.Bool   `tfsdk:"do_move_overwrite_in" helper:"doMoveOverwriteIn,default:false"`
		DoMoveOverwriteOut  types.Bool   `tfsdk:"do_move_overwrite_out" helper:"doMoveOverwriteOut,default:false"`
	} `tfsdk:"post_transmission_actions" helper:"postTransmissionActions,computed,optional"`*/

	AlternativeAddresses []struct {
		Host     types.String `tfsdk:"host" helper:",required"`
		Port     types.String `tfsdk:"port" helper:",required"`
		Position types.Int64  `tfsdk:"position" helper:",required"`
	} `tfsdk:"alternative_addresses" helper:"alternativeAddresses"`

	AdditionalAttributes types.Map `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`
}

func NewSTTransferSiteFolderMonitorModelResource() resource.Resource {
	return NewSTResource(&stTransferSiteFolderMonitorModel{}, "st_site_folder_monitoring", "", "/api/v2.0/sites", "/api/v2.0/sites/{id}").AddDiscriminator("[type=folder]")
}

func init() {
	registerResource(NewSTTransferSiteFolderMonitorModelResource)
}
