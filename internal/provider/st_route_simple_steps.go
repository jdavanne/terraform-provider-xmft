package provider

import (
	"terraform-provider-xmft/internal/tfhelper"
)

type stRouteSimpleSteps2 struct {
	//Compress           basetypes.ObjectValue `tfsdk:"compress" helper:"compress,elementtype:stRouteStepCompressResourceModel"`
	//Decompress         basetypes.ObjectValue `tfsdk:"decompress" helper:"decompress,elementtype:stRouteStepDecompressResourceModel"`
	//EncodingConversion basetypes.ObjectValue `tfsdk:"encoding_conversion" helper:"encodingConversion,elementtype:stRouteStepEncodingConversionModel"`
	//ExecuteRoute       types.Object `tfsdk:"execute_route" helper:"executeRoute,elementtype:stRouteStepExecuteRouteResourceModel"`
	//ExternalScript     types.Object `tfsdk:"external_script" helper:"externalScript,elementtype:stRouteStepExternalScriptResourceModel"`
	// LineEnding         types.Object `tfsdk:"line_ending" helper:"lineEnding,elementtype:stRouteStepLineEndingResourceModel"`
	/*LineFolding      types.Object `tfsdk:"line_folding" helper:"lineFolding,elementtype:stRouteStepLineFoldingResourceModel"`
	LinePadding      types.Object `tfsdk:"line_padding" helper:"linePadding,elementtype:stRouteStepLinePaddingResourceModel"`
	LineTruncating   types.Object `tfsdk:"line_truncating" helper:"lineTruncating,elementtype:stRouteStepLineTruncatingResourceModel"`
	PGPDecryption    types.Object `tfsdk:"pgp_decryption" helper:"pgpDecryption,elementtype:stRouteStepPgpDecryptionResourceModel"`
	PGPEncryption    types.Object `tfsdk:"pgp_encryption" helper:"pgpEncryption,elementtype:stRouteStepPgpEncryptionResourceModel"`
	Pluggable        types.Object `tfsdk:"pluggable" helper:"pluggable,elementtype:stRouteStepPluggableResourceModel"`*/
	//PublishToAccount types.Object `tfsdk:"publish_to_account" helper:"publishToAccount,elementtype:stRouteStepPublishToAccountResourceModel"`
	//PullFromPartner  types.Object `tfsdk:"pull_from_partner" helper:"pullFromPartner,elementtype:stRouteStepPullFromPartnerResourceModel"`
	//SendToPartner basetypes.ObjectValue `tfsdk:"send_to_partner" helper:"sendToPartner,elementtype:stRouteStepSendToPartnerResourceModel"`
}

type stRouteSimpleSteps struct {
	Compress   *stRouteStepCompressResourceModel   `tfsdk:"compress" helper:"Compress,optional"`
	Decompress *stRouteStepDecompressResourceModel `tfsdk:"decompress" helper:"Decompress,optional"`
	// EncodingConversion *stRouteStepEncodingConversionModel `tfsdk:"encoding_conversion" helper:"EncodingConversion,elementtype:stRouteStepEncodingConversionModel,optional"`
	// ExecuteRoute       types.Object `tfsdk:"execute_route" helper:"executeRoute,elementtype:stRouteStepExecuteRouteResourceModel"`
	// ExternalScript     types.Object `tfsdk:"external_script" helper:"externalScript,elementtype:stRouteStepExternalScriptResourceModel"`
	// LineEnding         types.Object `tfsdk:"line_ending" helper:"lineEnding,elementtype:stRouteStepLineEndingResourceModel"`
	// LineFolding      types.Object `tfsdk:"line_folding" helper:"lineFolding,elementtype:stRouteStepLineFoldingResourceModel"`
	// LinePadding      types.Object `tfsdk:"line_padding" helper:"linePadding,elementtype:stRouteStepLinePaddingResourceModel"`
	// LineTruncating   types.Object `tfsdk:"line_truncating" helper:"lineTruncating,elementtype:stRouteStepLineTruncatingResourceModel"`
	// PGPDecryption    types.Object `tfsdk:"pgp_decryption" helper:"pgpDecryption,elementtype:stRouteStepPgpDecryptionResourceModel"`
	// PGPEncryption    types.Object `tfsdk:"pgp_encryption" helper:"pgpEncryption,elementtype:stRouteStepPgpEncryptionResourceModel"`
	Pluggable        *stRouteStepPluggableResourceModel        `tfsdk:"pluggable" helper:"Pluggable,optional"`
	Rename           *stRouteStepRenameResourceModel           `tfsdk:"rename" helper:"Rename,optional"`
	PublishToAccount *stRouteStepPublishToAccountResourceModel `tfsdk:"publish_to_account" helper:"Publish,optional"`
	PullFromPartner  *stRouteStepPullFromPartnerResourceModel  `tfsdk:"pull_from_partner" helper:"PullFromPartner,optional"`
	SendToPartner    *stRouteStepSendToPartnerResourceModel    `tfsdk:"send_to_partner" helper:"SendToPartner,optional"`
}

func init() {
	tfhelper.RegisterType("stRouteSimpleSteps", &stRouteSimpleSteps{})
}
