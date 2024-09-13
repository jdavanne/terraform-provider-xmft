package provider

import (
	"terraform-provider-xmft/internal/tfhelper"
)

type stRouteSimpleSteps struct {
	CharactersReplace  *stRouteStepCharactersReplaceResourceModel `tfsdk:"characters_replace" helper:"CharactersReplace,optional"`
	Compress           *stRouteStepCompressResourceModel          `tfsdk:"compress" helper:"Compress,optional"`
	Decompress         *stRouteStepDecompressResourceModel        `tfsdk:"decompress" helper:"Decompress,optional"`
	EncodingConversion *stRouteStepEncodingConversionModel        `tfsdk:"encoding_conversion" helper:"EncodingConversion,optional"`
	ExecuteRoute       *stRouteStepExecuteRouteResourceModel      `tfsdk:"execute_route" helper:"ExecuteRoute,optional"`
	ExternalScript     *stRouteStepExternalScriptResourceModel    `tfsdk:"external_script" helper:"ExternalScript,optional"`
	LineEnding         *stRouteStepLineEndingResourceModel        `tfsdk:"line_ending" helper:"LineEnding,optional"`
	LineFolding        *stRouteStepLineFoldingResourceModel       `tfsdk:"line_folding" helper:"LineFolding,optional"`
	LinePadding        *stRouteStepLinePaddingResourceModel       `tfsdk:"line_padding" helper:"LinePadding,optional"`
	LineTruncating     *stRouteStepLineTruncatingResourceModel    `tfsdk:"line_truncating" helper:"LineTruncating,optional"`
	PGPDecryption      *stRouteStepPgpDecryptionResourceModel     `tfsdk:"pgp_decryption" helper:"PgpDecryption,optional"`
	PGPEncryption      *stRouteStepPgpEncryptionResourceModel     `tfsdk:"pgp_encryption" helper:"PgpEncryption,optional"`
	Pluggable          *stRouteStepPluggableResourceModel         `tfsdk:"pluggable" helper:"Pluggable,optional,optional"`
	Rename             *stRouteStepRenameResourceModel            `tfsdk:"rename" helper:"Rename,optional"`
	PublishToAccount   *stRouteStepPublishToAccountResourceModel  `tfsdk:"publish_to_account" helper:"Publish,optional"`
	PullFromPartner    *stRouteStepPullFromPartnerResourceModel   `tfsdk:"pull_from_partner" helper:"PullFromPartner,optional"`
	SendToPartner      *stRouteStepSendToPartnerResourceModel     `tfsdk:"send_to_partner" helper:"SendToPartner,optional"`
}

func init() {
	tfhelper.RegisterType("stRouteSimpleSteps", &stRouteSimpleSteps{})
}
