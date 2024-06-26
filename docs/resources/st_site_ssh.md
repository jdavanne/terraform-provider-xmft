---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "xmft_st_site_ssh Resource - xmft"
subcategory: ""
description: |-
  
---

# xmft_st_site_ssh (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account` (String)
- `host` (String)
- `name` (String)
- `user_name` (String)

### Optional

- `access_level` (String) default:PRIVATE
- `additional_attributes` (Map of String)
- `allowed_macs` (String)
- `alternative_addresses` (Attributes List) (see [below for nested schema](#nestedatt--alternative_addresses))
- `block_size` (Number) default:32768
- `buffer_size` (Number) default:32768
- `cipher_suites` (String)
- `client_certificate` (String) default:''
- `default` (Boolean) default:false
- `dmz` (String) default:none
- `download_folder` (String) default:''
- `download_folder_advanced_expression_enabled` (Boolean) default:false
- `download_pattern` (String) default:*
- `download_pattern_advanced_expression_enabled` (Boolean) default:false
- `download_pattern_type` (String) default:glob
- `finger_print` (String) default:''
- `fips_mode` (Boolean) default:false
- `key_exchange_algorithms` (String)
- `max_concurrent_connection` (Number) default:0
- `password` (String) default:''
- `port` (String) default:22
- `protocol` (String) default:ssh
- `protocols` (String) default:''
- `public_keys` (String)
- `socket_buffer_size` (Number) default:65536
- `socket_send_buffer_size` (Number) default:65536
- `socket_timeout` (Number) default:300
- `tcp_no_delay` (Boolean) default:true
- `transfer_mode` (String) default:AUTO_DETECT
- `transfer_type` (String) default:internal
- `type` (String) default:ssh
- `update_permissions_with_chmod_command` (String) default:''
- `upload_folder` (String) default:''
- `upload_folder_overridable` (Boolean) default:false
- `upload_permissions` (String) default:0644
- `use_password` (Boolean) default:true
- `use_password_expr` (Boolean) default:false
- `verify_finger` (Boolean) default:false

### Read-Only

- `id` (String) The ID of this resource.
- `last_updated` (String)

<a id="nestedatt--alternative_addresses"></a>
### Nested Schema for `alternative_addresses`

Required:

- `host` (String)
- `port` (String)
- `position` (Number)
