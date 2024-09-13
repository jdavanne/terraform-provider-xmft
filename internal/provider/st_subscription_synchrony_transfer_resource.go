package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type stSynchronyTransferModel struct {
	Id types.String `tfsdk:"id" helper:",computed,state"`
	// Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	Account             types.String `tfsdk:"account" helper:"account,default:"`
	Application         types.String `tfsdk:"application" helper:"application,default:"`
	FileRetentionPeriod types.Int64  `tfsdk:"file_retention_period" helper:"fileRetentionPeriod,default:0"`
	FlowAttributes      types.Map    `tfsdk:"flow_attributes" helper:"flowAttributes,elementtype:string,optional"`

	FlowAttrsMergeMode         types.String `tfsdk:"flow_attrs_merge_mode" helper:"flowAttrsMergeMode,enum:/preserve/overwrite/append,default:"`
	FlowName                   types.String `tfsdk:"flow_name" helper:"flowName,emptyIsNull,default:"`
	Folder                     types.String `tfsdk:"folder" helper:"folder,emptyIsNull,default:"`
	FolderMonitorScheduleCheck types.String `tfsdk:"folder_monitor_schedule_check" helper:"folderMonitorScheduleCheck,emptyIsNull,default:"`
	MaxParallelSitPulls        types.Int64  `tfsdk:"max_parallel_sit_pulls" helper:"maxParallelSitPulls,default:"`
	ScheduledFolderMonitor     types.String `tfsdk:"scheduled_folder_monitor" helper:"scheduledFolderMonitor,emptyIsNull,default:"`
	SubscriptionEncryptMode    types.String `tfsdk:"subscription_encrypt_mode" helper:"subscriptionEncryptMode,emptyIsNull,default:"`
	Type                       types.String `tfsdk:"type" helper:"type,default:SynchronyTransfer"`

	Schedules []Schedule `tfsdk:"schedules" helper:"schedules,fold:type,optional"`

	TransferConfigurations []stTransferConfiguration `tfsdk:"transfer_configurations" helper:"transferConfigurations"`
}

func NewSTSynchronyTransferModelResource() resource.Resource {
	return NewSTResource(&stSynchronyTransferModel{}, "st_subscription_synchrony_transfer", "", "/api/v2.0/subscriptions", "/api/v2.0/subscriptions/{id}").AddDiscriminator("[type=SynchronyTransfer]")
}

func init() {
	registerResource(NewSTSynchronyTransferModelResource)
}
