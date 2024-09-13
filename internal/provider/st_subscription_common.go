package provider

import (
	"terraform-provider-xmft/internal/tfhelper"

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
		TransformedNameExpressionEnabled types.Bool   `tfsdk:"transformed_name_expression_enabled" helper:"transformedNameExpressionEnabled,default:false"`
	} `tfsdk:"data_transformations" helper:"dataTransformations"`
}

type stSubscriptionPostTransmissionActions struct {
	MoveOverwrite                                     types.Bool   `tfsdk:"move_overwrite" helper:"moveOverwrite,default:false"`
	PtaOnTempfailInDoDelete                           types.Bool   `tfsdk:"pta_on_tempfail_in_do_delete" helper:"ptaOnTempfailInDoDelete,optional"`
	PtaOnTempfailInDoMove                             types.String `tfsdk:"pta_on_tempfail_in_do_move" helper:"ptaOnTempfailInDoMove,optional"`
	PtaOnPermfailInDoDelete                           types.Bool   `tfsdk:"pta_on_permfail_in_do_delete" helper:"ptaOnPermfailInDoDelete,optional"`
	PtaOnPermfailInDoMove                             types.String `tfsdk:"pta_on_permfail_in_do_move" helper:"ptaOnPermfailInDoMove,optional"`
	PtaOnPermfailDoAdvancedRouting                    types.Bool   `tfsdk:"pta_on_permfail_do_advanced_routing" helper:"ptaOnPermfailDoAdvancedRouting,optional"`
	PtaOnPermfailInDoAdvancedRoutingFailedFile        types.Bool   `tfsdk:"pta_on_permfail_in_do_advanced_routing_failed_file" helper:"ptaOnPermfailInDoAdvancedRoutingFailedFile,optional"`
	PtaOnPermfailInDoAdvancedRoutingWildcardPull      types.Bool   `tfsdk:"pta_on_permfail_in_do_advanced_routing_wildcard_pull" helper:"ptaOnPermfailInDoAdvancedRoutingWildcardPull,optional"`
	PtaOnTempfailInDoAdvancedRouting                  types.Bool   `tfsdk:"pta_on_tempfail_in_do_advanced_routing" helper:"ptaOnTempfailInDoAdvancedRouting,optional"`
	PtaOnTempfailInDoAdvancedRoutingProcessFailedFile types.Bool   `tfsdk:"pta_on_tempfail_in_do_advanced_routing_process_failed_file" helper:"ptaOnTempfailInDoAdvancedRoutingProcessFailedFile,optional"`
	PtaOnTempfailInDoAdvancedRoutingWildcardPull      types.Bool   `tfsdk:"pta_on_tempfail_in_do_advanced_routing_wildcard_pull" helper:"ptaOnTempfailInDoAdvancedRoutingWildcardPull,optional"`
	PtaOnSuccessDoInAdvancedRoutingWildcardPull       types.Bool   `tfsdk:"pta_on_success_do_in_advanced_routing_wildcard_pull" helper:"ptaOnSuccessDoInAdvancedRoutingWildcardPull,optional"`
	PtaOnSuccessTriggerRouteExecutionOnPeSITAck       types.Bool   `tfsdk:"pta_on_success_trigger_route_execution_on_pesit_ack" helper:"ptaOnSuccessTriggerRouteExecutionOnPeSITAck,optional"`
	PtaOnSuccessInDoDelete                            types.Bool   `tfsdk:"pta_on_success_in_do_delete" helper:"ptaOnSuccessInDoDelete,optional"`
	PtaOnSuccessInDoMove                              types.String `tfsdk:"pta_on_success_in_do_move" helper:"ptaOnSuccessInDoMove,optional"`
	PtaOnSuccessInDoMoveOverwrite                     types.Bool   `tfsdk:"pta_on_success_in_do_move_overwrite" helper:"ptaOnSuccessInDoMoveOverwrite,optional"`
	PtaOnPermfailOutDoDelete                          types.Bool   `tfsdk:"pta_on_permfail_out_do_delete" helper:"ptaOnPermfailOutDoDelete,optional"`
	PtaOnPermfailOutDoMove                            types.String `tfsdk:"pta_on_permfail_out_do_move" helper:"ptaOnPermfailOutDoMove,optional"`
	PtaOnSuccessOutDoDelete                           types.Bool   `tfsdk:"pta_on_success_out_do_delete" helper:"ptaOnSuccessOutDoDelete,optional"`
	PtaOnSuccessOutDoMove                             types.String `tfsdk:"pta_on_success_out_do_move" helper:"ptaOnSuccessOutDoMove,optional"`
	PtaOnSuccessOutDoMoveOverwrite                    types.Bool   `tfsdk:"pta_on_success_out_do_move_overwrite" helper:"ptaOnSuccessOutDoMoveOverwrite,optional"`
	PtaOnTempfailOutDoDelete                          types.Bool   `tfsdk:"pta_on_tempfail_out_do_delete" helper:"ptaOnTempfailOutDoDelete,optional"`
	PtaOnTempfailOutDoMove                            types.String `tfsdk:"pta_on_tempfail_out_do_move" helper:"ptaOnTempfailOutDoMove,optional"`
	TriggerOnConditionEnabled                         types.Bool   `tfsdk:"trigger_on_condition_enabled" helper:"triggerOnConditionEnabled,optional"`
	TriggerOnConditionExpression                      types.String `tfsdk:"trigger_on_condition_expression" helper:"triggerOnConditionExpression,optional"`
	TriggerOnSuccessfulWildcardPull                   types.Bool   `tfsdk:"trigger_on_successful_wildcard_pull" helper:"triggerOnSuccessfulWildcardPull,optional"`
	SubmitFilterType                                  types.String `tfsdk:"submit_filter_type" helper:"submitFilterType,enum:/TRIGGER_FILE_CONTENT/FILENAME_PATTERN,default:TRIGGER_FILE_CONTENT"`
	SubmitFilenamePatternExpression                   types.String `tfsdk:"submit_filename_pattern_expression" helper:"submitFilenamePatternExpression,optional"`
	TriggerFileOption                                 types.String `tfsdk:"trigger_file_option" helper:"triggerFileOption,enum:/fail/continue/retry,default:fail"`
	TriggerFileRetriesNumber                          types.Int64  `tfsdk:"trigger_file_retries_number" helper:"triggerFileRetriesNumber,optional"`
	TriggerFileRetryDelay                             types.Int64  `tfsdk:"trigger_file_retry_delay" helper:"triggerFileRetryDelay,optional"`
}

func init() {
	tfhelper.RegisterType("stSubscriptionPostTransmissionActions", &stSubscriptionPostTransmissionActions{})
	tfhelper.RegisterType("stTransferConfiguration", &stTransferConfiguration{})
}
