package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type stTransferSiteAdhocModel struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	AccessLevel types.String `tfsdk:"access_level" helper:"accessLevel,enum:/PRIVATE/PUBLIC/BUSINESS_UNIT,default:"`
	Account     types.String `tfsdk:"account" helper:"account,default:"`

	AlternativeAddresses []struct {
		Host     types.String `tfsdk:"host" helper:"host,default:"`
		Id       types.String `tfsdk:"id" helper:"id,state,computed"`
		Port     types.String `tfsdk:"port" helper:"port,default:"`
		Position types.Int64  `tfsdk:"position" helper:"position,default:"`
	} `tfsdk:"alternative_addresses" helper:"alternativeAddresses,optional"`

	BccEmailAddresses       types.String `tfsdk:"bcc_email_addresses" helper:"bccEmailAddresses,default:"`
	Body                    types.String `tfsdk:"body" helper:"body,default:"`
	CcEmailAddresses        types.String `tfsdk:"cc_email_addresses" helper:"ccEmailAddresses,default:"`
	Challenge               types.String `tfsdk:"challenge" helper:"challenge,default:"`
	ConfirmResponse         types.String `tfsdk:"confirm_response" helper:"confirmResponse,default:"`
	Default                 types.Bool   `tfsdk:"default" helper:",default:false"`
	Expiration              types.String `tfsdk:"expiration" helper:"expiration,enum:/1440 (1 day)/10080 (7 days)/43200 (30 days)/86400 (60 days)/Never,default:"`
	FromEmailAddress        types.String `tfsdk:"from_email_address" helper:"fromEmailAddress,default:"`
	MaxConcurrentConnection types.Int64  `tfsdk:"max_concurrent_connection" helper:"maxConcurrentConnection,default:0"`
	NotificationTemplate    types.String `tfsdk:"notification_template" helper:"notificationTemplate,default:"`
	Protocol                types.String `tfsdk:"protocol" helper:"protocol,default:adhoc"`
	Response                types.String `tfsdk:"response" helper:"response,default:"`
	SecurityLevel           types.String `tfsdk:"security_level" helper:"securityLevel,enum:/ANONYMOUS_LINK/CHALLENGED_LINK/EXISTING_ACCOUNT/ENROLL_UNLICENSED/ENROLL_LICENSED,default:"`
	SendAs                  types.String `tfsdk:"send_as" helper:"sendAs,default:"`
	SendFileAsAttachment    types.Bool   `tfsdk:"send_file_as_attachment" helper:"sendFileAsAttachment,default:false"`
	Subject                 types.String `tfsdk:"subject" helper:"subject,default:"`
	ToEmailAddresses        types.String `tfsdk:"to_email_addresses" helper:"toEmailAddresses,default:"`
	TransferType            types.String `tfsdk:"transfer_type" helper:"transferType,enum:/false/I/E,default:false"`
	Type                    types.String `tfsdk:"type" helper:"type,default:adhoc"`

	AdditionalAttributes types.Map `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`
}

func NewSTTransferSiteAdhocModelResource() resource.Resource {
	return NewSTResource(&stTransferSiteAdhocModel{}, "st_site_adhoc", "", "/api/v2.0/sites", "/api/v2.0/sites/{id}").AddDiscriminator("[type=adhoc]")
}

func init() {
	registerResource(NewSTTransferSiteAdhocModelResource)
}
