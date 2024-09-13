package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type stTransferSiteAS2Model struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	Type                    types.String `tfsdk:"type" helper:",default:AS2"`
	Protocol                types.String `tfsdk:"protocol" helper:",default:AS2"`
	TransferType            types.String `tfsdk:"transfer_type" helper:"transferType,default:internal"`
	MaxConcurrentConnection types.Int64  `tfsdk:"max_concurrent_connection" helper:"maxConcurrentConnection,default:0"`
	Default                 types.Bool   `tfsdk:"default" helper:",default:false"`
	AccessLevel             types.String `tfsdk:"access_level" helper:"accessLevel,default:PRIVATE"`
	Account                 types.String `tfsdk:"account" helper:",required"`
	As2Url                  types.String `tfsdk:"as2_url" helper:"as2Url,required"`
	Dmz                     types.String `tfsdk:"dmz" helper:",default:none"`
	FipsMode                types.Bool   `tfsdk:"fips_mode" helper:"fipsMode,default:false"`
	CipherSuites            types.String `tfsdk:"cipher_suites" helper:"cipherSuites,computed,optional"`
	Protocols               types.String `tfsdk:"protocols" helper:"protocols,default:"`
	AlternativeAddresses    []struct {
		Id       types.String `tfsdk:"id" helper:",computed,state"`
		Host     types.String `tfsdk:"host" helper:",required"`
		Port     types.String `tfsdk:"port" helper:",required"`
		Position types.Int64  `tfsdk:"position" helper:",required"`
	} `tfsdk:"alternative_addresses" helper:"alternativeAddresses,optional"`

	ReceiverAs2Id               types.String `tfsdk:"receiver_as2_id" helper:"receiverAs2Id,default:"`
	ReceiverEmail               types.String `tfsdk:"receiver_email" helper:"receiverEmail,default:"`
	RemoteEncryptionCertificate types.String `tfsdk:"remote_encryption_certificate" helper:"remoteEncryptionCertificate,emptyIsNull,default:"`
	RemoteSigningCertificate    types.String `tfsdk:"remote_signing_certificate" helper:"remoteSigningCertificate,emptyIsNull,default:"`
	SenderAs2Id                 types.String `tfsdk:"sender_as2_id" helper:"senderAs2Id,default:"`
	SenderEmail                 types.String `tfsdk:"sender_email" helper:"senderEmail,default:"`
	ServerEncryptionCertificate types.String `tfsdk:"server_encryption_certificate" helper:"serverEncryptionCertificate,emptyIsNull,default:"`
	ServerSigningCertificate    types.String `tfsdk:"server_signing_certificate" helper:"serverSigningCertificate,emptyIsNull,default:"`
	Subject                     types.String `tfsdk:"subject" helper:"subject,default:"`

	TransferSettings *struct {
		As2MdnOptions     types.String `tfsdk:"as2_mdn_options" helper:"as2MdnOptions,default:"`
		As2MdnReturnUrl   types.String `tfsdk:"as2_mdn_url" helper:"as2MdnReturnUrl,default:"`
		As2MdnUrl         types.Bool   `tfsdk:"as2_mdn_url" helper:"as2MdnUrl,default:"`
		ClientCertificate types.String `tfsdk:"client_certificate" helper:"clientCertificate,emptyIsNull,default:"`
		Compress          types.Bool   `tfsdk:"compress" helper:"compress,default:false"`
		DoAsIn            types.String `tfsdk:"do_as_in" helper:"doAsIn,default"`
		DoAsOut           types.String `tfsdk:"do_as_out" helper:"doAsOut,default"`
		EnableChunked     types.Bool   `tfsdk:"enable_chunked" helper:"enableChunked,default:false"`
		Encrypt           types.String `tfsdk:"encrypt" helper:"encrypt,emptyIsNull,default:"`
		MimeType          types.String `tfsdk:"mime_type" helper:"mimeType,default:"`
		ReceiverPassword  types.String `tfsdk:"receiver_password" helper:"receiverPassword,default:"`
		ReceiverUserName  types.String `tfsdk:"receiver_user_name" helper:"receiverUserName,default:"`
		RequireBA         types.Bool   `tfsdk:"require_ba" helper:"requireBA,default:false"`
		RequireEncrypt    types.Bool   `tfsdk:"require_encrypt" helper:"requireEncrypt,default:false"`
		RequireSign       types.Bool   `tfsdk:"require_sign" helper:"requireSign,default:false"`
		RequireSsl        types.Bool   `tfsdk:"require_ssl" helper:"requireSsl,default:false"`
		SenderPassword    types.String `tfsdk:"sender_password" helper:"senderPassword,default:"`
		SenderUserName    types.String `tfsdk:"sender_user_name" helper:"senderUserName,default:"`
		Sign              types.String `tfsdk:"sign" helper:"sign,emptyIsNull,default:"`
		Subject           types.String `tfsdk:"subject" helper:"subject,default:"`
		TransferTimeout   types.Int64  `tfsdk:"transfer_timeout" helper:"transferTimeout,default:600"`
	} `tfsdk:"transfer_settings" helper:"transferSettings,optional"`

	AdditionalAttributes types.Map `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`
}

func NewSTTransferSiteAS2ModelResource() resource.Resource {
	return NewSTResource(&stTransferSiteAS2Model{}, "st_site_as2", "", "/api/v2.0/sites", "/api/v2.0/sites/{id}").AddDiscriminator("[type=AS2]")
}

func init() {
	registerResource(NewSTTransferSiteAS2ModelResource)
}
