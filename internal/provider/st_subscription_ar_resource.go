package provider

import (
	"terraform-provider-xmft/internal/tfhelper"

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

type stSubscriptionPostTransmissionActions struct {
	MoveOverwrite                                     types.Bool   `tfsdk:"move_overwrite" helper:"moveOverwrite,default:false"`
	PtaOnTempfailInDoDelete                           types.Bool   `tfsdk:"pta_on_tempfail_in_do_delete" helper:"ptaOnTempfailInDoDelete,default:false"`
	PtaOnTempfailInDoMove                             types.String `tfsdk:"pta_on_tempfail_in_do_move" helper:"ptaOnTempfailInDoMove,default:"`
	PtaOnPermfailInDoDelete                           types.Bool   `tfsdk:"pta_on_permfail_in_do_delete" helper:"ptaOnPermfailInDoDelete,default:"`
	PtaOnPermfailInDoMove                             types.String `tfsdk:"pta_on_permfail_in_do_move" helper:"ptaOnPermfailInDoMove,default:"`
	PtaOnPermfailDoAdvancedRouting                    types.Bool   `tfsdk:"pta_on_permfail_do_advanced_routing" helper:"ptaOnPermfailDoAdvancedRouting,default:"`
	PtaOnPermfailInDoAdvancedRoutingFailedFile        types.Bool   `tfsdk:"pta_on_permfail_in_do_advanced_routing_failed_file" helper:"ptaOnPermfailInDoAdvancedRoutingFailedFile,default:"`
	PtaOnPermfailInDoAdvancedRoutingWildcardPull      types.Bool   `tfsdk:"pta_on_permfail_in_do_advanced_routing_wildcard_pull" helper:"ptaOnPermfailInDoAdvancedRoutingWildcardPull,default:"`
	PtaOnTempfailInDoAdvancedRouting                  types.Bool   `tfsdk:"pta_on_tempfail_in_do_advanced_routing" helper:"ptaOnTempfailInDoAdvancedRouting,default:"`
	PtaOnTempfailInDoAdvancedRoutingProcessFailedFile types.Bool   `tfsdk:"pta_on_tempfail_in_do_advanced_routing_process_failed_file" helper:"ptaOnTempfailInDoAdvancedRoutingProcessFailedFile,default:"`
	PtaOnTempfailInDoAdvancedRoutingWildcardPull      types.Bool   `tfsdk:"pta_on_tempfail_in_do_advanced_routing_wildcard_pull" helper:"ptaOnTempfailInDoAdvancedRoutingWildcardPull,default:"`
	PtaOnSuccessDoInAdvancedRoutingWildcardPull       types.Bool   `tfsdk:"pta_on_success_do_in_advanced_routing_wildcard_pull" helper:"ptaOnSuccessDoInAdvancedRoutingWildcardPull,default:"`
	PtaOnSuccessTriggerRouteExecutionOnPeSITAck       types.Bool   `tfsdk:"pta_on_success_trigger_route_execution_on_pesit_ack" helper:"ptaOnSuccessTriggerRouteExecutionOnPeSITAck,default:"`
	PtaOnSuccessInDoDelete                            types.Bool   `tfsdk:"pta_on_success_in_do_delete" helper:"ptaOnSuccessInDoDelete,default:"`
	PtaOnSuccessInDoMove                              types.String `tfsdk:"pta_on_success_in_do_move" helper:"ptaOnSuccessInDoMove,default:"`
	PtaOnSuccessInDoMoveOverwrite                     types.Bool   `tfsdk:"pta_on_success_in_do_move_overwrite" helper:"ptaOnSuccessInDoMoveOverwrite,default:"`
	PtaOnPermfailOutDoDelete                          types.Bool   `tfsdk:"pta_on_permfail_out_do_delete" helper:"ptaOnPermfailOutDoDelete,default:"`
	PtaOnPermfailOutDoMove                            types.String `tfsdk:"pta_on_permfail_out_do_move" helper:"ptaOnPermfailOutDoMove,default:"`
	PtaOnSuccessOutDoDelete                           types.Bool   `tfsdk:"pta_on_success_out_do_delete" helper:"ptaOnSuccessOutDoDelete,default:"`
	PtaOnSuccessOutDoMove                             types.String `tfsdk:"pta_on_success_out_do_move" helper:"ptaOnSuccessOutDoMove,default:"`
	PtaOnSuccessOutDoMoveOverwrite                    types.Bool   `tfsdk:"pta_on_success_out_do_move_overwrite" helper:"ptaOnSuccessOutDoMoveOverwrite,default:"`
	PtaOnTempfailOutDoDelete                          types.Bool   `tfsdk:"pta_on_tempfail_out_do_delete" helper:"ptaOnTempfailOutDoDelete,default:"`
	PtaOnTempfailOutDoMove                            types.String `tfsdk:"pta_on_tempfail_out_do_move" helper:"ptaOnTempfailOutDoMove,default:"`
	TriggerOnConditionEnabled                         types.Bool   `tfsdk:"trigger_on_condition_enabled" helper:"triggerOnConditionEnabled,default:"`
	TriggerOnConditionExpression                      types.String `tfsdk:"trigger_on_condition_expression" helper:"triggerOnConditionExpression,default:"`
	TriggerOnSuccessfulWildcardPull                   types.Bool   `tfsdk:"trigger_on_successful_wildcard_pull" helper:"triggerOnSuccessfulWildcardPull,default:"`
	SubmitFilterType                                  types.String `tfsdk:"submit_filter_type" helper:"submitFilterType,default:"`
	SubmitFilenamePatternExpression                   types.String `tfsdk:"submit_filename_pattern_expression" helper:"submitFilenamePatternExpression,default:"`
	TriggerFileOption                                 types.String `tfsdk:"trigger_file_option" helper:"triggerFileOption,default:"`
	TriggerFileRetriesNumber                          types.Int64  `tfsdk:"trigger_file_retries_number" helper:"triggerFileRetriesNumber,default:"`
	TriggerFileRetryDelay                             types.Int64  `tfsdk:"trigger_file_retry_delay" helper:"triggerFileRetryDelay,default:"`
}
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
	FlowAttributes             types.Map    `tfsdk:"flow_attributes" helper:"flowAttributes,elementtype:string,optional"`
	Schedules                  []Schedule   `tfsdk:"schedules" helper:"schedules,fold:type,optional"`

	TransferConfigurations []stTransferConfiguration `tfsdk:"transfer_configurations" helper:"transferConfigurations"`

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

	// PostTransmissionActions types.Object `tfsdk:"post_transmission_actions" helper:"postTransmissionActions,elementtype:stSubscriptionPostTransmissionActions,default:"`
}

func NewSTSubscriptionARModelResource() resource.Resource {
	return NewSTResource(&stSubscriptionModel{}, "st_subscription_ar", "", "/api/v2.0/subscriptions", "/api/v2.0/subscriptions/{id}").AddDiscriminator("[type=AdvancedRouting]")
}

func init() {
	registerResource(NewSTSubscriptionARModelResource)
	tfhelper.RegisterType("stSubscriptionPostTransmissionActions", &stSubscriptionPostTransmissionActions{})
}
