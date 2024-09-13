package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type stSharedFolderSubscriptionModel struct {
	Id types.String `tfsdk:"id" helper:",computed,state"`
	// Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	Type                       types.String `tfsdk:"type" helper:",default:SharedFolder"`
	Folder                     types.String `tfsdk:"folder" helper:",required"`
	Account                    types.String `tfsdk:"account" helper:",required"`
	Application                types.String `tfsdk:"application" helper:",required"`
	MaxParallelSitPulls        types.Int64  `tfsdk:"max_parallel_sit_pulls" helper:"maxParallelSitPulls,default:0"`
	FlowAttrsMergeMode         types.String `tfsdk:"flow_attrs_merge_mode" helper:"flowAttrsMergeMode,enum:/preserve/overwrite/append,default:preserve"`
	FolderMonitorScheduleCheck types.String `tfsdk:"folder_monitor_schedule_check" helper:"folderMonitorScheduleCheck,,emptyIsNull,default:"`
	FlowName                   types.String `tfsdk:"flow_name" helper:"flowName,,emptyIsNull,default:"`
	ScheduledFolderMonitor     types.String `tfsdk:"scheduled_folder_monitor" helper:"scheduledFolderMonitor,emptyIsNull,default:"`
	SubscriptionEncryptMode    types.String `tfsdk:"subscription_encrypt_mode" helper:"subscriptionEncryptMode,enum:/default/enabled/disabled,default:default"`
	FileRetentionPeriod        types.Int64  `tfsdk:"file_retention_period" helper:"fileRetentionPeriod,default:0"`
	FlowAttributes             types.Map    `tfsdk:"flow_attributes" helper:"flowAttributes,elementtype:string,default:"`

	Schedules []Schedule `tfsdk:"schedules" helper:"schedules,fold:type,optional"`

	TransferConfigurations     []stTransferConfiguration `tfsdk:"transfer_configurations" helper:"transferConfigurations"`
	TemplateAdvancedExpression types.Bool                `tfsdk:"template_advanced_expression" helper:"templateAdvancedExpression,default:false"`

	PostTransmissionActions struct {
		PtaOnTempfailInDoDelete types.Bool   `tfsdk:"pta_on_tempfail_in_do_delete" helper:"ptaOnTempfailInDoDelete,default:false"`
		PtaOnTempfailInDoMove   types.String `tfsdk:"pta_on_tempfail_in_do_move" helper:"ptaOnTempfailInDoMove,emptyIsNull,default:"`
		PtaOnPermfailInDoDelete types.Bool   `tfsdk:"pta_on_permfail_in_do_delete" helper:"ptaOnPermfailInDoDelete,default:"`
		PtaOnPermfailInDoMove   types.String `tfsdk:"pta_on_permfail_in_do_move" helper:"ptaOnPermfailInDoMove,,emptyIsNull,default:"`
		PtaOnSuccessInDoDelete  types.Bool   `tfsdk:"pta_on_success_in_do_delete" helper:"ptaOnSuccessInDoDelete,default:"`
		PtaOnSuccessInDoMove    types.String `tfsdk:"pta_on_success_in_do_move" helper:"ptaOnSuccessInDoMove,,emptyIsNull,default:"`
	} `tfsdk:"post_transmission_actions" helper:"postTransmissionActions,default:"`
}

func NewSTSharedFolderSubscriptionModelResource() resource.Resource {
	return NewSTResource(&stSharedFolderSubscriptionModel{}, "st_subscription_shared_folder", "", "/api/v2.0/subscriptions", "/api/v2.0/subscriptions/{id}").AddDiscriminator("[type=SharedFolder]")
}

func init() {
	registerResource(NewSTSharedFolderSubscriptionModelResource)
}
