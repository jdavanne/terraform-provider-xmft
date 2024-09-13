package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// see https://support.axway.com/kb/181603/language/en

type stCertificateModel struct {
	Id          types.String `tfsdk:"id" helper:",computed"`
	Name        types.String `tfsdk:"name" helper:",required,noread"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	Type       types.String `tfsdk:"type" helper:",enum:/x509/ssh,default:x509"`
	Usage      types.String `tfsdk:"usage" helper:",enum:/private/local/partner/login/trusted,default:private"`
	Password   types.String `tfsdk:"password" helper:",sensitive,noread"`
	CaPassword types.String `tfsdk:"ca_password" helper:"caPassword,sensitive,noread"`
	Overwrite  types.Bool   `tfsdk:"overwrite" helper:",noread,default:false"`

	Account     types.String `tfsdk:"account" helper:",default:"`
	AccessLevel types.String `tfsdk:"access_level" helper:"accessLevel,computed"`

	Content types.String `tfsdk:"content" helper:",required,noread"`

	AdditionalAttributes types.Map `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`

	Subject          types.String `tfsdk:"subject" helper:",computed"`
	ExpirationTime   types.String `tfsdk:"expiration_time" helper:"expirationTime,computed,nowrite"`
	CreationTime     types.Int64  `tfsdk:"creation_time" helper:"creationTime,computed,nowrite"`
	SignAlgorithm    types.String `tfsdk:"sign_algorithm" helper:"signAlgorithm,computed,nowrite"`
	KeySize          types.Int64  `tfsdk:"key_size" helper:"keySize,computed,nowrite"`
	KeyAlgorithm     types.String `tfsdk:"key_algorithm" helper:"keyAlgorithm,computed,nowrite"`
	Issuer           types.String `tfsdk:"issuer" helper:"issuer,computed,nowrite"`
	SerialNumber     types.String `tfsdk:"serial_number" helper:"serialNumber,computed,nowrite"`
	ValidityPeriod   types.Int64  `tfsdk:"validity_period" helper:"validityPeriod,computed,nowrite"`
	Fingerprint      types.String `tfsdk:"fingerprint" helper:"fingerprint,computed,nowrite"`
	ValidationStatus types.String `tfsdk:"validation_status" helper:"validationStatus,computed,nowrite"`
	Version          types.Int64  `tfsdk:"version" helper:"version,computed,nowrite"`

	// ExportPrivateKey    types.Bool `tfsdk:"export_private_key" helper:"exportPrivateKey,computed,nowrite"`
	// ExportSSHPublicKey  types.Bool `tfsdk:"export_ssh_public_key" helper:"exportSSHPublicKey,computed,nowrite"`
}

func NewSTCertificateModelResource() resource.Resource {
	return NewSTResource(&stCertificateModel{}, "st_certificate", "", "/api/v2.0/certificates", "/api/v2.0/certificates/{id}").AlwaysRecreate()
}

func init() {
	registerResource(NewSTCertificateModelResource)
}
