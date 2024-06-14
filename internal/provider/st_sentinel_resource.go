package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

/*
{
  "enabled": false,
  "host": "string",
  "port": 65535,
  "useSecureConnection": false,
  "shouldVerifyCert": false,
  "fipsEnabled": false,
  "heartbeatEnabled": false,
  "heartbeatDelay": 10,
  "heartbeatTimeUnit": "seconds",
  "eventStates": {},
  "mappingRules": {},
  "overflowFilePath": "string",
  "overflowFileSize": 1,
  "overflowFilePolicy": "ignore",
  "overflowFileThreshold": 94,
  "shouldPersistLinkData": false
}
*/

type stConfigurationSentinelModel struct {
	// Id          types.String `tfsdk:"id" helper:",computed,state"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`
	Name        types.String `tfsdk:"name" helper:",default:sentinel,noread,nowrite"`

	Enabled               types.Bool   `tfsdk:"enabled"`
	Host                  types.String `tfsdk:"host"`
	Port                  types.Int64  `tfsdk:"port"`
	UseSecureConnection   types.Bool   `tfsdk:"use_secure_connection" helper:"useSecureConnection,default:false"`
	ShouldVerifyCert      types.Bool   `tfsdk:"should_verify_cert" helper:"shouldVerifyCert,default:false"`
	FipsEnabled           types.Bool   `tfsdk:"fips_enabled" helper:"fipsEnabled,default:false"`
	HeartbeatEnabled      types.Bool   `tfsdk:"heartbeat_enabled" helper:"heartbeatEnabled,default:false"`
	HeartbeatDelay        types.Int64  `tfsdk:"heartbeat_delay" helper:"heartbeatDelay,default:10"`
	HeartbeatTimeUnit     types.String `tfsdk:"heartbeat_time_unit" helper:"heartbeatTimeUnit,default:seconds"`
	EventStates           types.Map    `tfsdk:"event_states" helper:"eventStates,elementtype:string,noread"`
	MappingRules          types.Map    `tfsdk:"mapping_rules" helper:"mappingRules,elementtype:string"`
	OverflowFilePath      types.String `tfsdk:"overflow_file_path" helper:"overflowFilePath"`
	OverflowFileSize      types.Int64  `tfsdk:"overflow_file_size" helper:"overflowFileSize,default:1"`
	OverflowFilePolicy    types.String `tfsdk:"overflow_file_policy" helper:"overflowFilePolicy,default:ignore"`
	OverflowFileThreshold types.Int64  `tfsdk:"overflow_file_threshold" helper:"overflowFileThreshold,default:94"`
	ShouldPersistLinkData types.Bool   `tfsdk:"should_persist_link_data" helper:"shouldPersistLinkData,default:false"`
}

func NewSTSentinelModelResource() resource.Resource {
	return NewSTResource(&stConfigurationSentinelModel{}, "st_sentinel", "", "/api/v2.0/configurations/sentinel", "/api/v2.0/configurations/sentinel").OnlyReplace()
}

func init() {
	registerResource(NewSTSentinelModelResource)
}
