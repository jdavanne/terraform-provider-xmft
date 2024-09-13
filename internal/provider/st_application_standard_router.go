package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type stStandardRouterApplication struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	Type                 types.String   `tfsdk:"type" helper:"type,default:StandardRouter"`
	BusinessUnits        []types.String `tfsdk:"business_units" helper:"businessUnits,default:"`
	AdditionalAttributes types.Map      `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`

	DropFolder     types.String `tfsdk:"drop_folder" helper:"dropFolder,default:"`
	InboxEnabled   types.Bool   `tfsdk:"inbox_enabled" helper:"inboxEnabled,default:false"`
	InboxFolder    types.Bool   `tfsdk:"inbox_folder" helper:"inboxFolder,default:false"`
	InboxIdPattern types.String `tfsdk:"inbox_id_pattern" helper:"inboxIdPattern,default:false"`
	ManagedByCG    types.Bool   `tfsdk:"managed_by_cg" helper:"managedByCG,default:false"`

	Notes                    types.String `tfsdk:"notes" helper:"notes,default:"`
	OutboxEnabled            types.Bool   `tfsdk:"outbox_enabled" helper:"outboxEnabled,default:false"`
	OutboxFileFormat         types.String `tfsdk:"outbox_file_format" helper:"outboxFileFormat,default:false"`
	OutboxFolder             types.String `tfsdk:"outbox_folder" helper:"outboxFolder,default:"`
	RenameFilesEnabled       types.Bool   `tfsdk:"rename_files_enabled" helper:"renameFilesEnabled,default:false"`
	SecureConnectionRequired types.Bool   `tfsdk:"secure_connection_required" helper:"secureConnectionRequired,default:false"`
}

func NewSTStandardRouterApplicationResource() resource.Resource {
	return NewSTResource(&stStandardRouterApplication{}, "st_standard_router_application", "", "/api/v2.0/applications", "/api/v2.0/applications/{name}").UseSwaggerUri("/api/v2.0/applications/{name}[type=StandardRouter]")
}

func init() {
	registerResource(NewSTStandardRouterApplicationResource)
}
