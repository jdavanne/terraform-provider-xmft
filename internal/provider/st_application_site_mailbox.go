package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type stSiteMailboxApplication struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	BusinessUnits        []types.String `tfsdk:"business_units" helper:"businessUnits,default:"`
	AdditionalAttributes types.Map      `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`

	DropFolder   types.String `tfsdk:"drop_folder" helper:"dropFolder,default:"`
	InboxFolder  types.String `tfsdk:"inbox_folder" helper:"inboxFolder,default:"`
	ManagedByCG  types.Bool   `tfsdk:"managed_by_cg" helper:"managedByCG,default:false"`
	Notes        types.String `tfsdk:"notes" helper:"notes,default:"`
	OutboxFolder types.String `tfsdk:"outbox_folder" helper:"outboxFolder,default:"`
	Type         types.String `tfsdk:"type" helper:"type,default:SiteMailbox"`
}

func NewSTSiteMailboxApplicationResource() resource.Resource {
	return NewSTResource(&stSiteMailboxApplication{}, "st_site_mailbox_application", "", "/api/v2.0/applications", "/api/v2.0/applications/{name}").UseSwaggerUri("/api/v2.0/applications/{name}[type=SiteMailbox]")
}

func init() {
	registerResource(NewSTSiteMailboxApplicationResource)
}
