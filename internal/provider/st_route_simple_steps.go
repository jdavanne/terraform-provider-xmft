package provider

import (
	"terraform-provider-xmft/internal/tfhelper"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type stRouteSimpleSteps struct {
	Compress           types.Object `tfsdk:"compress" helper:"compress,elementtype:stRouteStepCompressResourceModel"`
	Decompress         types.Object `tfsdk:"decompress" helper:"decompress,elementtype:stRouteStepDecompressResourceModel"`
	EncodingConversion types.Object `tfsdk:"encoding_conversion" helper:"encodingConversion,elementtype:stRouteStepEncodingConversionModel"`
	ExecuteRoute       types.Object `tfsdk:"execute_route" helper:"executeRoute,elementtype:stRouteStepExecuteRouteResourceModel"`
	ExternalScript     types.Object `tfsdk:"external_script" helper:"externalScript,elementtype:stRouteStepScriptResourceModel"`
	LineEnding         types.Object `tfsdk:"line_ending" helper:"lineEnding,elementtype:stRouteStepLineEndingResourceModel"`
	LineFolding        types.Object `tfsdk:"line_folding" helper:"lineFolding,elementtype:stRouteStepLineFoldingResourceModel"`
	LinePadding        types.Object `tfsdk:"line_padding" helper:"linePadding,elementtype:stRouteStepLinePaddingResourceModel"`
	LineTruncating     types.Object `tfsdk:"line_truncating" helper:"lineTruncating,elementtype:stRouteStepLineTruncatingResourceModel"`
	PGPDecryption      types.Object `tfsdk:"pgp_decryption" helper:"pgpDecryption,elementtype:stRouteStepPGPDecryptionResourceModel"`
	PGPEncryption      types.Object `tfsdk:"pgp_encryption" helper:"pgpEncryption,elementtype:stRouteStepPGPEncryptionResourceModel"`
	Pluggable          types.Object `tfsdk:"pluggable" helper:"pluggable,elementtype:stRouteStepPluggableResourceModel"`
	PublishToAccount   types.Object `tfsdk:"publish_to_account" helper:"publishToAccount,elementtype:stRouteStepPublishToAccountResourceModel"`
	PullFromPartner    types.Object `tfsdk:"pull_from_partner" helper:"pullFromPartner,elementtype:stRouteStepPullFromPartnerResourceModel"`
	SendToPartner      types.Object `tfsdk:"push_to_partner" helper:"sendToPartner,elementtype:stRouteStepSendToPartnerResourceModel"`
}

func init() {
	tfhelper.RegisterType("stRouteSimpleSteps", &stRouteSimpleSteps{})
}
