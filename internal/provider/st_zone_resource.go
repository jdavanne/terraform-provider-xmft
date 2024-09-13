package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type stZoneModel struct {
	// Id              types.String   `tfsdk:"id" helper:",computed,state"`
	Name                   types.String `tfsdk:"name" helper:"name,state,required"`
	Description            types.String `tfsdk:"description" helper:"description,emptyIsNull,optional"`
	PublicURLPrefix        types.String `tfsdk:"public_url_prefix" helper:"publicURLPrefix,emptyIsNull,optional"`
	SsoSpEntityId          types.String `tfsdk:"sso_sp_entity_id" helper:"ssoSpEntityId,emptyIsNull,optional"`
	IsDnsResolutionEnabled types.Bool   `tfsdk:"is_dns_resolution_enabled" helper:"isDnsResolutionEnabled,optional"`
	IsDefault              types.Bool   `tfsdk:"is_default" helper:"isDefault,optional"`
	Edges                  []struct {
		Title                      types.String `tfsdk:"title" helper:"title,default:"`
		Notes                      types.String `tfsdk:"notes" helper:"notes,default:"`
		DeploymentSite             types.String `tfsdk:"deployment_site" helper:"deploymentSite,optional"`
		EnabledProxy               types.Bool   `tfsdk:"enabled_proxy" helper:"enabledProxy,optional"`
		IsAutoDiscoverable         types.Bool   `tfsdk:"is_auto_discoverable" helper:"isAutoDiscoverable,optional"`
		DynamicNodeIpDiscoveryFqdn types.String `tfsdk:"dynamic_node_ip_discovery_fqdn" helper:"dynamicNodeIpDiscoveryFqdn,optional"`
		IpAddresses                []struct {
			IpAddress types.String `tfsdk:"ip_address" helper:"ipAddress,required"`
		} `tfsdk:"ip_addresses" helper:"ipAddresses,optional"`
		ConfigurationId types.String `tfsdk:"configuration_id" helper:"configurationId,optional"`
		Descriptor      types.String `tfsdk:"descriptor" helper:"descriptor,optional"`
		Proxies         []struct {
			ProxyProtocol types.String `tfsdk:"proxy_protocol" helper:"proxyProtocol,optional"`
			Port          types.Int64  `tfsdk:"port" helper:"port,optional"`
			IsEnabled     types.Bool   `tfsdk:"is_enabled" helper:"isEnabled,optional"`
			UserName      types.String `tfsdk:"username" helper:"username,optional"`
			IsUsePassword types.Bool   `tfsdk:"is_use_password" helper:"isUsePassword,optional"`
			Password      types.String `tfsdk:"password" helper:"password,optional"`
		} `tfsdk:"proxies" helper:"proxies,optional"`
		Protocols []struct {
			StreamingProtocol types.String `tfsdk:"streaming_protocol" helper:"streamingProtocol,required"`
			Port              types.Int64  `tfsdk:"port" helper:"port,required"`
			IsEnabled         types.Bool   `tfsdk:"is_enabled" helper:"isEnabled,default:true"`
			SslAlias          types.String `tfsdk:"ssl_alias" helper:"sslAlias,emptyIsNull,default"`
		} `tfsdk:"protocols" helper:"protocols,optional"`
	} `tfsdk:"edges" helper:",optional"`
}

func NewSTZoneModelResource() resource.Resource {
	return NewSTResource(&stZoneModel{}, "st_zone", "", "/api/v2.0/zones", "/api/v2.0/zones/{name}")
}

func init() {
	registerResource(NewSTZoneModelResource)
}
