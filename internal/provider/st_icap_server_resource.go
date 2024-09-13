package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type stIcapServerModel struct {
	// Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	AdvancedConnectionSettings struct {
		ConnectionTimeout types.Int64  `tfsdk:"connection_timeout" helper:"connectionTimeout,default:"`
		EnabledCiphers    types.String `tfsdk:"enabled_ciphers" helper:"enabledCiphers,default:"`
		EnabledProtocols  types.String `tfsdk:"enabled_protocols" helper:"enabledProtocols,default:"`
		ReadTimeout       types.Int64  `tfsdk:"read_timeout" helper:"readTimeout,default:"`
		RetryAttempts     types.Int64  `tfsdk:"retry_attempts" helper:"retryAttempts,default:"`
		RetryDelay        types.Int64  `tfsdk:"retry_delay" helper:"retryDelay,default:"`
	} `tfsdk:"advanced_connection_settings" helper:"advancedConnectionSettings,default:"`

	AdvancedIcapSettings struct {
		EnableWinNtFormat               types.Bool `tfsdk:"enable_win_nt_format" helper:"enableWinNtFormat,default:false"`
		StopTransfersModifyOrNotHandled types.Bool `tfsdk:"stop_transfers_modify_or_not_handled" helper:"stopTransfersModifyOrNotHandled,default:false"`
		TreatModifyAsBlock              types.Bool `tfsdk:"treat_modify_as_block" helper:"treatModifyAsBlock,default:false"`
	} `tfsdk:"advanced_icap_settings" helper:"advancedIcapSettings,default:"`

	BasicSettings struct {
		ClientCertificate         types.String `tfsdk:"client_certificate" helper:"clientCertificate,default:"`
		ConnectionSecure          types.Bool   `tfsdk:"connection_secure" helper:"connectionSecure,default:false"`
		DenyOnConnectionError     types.Bool   `tfsdk:"deny_on_connection_error" helper:"denyOnConnectionError,default:false"`
		FipsEnabled               types.Bool   `tfsdk:"fips_enabled" helper:"fipsEnabled,default:false"`
		IncludeFilename           types.Bool   `tfsdk:"include_filename" helper:"includeFilename,default:true"`
		MailAddressesOnIcapDenied types.String `tfsdk:"mail_addresses_on_icap_denied" helper:"mailAddressesOnIcapDenied,default:"`
		MailAddressesOnIcapError  types.String `tfsdk:"mail_addresses_on_icap_error" helper:"mailAddressesOnIcapError,default:"`
		MaxSize                   types.Int64  `tfsdk:"max_size" helper:"maxSize,default:"`
		Name                      types.String `tfsdk:"name" helper:"name,default:string"`
		NotifyOnIcapDenied        types.Bool   `tfsdk:"notify_on_icap_denied" helper:"notifyOnIcapDenied,default:false"`
		NotifyOnIcapError         types.Bool   `tfsdk:"notify_on_icap_error" helper:"notifyOnIcapError,default:false"`
		PreviewSize               types.Int64  `tfsdk:"preview_size" helper:"previewSize,default:"`
		Type                      types.String `tfsdk:"type" helper:"type,enum:[BOTH INCOMING OUTGOING],default:"`
		Url                       types.String `tfsdk:"url" helper:"url,default:"`
		VerifyCertificate         types.Bool   `tfsdk:"verify_certificate" helper:"verifyCertificate,default:false"`
	} `tfsdk:"basic_settings" helper:"basicSettings,default:"`

	HeaderSettings struct {
		HeadersMapping types.String `tfsdk:"headers_mapping" helper:"headersMapping,default:"`
	} `tfsdk:"header_settings" helper:"headerSettings,default:"`

	ScanFilteringSettings struct {
		IgnoredFileType        types.String `tfsdk:"ignored_file_type" helper:"ignoredFileType,default:"`
		OnlyIfPartnerRecipient types.Bool   `tfsdk:"only_if_partner_recipient" helper:"onlyIfPartnerRecipient,default:false"`
		PolicyExpression       types.String `tfsdk:"policy_expression" helper:"policyExpression,default:"`
		WithoutBU              types.Bool   `tfsdk:"without_bu" helper:"withoutBU,default:false"`
	} `tfsdk:"scan_filtering_settings" helper:"scanFilteringSettings,default:"`
	ServerEnabled types.Bool `tfsdk:"server_enabled" helper:"serverEnabled,default:false"`
}

func NewSTIcapServerModelResource() resource.Resource {
	return NewSTResource(&stIcapServerModel{}, "st_icap_server", "", "/api/v2.0/icapServers", "/api/v2.0/icapServers/{name}")
}

func init() {
	registerResource(NewSTIcapServerModelResource)
}
