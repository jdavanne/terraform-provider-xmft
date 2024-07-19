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
	Type   types.String `tfsdk:"type" helper:",default:SendToPartner"`
	Status types.String `tfsdk:"status" helper:",enum:/ENABLED/DISABLED,default:ENABLED"`

	// PrecedingStep types.String `tfsdk:"preceding_step"`
	ConditionType       types.String `tfsdk:"condition_type" helper:"conditionType,enum:/ALWAYS/EL,default:ALWAYS"`
	Condition           types.String `tfsdk:"condition" helper:",emptyIsNull,default:"`
	ActionOnStepSuccess types.String `tfsdk:"action_on_step_success" helper:"actionOnStepSuccess,enum:/PROCEED/STOP,default:PROCEED"`
	ActionOnStepFailure types.String `tfsdk:"action_on_step_failure" helper:"actionOnStepFailure,enum:/FAIL/PROCEED,default:FAIL"`

	// Autostart                            types.Bool   `tfsdk:"autostart" helper:"autostart,default:false"`
	// UsePrecedingStepFiles                types.Bool   `tfsdk:"use_preceding_step_files" helper:"usePrecedingStepFiles,default:false"`
	FileFilterExpression                 types.String `tfsdk:"file_filter_expression" helper:"fileFilterExpression,default:*"`
	FileFilterExpressionType             types.String `tfsdk:"file_filter_expression_type" helper:"fileFilterExpressionType,enum:/GLOB/REGEXP/TEXT_FILES,default:GLOB"`
	TargetAccountExpression              types.String `tfsdk:"target_account_expression" helper:"targetAccountExpression,emptyIsNull,default:"`
	TargetAccountExpressionType          types.String `tfsdk:"target_account_expression_type" helper:"targetAccountExpressionType,enum:/NAME/EXPRESSION,emptyIsNull,default"`
	TransferSiteExpression               types.String `tfsdk:"transfer_site_expression" helper:"transferSiteExpression,default"`
	TransferSiteExpressionType           types.String `tfsdk:"transfer_site_expression_type" helper:"transferSiteExpressionType,enum:/LIST/EXPRESSION_WILDCARD,default:LIST"`
	UploadFolder                         types.String `tfsdk:"upload_folder" helper:"uploadFolder,default:/"`
	TransferProfileExpression            types.String `tfsdk:"transfer_profile_expression" helper:"transferProfileExpression,emptyIsNull,default:"`
	TransferProfileExpressionType        types.String `tfsdk:"transfer_profile_expression_type" helper:"transferProfileExpressionType,computed,optional"`
	StoreAndForwardMode                  types.String `tfsdk:"store_and_forward_mode" helper:"storeAndForwardMode,emptyIsNull,default:"`
	VirtualFilename                      types.String `tfsdk:"virtual_filename" helper:"virtualFilename,emptyIsNull,default:"`
	DataEncoding                         types.String `tfsdk:"data_encoding" helper:"dataEncoding,emptyIsNull,default:"`
	RecordFormat                         types.String `tfsdk:"record_format" helper:"recordFormat,emptyIsNull,default:"`
	RecordLength                         types.String `tfsdk:"record_length" helper:"recordLength,emptyIsNull,default:"`
	FileLabel                            types.String `tfsdk:"file_label" helper:"fileLabel,emptyIsNull,default:"`
	Originator                           types.String `tfsdk:"originator" helper:"originator,emptyIsNull,default:"`
	FinalDestination                     types.String `tfsdk:"final_destination" helper:"finalDestination,emptyIsNull,default:"`
	UserMessage                          types.String `tfsdk:"user_message" helper:"userMessage,emptyIsNull,default:"`
	TriggerFileName                      types.String `tfsdk:"trigger_file_name" helper:"triggerFileName,emptyIsNull,default:"`
	TriggerTargetAccountExpression       types.String `tfsdk:"trigger_target_account_expression" helper:"triggerTargetAccountExpression,emptyIsNull,default:"`
	TriggerTransferSiteExpression        types.String `tfsdk:"trigger_transfer_site_expression" helper:"triggerTransferSiteExpression,emptyIsNull,default:"`
	TriggerTransferSiteExpressionType    types.String `tfsdk:"trigger_transfer_site_expression_type" helper:"triggerTransferSiteExpressionType,enum:/LIST/EXPRESSION_WILDCARD,emptyIsNull,default:"`
	RouteFileAs                          types.String `tfsdk:"route_file_as" helper:"routeFileAs,emptyIsNull,default:"`
	TriggerFileContent                   types.String `tfsdk:"trigger_file_content" helper:"triggerFileContent,computed,optional"`
	PostRoutingActionType                types.String `tfsdk:"post_routing_action_type" helper:"postRoutingActionType,enum:/DELETE/RENAME,emptyIsNull,default:"`
	SleepIncrementBetweenRetries         types.Int64  `tfsdk:"sleep_increment_between_retries" helper:"sleepIncrementBetweenRetries,computed,optional"`
	SleepBetweenRetries                  types.Int64  `tfsdk:"sleep_between_retries" helper:"sleepBetweenRetries,computed,optional"`
	MaxParallelClients                   types.Int64  `tfsdk:"max_parallel_clients" helper:"maxParallelClients,computed,optional"`
	MaxNumberOfRetries                   types.Int64  `tfsdk:"max_number_of_retries" helper:"maxNumberOfRetries,computed,optional"`
	TriggerFileForEach                   types.Bool   `tfsdk:"trigger_file_for_each" helper:"triggerFileForEach,computed,optional"`
	TriggerTransferProfileExpression     types.String `tfsdk:"trigger_transfer_profile_expression" helper:"triggerTransferProfileExpression,computed,optional"`
	TriggerUploadFolder                  types.String `tfsdk:"trigger_upload_folder" helper:"triggerUploadFolder,computed,optional"`
	ArchivePolicyOnFailure               types.String `tfsdk:"archive_policy_on_failure" helper:"archivePolicyOnFailure,computed,optional"`
	ArchivePolicyOnSuccess               types.String `tfsdk:"archive_policy_on_success" helper:"archivePolicyOnSuccess,computed,optional"`
	PostRoutingActionRenameExpression    types.String `tfsdk:"post_routing_action_rename_expression" helper:"postRoutingActionRenameExpression,computed,optional"`
	TriggerTargetAccountExpressionType   types.String `tfsdk:"trigger_target_account_expression_type" helper:"triggerTargetAccountExpressionType,computed,optional"`
	TriggerTransferProfileExpressionType types.String `tfsdk:"trigger_transfer_profile_expression_type" helper:"triggerTransferProfileExpressionType,computed,optional"`
}

func init() {
	tfhelper.RegisterType("stRouteStepSendToPartnerResourceModel", &stRouteStepSendToPartnerResourceModel{})
}
