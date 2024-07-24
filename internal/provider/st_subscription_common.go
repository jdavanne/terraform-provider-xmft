package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type stTransferConfiguration struct {
	Id                  types.String `tfsdk:"id" helper:",computed,state"`
	Tag                 types.String `tfsdk:"tag" helper:"tag,enum:/PARTNER-IN/PARTNER-OUT,required"`
	Outbound            types.Bool   `tfsdk:"outbound" helper:"outbound,default:false"`
	Site                types.String `tfsdk:"site" helper:"site"`
	TransferProfile     types.String `tfsdk:"transfer_profile" helper:"transferProfile"`
	DataTransformations []struct {
		Id                               types.String `tfsdk:"id" helper:",computed,state"`
		Type                             types.String `tfsdk:"type" helper:"type"`
		AsciiArmor                       types.Bool   `tfsdk:"ascii_armor" helper:"asciiArmor"`
		CompressionAlgorithm             types.Int64  `tfsdk:"compression_algorithm" helper:"compressionAlgorithm,enum:/prefered:-1/no-compression:0/zip:1/zlib:2/bzip2:3,default:0"`
		CompressionLevel                 types.Int64  `tfsdk:"compression_level" helper:"compressionLevel,enum:/fast:1/normal:5/good:7/best:9,default:2"`
		EncryptEnabled                   types.Bool   `tfsdk:"encrypt_enabled" helper:"encryptEnabled,default:false"`
		LocalSignCertificate             types.String `tfsdk:"local_sign_certificate" helper:"localSignCertificate"`
		OriginalNameExpression           types.String `tfsdk:"original_name_expression" helper:"originalNameExpression"`
		OriginalNameExpressionEnabled    types.Bool   `tfsdk:"original_name_expression_enabled" helper:"originalNameExpressionEnabled,default:false"`
		PartnerEncryptCertificate        types.String `tfsdk:"partner_encrypt_certificate" helper:"partnerEncryptCertificate"`
		RequireEncryption                types.Bool   `tfsdk:"require_encryption" helper:"requireEncryption"`
		RequireSignature                 types.Bool   `tfsdk:"require_signature" helper:"requireSignature"`
		SigningEnabled                   types.Bool   `tfsdk:"signing_enabled" helper:"signingEnabled,default:false"`
		TransformedNameExpression        types.String `tfsdk:"transformed_name_expression" helper:"transformedNameExpression"`
		TransformedNameExpressionEnabled types.Bool   `tfsdk:"transformed_name_expression_enabled" helper:"transformedNameExpressionEnabled:false"`
	} `tfsdk:"data_transformations" helper:"dataTransformations"`
}
