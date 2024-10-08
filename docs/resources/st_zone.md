---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "xmft_st_zone Resource - xmft"
subcategory: ""
description: |-
  
---

# xmft_st_zone (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the DMZ zone. Zone name can not contain characters "\", "/", ";" and "'"

### Optional

- `description` (String) The description of DMZ Zone.
- `edges` (Attributes List) (see [below for nested schema](#nestedatt--edges))
- `is_default` (Boolean) Define whether the DMZ zone is set to default or not.
- `is_dns_resolution_enabled` (Boolean) Define whether DNS resolution on Edge is enabled or not (true if DNS resolution on Edge is enabled and false-otherwise).
- `public_url_prefix` (String) The public URL prefix for public access to the DMZ via HTTP(s).
- `sso_sp_entity_id` (String) The SSO Service Provider EntityId to identify to an Identity Provider.

<a id="nestedatt--edges"></a>
### Nested Schema for `edges`

Optional:

- `configuration_id` (String) The identifier of the machine
- `deployment_site` (String) The cluster deployment site of the edge
- `descriptor` (String) The unique identifier for current node (IP or hostname ..etc).
- `dynamic_node_ip_discovery_fqdn` (String) **Note:** Beta Feature - Do not use in Production. Check 'Beta.Dmz.DynamicNodeIpDiscovery.Enabled' configuration option.
**Note:** The property is available only for zones different than the 'Private' one and only on Backend.

Fqdn to be used for resolving multiple node addresses.
- `enabled_proxy` (Boolean) Check status of the proxy whether it is enabled/disabled (true if proxy is enabled, false otherwise)
- `ip_addresses` (Attributes List) (see [below for nested schema](#nestedatt--edges--ip_addresses))
- `is_auto_discoverable` (Boolean) **Note:** Beta Feature - Do not use in Production. Check 'Beta.Dmz.DynamicNodeIpDiscovery.Enabled' configuration option.
**Note:** The property is available only for zones different than the 'Private' one and only on Backend.

Check whether the zone type is auto-discoverable or static (true if zone is auto-discoverable, false - if static).
- `notes` (String) default:'', Notes for the DMZ edge
- `protocols` (Attributes List) (see [below for nested schema](#nestedatt--edges--protocols))
- `proxies` (Attributes List) (see [below for nested schema](#nestedatt--edges--proxies))
- `title` (String) default:'', Title of the DMZ edge. Edge title can not contain characters "\", "/", ";" and "'"

<a id="nestedatt--edges--ip_addresses"></a>
### Nested Schema for `edges.ip_addresses`

Required:

- `ip_address` (String) The IP address for DMZ edge.


<a id="nestedatt--edges--protocols"></a>
### Nested Schema for `edges.protocols`

Required:

- `port` (Number) The number of port for DMZ edge on the current protocol
- `streaming_protocol` (String) The streaming protocol for DMZ edge

Optional:

- `is_enabled` (Boolean) default:true, Define whether the protocol is enabled to be used from the edge
- `ssl_alias` (String) default:'', Certificate reference for DMZ edge on the current protocol.


<a id="nestedatt--edges--proxies"></a>
### Nested Schema for `edges.proxies`

Optional:

- `is_enabled` (Boolean) Define whether the proxy is enabled
- `is_use_password` (Boolean) Shows if the password should be preserved/overwritten or removed (true if the password should be preserved/overwritten with non empty one and false when the password should be removed)
- `password` (String) The password for the edge proxy. Password can be set only if 'isUsePassword' is set to 'true'. This property can not be get via the ST REST API.
- `port` (Number) The number of port for the proxi
- `proxy_protocol` (String) The proxy protocol
- `username` (String) The user name to authenticate to the proxy
