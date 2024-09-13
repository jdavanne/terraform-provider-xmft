package provider

import (
	"terraform-provider-xmft/internal/tfhelper"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type stBasicSubscriptionPostTransmissionActions struct {
	PtaOnTempfailInDoDelete  types.Bool   `tfsdk:"pta_on_tempfail_in_do_delete" helper:"ptaOnTempfailInDoDelete,default:false"`
	PtaOnTempfailInDoMove    types.String `tfsdk:"pta_on_tempfail_in_do_move" helper:"ptaOnTempfailInDoMove,emptyIsNull,default:"`
	PtaOnPermfailInDoDelete  types.Bool   `tfsdk:"pta_on_permfail_in_do_delete" helper:"ptaOnPermfailInDoDelete,default:"`
	PtaOnPermfailInDoMove    types.String `tfsdk:"pta_on_permfail_in_do_move" helper:"ptaOnPermfailInDoMove,emptyIsNull,default:"`
	PtaOnSuccessInDoDelete   types.Bool   `tfsdk:"pta_on_success_in_do_delete" helper:"ptaOnSuccessInDoDelete,default:"`
	PtaOnSuccessInDoMove     types.String `tfsdk:"pta_on_success_in_do_move" helper:"ptaOnSuccessInDoMove,emptyIsNull,default:"`
	PtaOnPermfailOutDoDelete types.Bool   `tfsdk:"pta_on_permfail_out_do_delete" helper:"ptaOnPermfailOutDoDelete,default:"`
	PtaOnPermfailOutDoMove   types.String `tfsdk:"pta_on_permfail_out_do_move" helper:"ptaOnPermfailOutDoMove,emptyIsNull,default:"`
	PtaOnSuccessOutDoDelete  types.Bool   `tfsdk:"pta_on_success_out_do_delete" helper:"ptaOnSuccessOutDoDelete,default:"`
	PtaOnSuccessOutDoMove    types.String `tfsdk:"pta_on_success_out_do_move" helper:"ptaOnSuccessOutDoMove,emptyIsNull,default:"`
}

type stBasicSubscriptionModel struct {
	Id types.String `tfsdk:"id" helper:",computed,state"`
	// Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	Type                       types.String `tfsdk:"type" helper:",default:Basic"`
	Folder                     types.String `tfsdk:"folder" helper:",required"`
	Account                    types.String `tfsdk:"account" helper:",required"`
	Application                types.String `tfsdk:"application" helper:",required"`
	MaxParallelSitPulls        types.Int64  `tfsdk:"max_parallel_sit_pulls" helper:"maxParallelSitPulls,default:0"`
	FlowAttrsMergeMode         types.String `tfsdk:"flow_attrs_merge_mode" helper:"flowAttrsMergeMode,enum:/preserve/overwrite/append,default:preserve"`
	FolderMonitorScheduleCheck types.String `tfsdk:"folder_monitor_schedule_check" helper:"folderMonitorScheduleCheck,emptyIsNull,default:"`
	FlowName                   types.String `tfsdk:"flow_name" helper:"flowName,,emptyIsNull,default:"`
	ScheduledFolderMonitor     types.String `tfsdk:"scheduled_folder_monitor" helper:"scheduledFolderMonitor,,emptyIsNull,default:"`
	SubscriptionEncryptMode    types.String `tfsdk:"subscription_encrypt_mode" helper:"subscriptionEncryptMode,default:DEFAULT"`
	FileRetentionPeriod        types.Int64  `tfsdk:"file_retention_period" helper:"fileRetentionPeriod,default:0"`
	FlowAttributes             types.Map    `tfsdk:"flow_attributes" helper:"flowAttributes,elementtype:string,optional"`

	Schedules []Schedule `tfsdk:"schedules" helper:"schedules,fold:type,optional"`

	TransferConfigurations []stTransferConfiguration `tfsdk:"transfer_configurations" helper:"transferConfigurations"`

	TemplateAdvancedExpression types.Bool `tfsdk:"template_advanced_expression" helper:"templateAdvancedExpression,default:false"`

	// PostTransmissionActions *stBasicSubscriptionPostTransmissionActions `tfsdk:"post_transmission_actions" helper:"postTransmissionActions,elementtype:stBasicSubscriptionPostTransmissionActions,optional"`
	PostTransmissionActions types.Object `tfsdk:"post_transmission_actions" helper:"postTransmissionActions,elementtype:stBasicSubscriptionPostTransmissionActions,default:"`
}

func NewSTBasicSubscriptionARModelResource() resource.Resource {
	return NewSTResource(&stBasicSubscriptionModel{}, "st_subscription_basic", "", "/api/v2.0/subscriptions", "/api/v2.0/subscriptions/{id}").AddDiscriminator("[type=Basic]")
}

func init() {
	registerResource(NewSTBasicSubscriptionARModelResource)
	tfhelper.RegisterType("stBasicSubscriptionPostTransmissionActions", &stBasicSubscriptionPostTransmissionActions{})
}
