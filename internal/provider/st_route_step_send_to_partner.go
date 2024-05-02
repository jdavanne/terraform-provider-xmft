package provider

import (
	"terraform-provider-xmft/internal/tfhelper"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

/*
 {
            "type": "SendToPartner",
            "status": "ENABLED",
            "precedingStep": "string",
            "conditionType": "ALWAYS",
            "condition": "string",
            "actionOnStepSuccess": "PROCEED",
            "actionOnStepFailure": "FAIL",
            "autostart": false,
            "usePrecedingStepFiles": false,
            "fileFilterExpression": "string",
            "fileFilterExpressionType": "GLOB",
            "targetAccountExpression": "string",
            "targetAccountExpressionType": "string",
            "transferSiteExpression": "string",
            "transferSiteExpressionType": "LIST",
            "uploadFolder": "string",
            "transferProfileExpression": "string",
            "transferProfileExpressionType": "string",
            "storeAndForwardMode": "string",
            "virtualFilename": "string",
            "dataEncoding": "string",
            "recordFormat": "string",
            "recordLength": "string",
            "fileLabel": "string",
            "originator": "string",
            "finalDestination": "string",
            "userMessage": "string",
            "triggerFileName": "string",
            "triggerTargetAccountExpression": "string",
            "triggerTransferSiteExpression": "string",
            "triggerTransferSiteExpressionType": "string",
            "routeFileAs": "string",
            "triggerFileContent": "string",
            "postRoutingActionType": "string",
            "sleepIncrementBetweenRetries": 0,
            "sleepBetweenRetries": 0,
            "maxParallelClients": 0,
            "maxNumberOfRetries": 0,
            "triggerFileForEach": false,
            "triggerTransferProfileExpression": "string",
            "triggerUploadFolder": "string",
            "archivePolicyOnFailure": "DEFAULT",
            "archivePolicyOnSuccess": "DEFAULT",
            "postRoutingActionRenameExpression": "string",
            "triggerTargetAccountExpressionType": "NAME",
            "triggerTransferProfileExpressionType": "string"
        },
*/

type stRouteStepSendToPartnerResourceModel struct {
	Id     types.String `tfsdk:"id" helper:",computed,state"`
	Type   types.String `tfsdk:"type" helper:",required"`
	Status types.String `tfsdk:"status" helper:",default:ENABLED"`
	// PrecedingStep types.String `tfsdk:"preceding_step"`
	// ConditionType                        types.String `tfsdk:"condition_type" helper:"conditionType,default:ALWAYS"`
	// Condition                            types.String `tfsdk:"condition"`
	ActionOnStepSuccess types.String `tfsdk:"action_on_step_success" helper:"actionOnStepSuccess,default:PROCEED"`
	ActionOnStepFailure types.String `tfsdk:"action_on_step_failure" helper:"actionOnStepFailure,default:FAIL"`
	// Autostart                            types.Bool   `tfsdk:"autostart" helper:"autostart,default:false"`
	// UsePrecedingStepFiles                types.Bool   `tfsdk:"use_preceding_step_files" helper:"usePrecedingStepFiles,default:false"`
	FileFilterExpression        types.String `tfsdk:"file_filter_expression" helper:"fileFilterExpression,default:*"`
	FileFilterExpressionType    types.String `tfsdk:"file_filter_expression_type" helper:"fileFilterExpressionType,default:GLOB"`
	TargetAccountExpression     types.String `tfsdk:"target_account_expression" helper:"targetAccountExpression,default"`
	TargetAccountExpressionType types.String `tfsdk:"target_account_expression_type" helper:"targetAccountExpressionType,default"`
	TransferSiteExpression      types.String `tfsdk:"transfer_site_expression" helper:"transferSiteExpression,default"`
	TransferSiteExpressionType  types.String `tfsdk:"transfer_site_expression_type" helper:"transferSiteExpressionType,default:LIST"`
	UploadFolder                types.String `tfsdk:"upload_folder" helper:"uploadFolder,default"`
	// TransferProfileExpression            types.String `tfsdk:"transfer_profile_expression" helper:"transferProfileExpression"`
	// TransferProfileExpressionType        types.String `tfsdk:"transfer_profile_expression_type" helper:"transferProfileExpressionType"`
	// StoreAndForwardMode                  types.String `tfsdk:"store_and_forward_mode" helper:"storeAndForwardMode"`
	// VirtualFilename                      types.String `tfsdk:"virtual_filename" helper:"virtualFilename"`
	// DataEncoding                         types.String `tfsdk:"data_encoding" helper:"dataEncoding"`
	// RecordFormat                         types.String `tfsdk:"record_format" helper:"recordFormat"`
	// RecordLength                         types.String `tfsdk:"record_length" helper:"recordLength"`
	// FileLabel                            types.String `tfsdk:"file_label" helper:"fileLabel"`
	// Originator                           types.String `tfsdk:"originator" helper:"originator"`
	// FinalDestination                     types.String `tfsdk:"final_destination" helper:"finalDestination"`
	// UserMessage                          types.String `tfsdk:"user_message" helper:"userMessage"`
	// TriggerFileName                      types.String `tfsdk:"trigger_file_name" helper:"triggerFileName"`
	// TriggerTargetAccountExpression       types.String `tfsdk:"trigger_target_account_expression" helper:"triggerTargetAccountExpression"`
	// TriggerTransferSiteExpression        types.String `tfsdk:"trigger_transfer_site_expression" helper:"triggerTransferSiteExpression"`
	// TriggerTransferSiteExpressionType    types.String `tfsdk:"trigger_transfer_site_expression_type" helper:"triggerTransferSiteExpressionType"`
	// RouteFileAs                          types.String `tfsdk:"route_file_as" helper:"routeFileAs"`
	// TriggerFileContent                   types.String `tfsdk:"trigger_file_content" helper:"triggerFileContent"`
	// PostRoutingActionType                types.String `tfsdk:"post_routing_action_type" helper:"postRoutingActionType"`
	// SleepIncrementBetweenRetries         types.Int64  `tfsdk:"sleep_increment_between_retries" helper:"sleepIncrementBetweenRetries"`
	// SleepBetweenRetries                  types.Int64  `tfsdk:"sleep_between_retries" helper:"sleepBetweenRetries"`
	// MaxParallelClients                   types.Int64  `tfsdk:"max_parallel_clients" helper:"maxParallelClients"`
	// MaxNumberOfRetries                   types.Int64  `tfsdk:"max_number_of_retries" helper:"maxNumberOfRetries"`
	// TriggerFileForEach                   types.Bool   `tfsdk:"trigger_file_for_each" helper:"triggerFileForEach"`
	// TriggerTransferProfileExpression     types.String `tfsdk:"trigger_transfer_profile_expression" helper:"triggerTransferProfileExpression"`
	// TriggerUploadFolder                  types.String `tfsdk:"trigger_upload_folder" helper:"triggerUploadFolder"`
	// ArchivePolicyOnFailure               types.String `tfsdk:"archive_policy_on_failure" helper:"archivePolicyOnFailure"`
	// ArchivePolicyOnSuccess               types.String `tfsdk:"archive_policy_on_success" helper:"archivePolicyOnSuccess"`
	// PostRoutingActionRenameExpression    types.String `tfsdk:"post_routing_action_rename_expression" helper:"postRoutingActionRenameExpression"`
	// TriggerTargetAccountExpressionType   types.String `tfsdk:"trigger_target_account_expression_type" helper:"triggerTargetAccountExpressionType"`
	// TriggerTransferProfileExpressionType types.String `tfsdk:"trigger_transfer_profile_expression_type" helper:"triggerTransferProfileExpressionType"`
}

func init() {
	tfhelper.RegisterType("stRouteStepSendToPartnerResourceModel", &stRouteStepSendToPartnerResourceModel{})
}
