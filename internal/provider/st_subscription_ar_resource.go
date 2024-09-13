package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

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
	FolderMonitorScheduleCheck types.String `tfsdk:"folder_monitor_schedule_check" helper:"folderMonitorScheduleCheck,emptyIsNull,default:"`
	FlowName                   types.String `tfsdk:"flow_name" helper:"flowName,,emptyIsNull,default:"`
	ScheduledFolderMonitor     types.String `tfsdk:"scheduled_folder_monitor" helper:"scheduledFolderMonitor,,emptyIsNull,default:"`
	SubscriptionEncryptMode    types.String `tfsdk:"subscription_encrypt_mode" helper:"subscriptionEncryptMode,default:DEFAULT"`
	FileRetentionPeriod        types.Int64  `tfsdk:"file_retention_period" helper:"fileRetentionPeriod,default:0"`
	FlowAttributes             types.Map    `tfsdk:"flow_attributes" helper:"flowAttributes,elementtype:string,optional"`

	// Schedules              []Schedule                `tfsdk:"schedules" helper:"schedules,fold:type,optional"`
	Schedules types.List `tfsdk:"schedules" helper:"schedules,elementtype:Schedule,fold:type,optional"`
	// TransferConfigurations []stTransferConfiguration `tfsdk:"transfer_configurations" helper:"transferConfigurations,optional"`
	TransferConfigurations types.List `tfsdk:"transfer_configurations" helper:"transferConfigurations,elementtype:stTransferConfiguration,optional"`

	PostClientDownloads struct {
		PostClientDownloadActionType                                   types.String `tfsdk:"post_client_download_action_type" helper:"postClientDownloadActionType,emptyIsNull,default:"`
		PostClientDownloadActionTypeFailure                            types.String `tfsdk:"post_client_download_action_type_failure" helper:"postClientDownloadActionTypeFailure,emptyIsNull,default:"`
		PostClientDownloadTypeOnPermfailDoDelete                       types.Bool   `tfsdk:"post_client_download_type_on_permfail_do_delete" helper:"postClientDownloadTypeOnPermfailDoDelete,emptyIsNull,default:"`
		PostClientDownloadTypeOnSuccessDoAdvancedRouting               types.Bool   `tfsdk:"post_client_download_type_on_success_do_advanced_routing" helper:"postClientDownloadTypeOnSuccessDoAdvancedRouting,emptyIsNull,default:"`
		PostClientDownloadTypeOnSuccessDoAdvancedRoutingProcessFile    types.Bool   `tfsdk:"post_client_download_type_on_success_do_advanced_routing_process_file" helper:"postClientDownloadTypeOnSuccessDoAdvancedRoutingProcessFile,emptyIsNull,default:"`
		PostClientDownloadTypeOnFailDoAdvancedRouting                  types.Bool   `tfsdk:"post_client_download_type_on_fail_do_advanced_routing" helper:"postClientDownloadTypeOnFailDoAdvancedRouting,emptyIsNull,default:"`
		PostClientDownloadTypeOnFailDoAdvancedRoutingProcessFailedFile types.Bool   `tfsdk:"post_client_download_type_on_fail_do_advanced_routing_process_failed_file" helper:"postClientDownloadTypeOnFailDoAdvancedRoutingProcessFailedFile,emptyIsNull,default:"`
	} `tfsdk:"post_client_downloads" helper:"postClientDownloads,default:"`

	PostProcessingActions struct {
		PpaOnFailInDoDelete    types.Bool   `tfsdk:"ppa_on_fail_in_do_delete" helper:"ppaOnFailInDoDelete,emptyIsNull,default:"`
		PpaOnFailInDoMove      types.String `tfsdk:"ppa_on_fail_in_do_move" helper:"ppaOnFailInDoMove,emptyIsNull,default:"`
		PpaOnSuccessInDoDelete types.Bool   `tfsdk:"ppa_on_success_in_do_delete" helper:"ppaOnSuccessInDoDelete,emptyIsNull,default:"`
		PpaOnSuccessInDoMove   types.String `tfsdk:"ppa_on_success_in_do_move" helper:"ppaOnSuccessInDoMove,emptyIsNull,default:"`
	} `tfsdk:"post_processing_actions" helper:"postProcessingActions,default:"`

	// PostTransmissionActions stSubscriptionPostTransmissionActions `tfsdk:"post_transmission_actions" helper:"postTransmissionActions,default:"`
	PostTransmissionActions types.Object `tfsdk:"post_transmission_actions" helper:"postTransmissionActions,elementtype:stSubscriptionPostTransmissionActions,default:"`
}

func NewSTSubscriptionARModelResource() resource.Resource {
	return NewSTResource(&stSubscriptionModel{}, "st_subscription_ar", "", "/api/v2.0/subscriptions", "/api/v2.0/subscriptions/{id}").AddDiscriminator("[type=AdvancedRouting]")
}

func init() {
	registerResource(NewSTSubscriptionARModelResource)
}
