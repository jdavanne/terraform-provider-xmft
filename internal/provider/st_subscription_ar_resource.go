package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

/*
{
  "type": "AdvancedRouting",
  "folder": "string",
  "account": "string",
  "application": "string",
  "maxParallelSitPulls": 0,
  "flowAttrsMergeMode": "preserve",
  "folderMonitorScheduleCheck": "string",
  "flowName": "string",
  "scheduledFolderMonitor": "string",
  "subscriptionEncryptMode": "string",
  "fileRetentionPeriod": 0,
  "flowAttributes": {
    "additionalProp1": "string",
    "additionalProp2": "string",
    "additionalProp3": "string"
  },
  "schedules": [
    {
      "tag": "string",
      "type": "ONCE",
      "executionTimes": [
        "string"
      ],
      "startDate": {},
      "skipHolidays": false
    }
  ],
  "transferConfigurations": [
    {
      "tag": "PARTNER-IN",
      "outbound": false,
      "site": "string",
      "transferProfile": "string",
      "dataTransformations": [
        {
          "type": "string",
          "asciiArmor": false,
          "compressionAlgorithm": 0,
          "compressionLevel": 0,
          "encryptEnabled": false,
          "localSignCertificate": "string",
          "originalNameExpression": "string",
          "originalNameExpressionEnabled": false,
          "partnerEncryptCertificate": "string",
          "requireEncryption": false,
          "requireSignature": false,
          "signingEnabled": false,
          "transformedNameExpression": "string",
          "transformedNameExpressionEnabled": false
        }
      ]
    }
  ],
  "postClientDownloads": {
    "postClientDownloadActionType": "string",
    "postClientDownloadActionTypeFailure": "string",
    "postClientDownloadTypeOnPermfailDoDelete": false,
    "postClientDownloadTypeOnSuccessDoAdvancedRouting": false,
    "postClientDownloadTypeOnSuccessDoAdvancedRoutingProcessFile": false,
    "postClientDownloadTypeOnFailDoAdvancedRouting": false,
    "postClientDownloadTypeOnFailDoAdvancedRoutingProcessFailedFile": false
  },
  "postProcessingActions": {
    "ppaOnFailInDoDelete": false,
    "ppaOnFailInDoMove": "string",
    "ppaOnSuccessInDoDelete": false,
    "ppaOnSuccessInDoMove": "string"
  },
  "postTransmissionActions": {
    "moveOverwrite": false,
    "ptaOnTempfailInDoDelete": false,
    "ptaOnTempfailInDoMove": "string",
    "ptaOnPermfailInDoDelete": false,
    "ptaOnPermfailInDoMove": "string",
    "ptaOnPermfailDoAdvancedRouting": false,
    "ptaOnPermfailInDoAdvancedRoutingFailedFile": false,
    "ptaOnPermfailInDoAdvancedRoutingWildcardPull": false,
    "ptaOnTempfailInDoAdvancedRouting": false,
    "ptaOnTempfailInDoAdvancedRoutingProcessFailedFile": false,
    "ptaOnTempfailInDoAdvancedRoutingWildcardPull": false,
    "ptaOnSuccessDoInAdvancedRoutingWildcardPull": false,
    "ptaOnSuccessTriggerRouteExecutionOnPeSITAck": false,
    "ptaOnSuccessInDoDelete": false,
    "ptaOnSuccessInDoMove": "string",
    "ptaOnSuccessInDoMoveOverwrite": false,
    "ptaOnPermfailOutDoDelete": false,
    "ptaOnPermfailOutDoMove": "string",
    "ptaOnSuccessOutDoDelete": false,
    "ptaOnSuccessOutDoMove": "string",
    "ptaOnSuccessOutDoMoveOverwrite": false,
    "ptaOnTempfailOutDoDelete": false,
    "ptaOnTempfailOutDoMove": "string",
    "triggerOnConditionEnabled": false,
    "triggerOnConditionExpression": "string",
    "triggerOnSuccessfulWildcardPull": false,
    "submitFilterType": "string",
    "submitFilenamePatternExpression": "string",
    "triggerFileOption": "string",
    "triggerFileRetriesNumber": 0,
    "triggerFileRetryDelay": 0
  }
}
*/

type stSubscriptionModel struct {
	Id types.String `tfsdk:"id" helper:",computed,state"`
	// Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	Type                       types.String `tfsdk:"type" helper:",default:AdvancedRouting"`
	Folder                     types.String `tfsdk:"folder" helper:",required"`
	Account                    types.String `tfsdk:"account" helper:",required"`
	Application                types.String `tfsdk:"application" helper:",required"`
	MaxParallelSitPulls        types.Int64  `tfsdk:"max_parallel_sit_pulls" helper:"maxParallelSitPulls,default:0"`
	FlowAttrsMergeMode         types.String `tfsdk:"flow_attrs_merge_mode" helper:"flowAttrsMergeMode,default:preserve"`
	FolderMonitorScheduleCheck types.String `tfsdk:"folder_monitor_schedule_check" helper:"folderMonitorScheduleCheck,default:"`
	FlowName                   types.String `tfsdk:"flow_name" helper:"flowName,default"`
	ScheduledFolderMonitor     types.String `tfsdk:"scheduled_folder_monitor" helper:"scheduledFolderMonitor,default"`
	SubscriptionEncryptMode    types.String `tfsdk:"subscription_encrypt_mode" helper:"subscriptionEncryptMode,default:DEFAULT"`
	FileRetentionPeriod        types.Int64  `tfsdk:"file_retention_period" helper:"fileRetentionPeriod,default:0"`
	// FlowAttributes             types.Object `tfsdk:"flow_attributes" helper:"flowAttributes"`
	Schedules []struct {
		Tag            types.String   `tfsdk:"tag" helper:"tag"`
		Type           types.String   `tfsdk:"type" helper:"type"`
		ExecutionTimes []types.String `tfsdk:"execution_times" helper:"executionTimes"`
		StartDate      types.String   `tfsdk:"start_date" helper:"startDate"`
		SkipHolidays   types.Bool     `tfsdk:"skip_holidays" helper:"skipHolidays"`
	} `tfsdk:"schedules" helper:"schedules"`

	/*TransferConfigurations []struct {
		Tag                 types.String `tfsdk:"tag" helper:"tag"`
		Outbound            types.Bool   `tfsdk:"outbound" helper:"outbound"`
		Site                types.String `tfsdk:"site" helper:"site"`
		TransferProfile     types.String `tfsdk:"transfer_profile" helper:"transferProfile"`
		DataTransformations []struct {
			Type                             types.String `tfsdk:"type" helper:"type"`
			AsciiArmor                       types.Bool   `tfsdk:"ascii_armor" helper:"asciiArmor"`
			CompressionAlgorithm             types.Int64  `tfsdk:"compression_algorithm" helper:"compressionAlgorithm"`
			CompressionLevel                 types.Int64  `tfsdk:"compression_level" helper:"compressionLevel"`
			EncryptEnabled                   types.Bool   `tfsdk:"encrypt_enabled" helper:"encryptEnabled"`
			LocalSignCertificate             types.String `tfsdk:"local_sign_certificate" helper:"localSignCertificate"`
			OriginalNameExpression           types.String `tfsdk:"original_name_expression" helper:"originalNameExpression"`
			OriginalNameExpressionEnabled    types.Bool   `tfsdk:"original_name_expression_enabled" helper:"originalNameExpressionEnabled"`
			PartnerEncryptCertificate        types.String `tfsdk:"partner_encrypt_certificate" helper:"partnerEncryptCertificate"`
			RequireEncryption                types.Bool   `tfsdk:"require_encryption" helper:"requireEncryption"`
			RequireSignature                 types.Bool   `tfsdk:"require_signature" helper:"requireSignature"`
			SigningEnabled                   types.Bool   `tfsdk:"signing_enabled" helper:"signingEnabled"`
			TransformedNameExpression        types.String `tfsdk:"transformed_name_expression" helper:"transformedNameExpression"`
			TransformedNameExpressionEnabled types.Bool   `tfsdk:"transformed_name_expression_enabled" helper:"transformedNameExpressionEnabled"`
		} `tfsdk:"data_transformations" helper:"dataTransformations"`
	} `tfsdk:"transfer_configurations" helper:"transferConfigurations"`*/

	/*PostClientDownloads struct {
		PostClientDownloadActionType                                   types.String `tfsdk:"post_client_download_action_type" helper:"postClientDownloadActionType"`
		PostClientDownloadActionTypeFailure                            types.String `tfsdk:"post_client_download_action_type_failure" helper:"postClientDownloadActionTypeFailure"`
		PostClientDownloadTypeOnPermfailDoDelete                       types.Bool   `tfsdk:"post_client_download_type_on_permfail_do_delete" helper:"postClientDownloadTypeOnPermfailDoDelete"`
		PostClientDownloadTypeOnSuccessDoAdvancedRouting               types.Bool   `tfsdk:"post_client_download_type_on_success_do_advanced_routing" helper:"postClientDownloadTypeOnSuccessDoAdvancedRouting"`
		PostClientDownloadTypeOnSuccessDoAdvancedRoutingProcessFile    types.Bool   `tfsdk:"post_client_download_type_on_success_do_advanced_routing_process_file" helper:"postClientDownloadTypeOnSuccessDoAdvancedRoutingProcessFile"`
		PostClientDownloadTypeOnFailDoAdvancedRouting                  types.Bool   `tfsdk:"post_client_download_type_on_fail_do_advanced_routing" helper:"postClientDownloadTypeOnFailDoAdvancedRouting"`
		PostClientDownloadTypeOnFailDoAdvancedRoutingProcessFailedFile types.Bool   `tfsdk:"post_client_download_type_on_fail_do_advanced_routing_process_failed_file" helper:"postClientDownloadTypeOnFailDoAdvancedRoutingProcessFailedFile"`
	} `tfsdk:"post_client_downloads" helper:"postClientDownloads"`*/

	/*PostProcessingActions struct {
		PpaOnFailInDoDelete    types.Bool   `tfsdk:"ppa_on_fail_in_do_delete" helper:"ppaOnFailInDoDelete"`
		PpaOnFailInDoMove      types.String `tfsdk:"ppa_on_fail_in_do_move" helper:"ppaOnFailInDoMove"`
		PpaOnSuccessInDoDelete types.Bool   `tfsdk:"ppa_on_success_in_do_delete" helper:"ppaOnSuccessInDoDelete"`
		PpaOnSuccessInDoMove   types.String `tfsdk:"ppa_on_success_in_do_move" helper:"ppaOnSuccessInDoMove"`
	} `tfsdk:"post_processing_actions" helper:"postProcessingActions"`*/

	/*PostTransmissionActions *struct {
		MoveOverwrite                                     types.Bool   `tfsdk:"move_overwrite" helper:"moveOverwrite"`
		PtaOnTempfailInDoDelete                           types.Bool   `tfsdk:"pta_on_tempfail_in_do_delete" helper:"ptaOnTempfailInDoDelete"`
		PtaOnTempfailInDoMove                             types.String `tfsdk:"pta_on_tempfail_in_do_move" helper:"ptaOnTempfailInDoMove"`
		PtaOnPermfailInDoDelete                           types.Bool   `tfsdk:"pta_on_permfail_in_do_delete" helper:"ptaOnPermfailInDoDelete"`
		PtaOnPermfailInDoMove                             types.String `tfsdk:"pta_on_permfail_in_do_move" helper:"ptaOnPermfailInDoMove"`
		PtaOnPermfailDoAdvancedRouting                    types.Bool   `tfsdk:"pta_on_permfail_do_advanced_routing" helper:"ptaOnPermfailDoAdvancedRouting"`
		PtaOnPermfailInDoAdvancedRoutingFailedFile        types.Bool   `tfsdk:"pta_on_permfail_in_do_advanced_routing_failed_file" helper:"ptaOnPermfailInDoAdvancedRoutingFailedFile"`
		PtaOnPermfailInDoAdvancedRoutingWildcardPull      types.Bool   `tfsdk:"pta_on_permfail_in_do_advanced_routing_wildcard_pull" helper:"ptaOnPermfailInDoAdvancedRoutingWildcardPull"`
		PtaOnTempfailInDoAdvancedRouting                  types.Bool   `tfsdk:"pta_on_tempfail_in_do_advanced_routing" helper:"ptaOnTempfailInDoAdvancedRouting"`
		PtaOnTempfailInDoAdvancedRoutingProcessFailedFile types.Bool   `tfsdk:"pta_on_tempfail_in_do_advanced_routing_process_failed_file" helper:"ptaOnTempfailInDoAdvancedRoutingProcessFailedFile"`
		PtaOnTempfailInDoAdvancedRoutingWildcardPull      types.Bool   `tfsdk:"pta_on_tempfail_in_do_advanced_routing_wildcard_pull" helper:"ptaOnTempfailInDoAdvancedRoutingWildcardPull"`
		PtaOnSuccessDoInAdvancedRoutingWildcardPull       types.Bool   `tfsdk:"pta_on_success_do_in_advanced_routing_wildcard_pull" helper:"ptaOnSuccessDoInAdvancedRoutingWildcardPull"`
		PtaOnSuccessTriggerRouteExecutionOnPeSITAck       types.Bool   `tfsdk:"pta_on_success_trigger_route_execution_on_pesit_ack" helper:"ptaOnSuccessTriggerRouteExecutionOnPeSITAck"`
		PtaOnSuccessInDoDelete                            types.Bool   `tfsdk:"pta_on_success_in_do_delete" helper:"ptaOnSuccessInDoDelete"`
		PtaOnSuccessInDoMove                              types.String `tfsdk:"pta_on_success_in_do_move" helper:"ptaOnSuccessInDoMove"`
		PtaOnSuccessInDoMoveOverwrite                     types.Bool   `tfsdk:"pta_on_success_in_do_move_overwrite" helper:"ptaOnSuccessInDoMoveOverwrite"`
		PtaOnPermfailOutDoDelete                          types.Bool   `tfsdk:"pta_on_permfail_out_do_delete" helper:"ptaOnPermfailOutDoDelete"`
		PtaOnPermfailOutDoMove                            types.String `tfsdk:"pta_on_permfail_out_do_move" helper:"ptaOnPermfailOutDoMove"`
		PtaOnSuccessOutDoDelete                           types.Bool   `tfsdk:"pta_on_success_out_do_delete" helper:"ptaOnSuccessOutDoDelete"`
		PtaOnSuccessOutDoMove                             types.String `tfsdk:"pta_on_success_out_do_move" helper:"ptaOnSuccessOutDoMove"`
		PtaOnSuccessOutDoMoveOverwrite                    types.Bool   `tfsdk:"pta_on_success_out_do_move_overwrite" helper:"ptaOnSuccessOutDoMoveOverwrite"`
		PtaOnTempfailOutDoDelete                          types.Bool   `tfsdk:"pta_on_tempfail_out_do_delete" helper:"ptaOnTempfailOutDoDelete"`
		PtaOnTempfailOutDoMove                            types.String `tfsdk:"pta_on_tempfail_out_do_move" helper:"ptaOnTempfailOutDoMove"`
		TriggerOnConditionEnabled                         types.Bool   `tfsdk:"trigger_on_condition_enabled" helper:"triggerOnConditionEnabled"`
		TriggerOnConditionExpression                      types.String `tfsdk:"trigger_on_condition_expression" helper:"triggerOnConditionExpression"`
		TriggerOnSuccessfulWildcardPull                   types.Bool   `tfsdk:"trigger_on_successful_wildcard_pull" helper:"triggerOnSuccessfulWildcardPull"`
		SubmitFilterType                                  types.String `tfsdk:"submit_filter_type" helper:"submitFilterType"`
		SubmitFilenamePatternExpression                   types.String `tfsdk:"submit_filename_pattern_expression" helper:"submitFilenamePatternExpression"`
		TriggerFileOption                                 types.String `tfsdk:"trigger_file_option" helper:"triggerFileOption"`
		TriggerFileRetriesNumber                          types.Int64  `tfsdk:"trigger_file_retries_number" helper:"triggerFileRetriesNumber"`
		TriggerFileRetryDelay                             types.Int64  `tfsdk:"trigger_file_retry_delay" helper:"triggerFileRetryDelay"`
	} `tfsdk:"post_transmission_actions" helper:"postTransmissionActions"`*/
}

func NewSTSubscriptionARModelResource() resource.Resource {
	return NewSTResource(&stSubscriptionModel{}, "st_subscription_ar", "", "/api/v2.0/subscriptions", "/api/v2.0/subscriptions/{id}")
}

func init() {
	registerResource(NewSTSubscriptionARModelResource)
}
