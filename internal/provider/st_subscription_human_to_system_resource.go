package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type stHumanToSystemSubscriptionModel struct {
	Id types.String `tfsdk:"id" helper:",computed,state"`
	// Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	Type                       types.String `tfsdk:"type" helper:",default:HumanSystem"`
	Folder                     types.String `tfsdk:"folder" helper:",required"`
	Account                    types.String `tfsdk:"account" helper:",required"`
	Application                types.String `tfsdk:"application" helper:",required"`
	MaxParallelSitPulls        types.Int64  `tfsdk:"max_parallel_sit_pulls" helper:"maxParallelSitPulls,default:0"`
	FlowAttrsMergeMode         types.String `tfsdk:"flow_attrs_merge_mode" helper:"flowAttrsMergeMode,enum:/preserve/overwrite/append,default:preserve"`
	FolderMonitorScheduleCheck types.String `tfsdk:"folder_monitor_schedule_check" helper:"folderMonitorScheduleCheck,emptyIsNull,default:"`
	FlowName                   types.String `tfsdk:"flow_name" helper:"flowName,default"`
	ScheduledFolderMonitor     types.String `tfsdk:"scheduled_folder_monitor" helper:"scheduledFolderMonitor,default:"`
	SubscriptionEncryptMode    types.String `tfsdk:"subscription_encrypt_mode" helper:"subscriptionEncryptMode,enum:/default/enabled/disabled,default:default"`
	FileRetentionPeriod        types.Int64  `tfsdk:"file_retention_period" helper:"fileRetentionPeriod,default:0"`
	FlowAttributes             types.Map    `tfsdk:"flow_attributes" helper:"flowAttributes,elementtype:string,default:"`

	Schedules []Schedule `tfsdk:"schedules" helper:"schedules,fold:type,optional"`

	TransferConfigurations []stTransferConfiguration `tfsdk:"transfer_configurations" helper:"transferConfigurations"`

	Rules []struct {
		Enabled           types.Bool   `tfsdk:"enabled" helper:"enabled,default:false"`
		RecipientPattern  types.String `tfsdk:"recipient_pattern" helper:"recipientPattern,default:"`
		FileFilterPattern types.String `tfsdk:"file_filter_pattern" helper:"fileFilterPattern,default:"`
		TargetFolder      types.String `tfsdk:"target_folder" helper:"targetFolder,default:"`
	} `tfsdk:"rules" helper:"rules,optional"`
}

func NewSTHumanToSystemSubscriptionModelResource() resource.Resource {
	return NewSTResource(&stHumanToSystemSubscriptionModel{}, "st_subscription_human_to_system", "", "/api/v2.0/subscriptions", "/api/v2.0/subscriptions/{id}").AddDiscriminator("[type=HumanSystem]")
}

func init() {
	registerResource(NewSTHumanToSystemSubscriptionModelResource)
}
