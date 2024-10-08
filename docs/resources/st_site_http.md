---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "xmft_st_site_http Resource - xmft"
subcategory: ""
description: |-
  
---

# xmft_st_site_http (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account` (String) The account for which the site is created
- `host` (String) Specify the remote site server host.
- `name` (String) The type of the site. Acts as a discriminator.
- `user_name` (String) Specify the site login username.

### Optional

- `access_level` (String) default:PRIVATE, The access level for this site.
- `additional_attributes` (Map of String) Additional attributes which are defined with "key": "value" pairs. Keys must start with "userVars." prefix, follow the pattern: [a-zA-Z0-9_.]+
and have length between 10 and 255 characters (including the prefix). Non prefixed part of key should not start with "userVars.", since it is
a reserved word. Both key and value cannot be blank.
- `alternative_addresses` (Attributes List) (see [below for nested schema](#nestedatt--alternative_addresses))
- `cipher_suites` (String) Cipher suites for http site.
- `client_certificate` (String) default:'', The client certificate ID used for mutual authentication.
- `default` (Boolean) default:false, Defines whether it is a default site. Only site from type PeSIT can be marked as 'Default'
- `dmz` (String) default:none, Specify the remote site network zone.
- `download_folder` (String) default:/, Specify the download folder.
- `download_folder_advanced_expression_enabled` (Boolean) default:false, When it is 'true' the download folder is evaluated using expression language.
- `download_pattern` (String) default:*, Specify the download pattern.
- `download_pattern_advanced_expression_enabled` (Boolean) default:false, When it is 'true' the download pattern is evaluated using expression language.
- `fips_mode` (Boolean) default:false, Specify if the FIPS Transfer Mode is enabled or disabled. This property can be set if HTTPS (issecure) is used.
- `is_secure` (Boolean) default:false, Specify if HTTPS is used or not.
- `max_concurrent_connection` (Number) default:0, The max concurrent connection of the site
- `password` (String) default:'', Specify the site login password encrypted in 'AES128'. This property should be set if 'usePassword' is 'true'.
- `port` (String) default:22, Specify the remote site server port.
- `post_transmission_actions` (Attributes) default:{} (see [below for nested schema](#nestedatt--post_transmission_actions))
- `protocol` (String) default:http, <nil>
- `protocols` (String) Enabled SSL protocols for http site.
- `request_mode` (String) enum:/GET/POST, default:GET, The value of the requestMode property
- `transfer_mode` (String) enum:/AUTO_DETECT/BINARY/ASCII, default:AUTO_DETECT, Specify the transfer mode. It can be A (Ascii), I (Binary). When it is not define the transfer mode is Auto detected.
- `transfer_type` (String) enum:/false/I/E, default:internal, The transfer type of the site. It can be unspecified (N), internal (I), partner(E)
- `type` (String) default:http, <nil>
- `upload_folder` (String) default:'', Specify the upload folder.
- `upload_folder_overridable` (Boolean) default:false, Defines if the upload folder can be modified by the Send To Partner routing step.
- `uri` (String) default:'', Specify the partner URL.
- `use_password` (Boolean) default:true, Specify if the login password should be set.
- `use_password_expr` (Boolean) default:false, Specify whether to have password expressions or not.
- `use_uri` (Boolean) default:false, Specify whether to define partner URL or not. If it's 'false', 'host' and 'port' should be specified. Otherwise 'url' should be specified.
- `verify_cert` (Boolean) default:false, Verify certificate for this site.

### Read-Only

- `id` (String) The id of the site.
- `last_updated` (String)

<a id="nestedatt--alternative_addresses"></a>
### Nested Schema for `alternative_addresses`

Required:

- `host` (String) The host/url of the alternative address.
- `port` (String) The port of the alternative address.
- `position` (Number) The position when alternate addresses.

Read-Only:

- `id` (String) The id of the site alternative address


<a id="nestedatt--post_transmission_actions"></a>
### Nested Schema for `post_transmission_actions`

Optional:

- `delete_on_perm_fail_in` (Boolean) default:false, Defines whether to delete the source file on failure after the transmission.
- `delete_on_perm_fail_out` (Boolean) default:false, Defines whether to delete the destination file on failure after the transmission.
- `delete_on_success_in` (Boolean) default:false, Defines whether to delete the source file on success after the transmission.
- `delete_on_temp_fail_out` (Boolean) default:false, Defines whether to delete the destination file on temporary failure after the transmission.
- `do_as_in` (String) default:'', Specify a value to receive the file with a different name. An expression language can be used to specify a file name e.g. ${stenv['target']}_${random()}.
- `do_as_out` (String) default:'', Specify a value to send the file with a different name. An expression language can be used to specify a file name e.g. ${stenv['target']}_${random()}.
