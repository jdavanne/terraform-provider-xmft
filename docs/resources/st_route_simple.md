---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "xmft_st_route_simple Resource - xmft"
subcategory: ""
description: |-
  
---

# xmft_st_route_simple (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)

### Optional

- `additional_attributes` (Map of String)
- `condition` (String)
- `condition_type` (String) default:ALWAYS
- `description` (String) default:''
- `failure_email_name` (String)
- `failure_email_notification` (Boolean)
- `failure_email_template` (String)
- `steps` (Attributes List) (see [below for nested schema](#nestedatt--steps))
- `success_email_name` (String)
- `success_email_notification` (Boolean)
- `success_email_template` (String)
- `triggering_email_name` (String)
- `triggering_email_notification` (Boolean)
- `triggering_email_template` (String)
- `type` (String) default:SIMPLE

### Read-Only

- `id` (String) The ID of this resource.
- `last_updated` (String)

<a id="nestedatt--steps"></a>
### Nested Schema for `steps`

Optional:

- `compress` (Attributes) (see [below for nested schema](#nestedatt--steps--compress))
- `decompress` (Attributes) (see [below for nested schema](#nestedatt--steps--decompress))
- `execute_route` (Attributes) (see [below for nested schema](#nestedatt--steps--execute_route))
- `pluggable` (Attributes) (see [below for nested schema](#nestedatt--steps--pluggable))
- `publish_to_account` (Attributes) (see [below for nested schema](#nestedatt--steps--publish_to_account))
- `pull_from_partner` (Attributes) (see [below for nested schema](#nestedatt--steps--pull_from_partner))
- `rename` (Attributes) (see [below for nested schema](#nestedatt--steps--rename))
- `send_to_partner` (Attributes) (see [below for nested schema](#nestedatt--steps--send_to_partner))

<a id="nestedatt--steps--compress"></a>
### Nested Schema for `steps.compress`

Optional:

- `action_on_step_failure` (String) enum:/FAIL/PROCEED, default:FAIL
- `action_on_step_success` (String) enum:/PROCEED/STOP, default:PROCEED
- `compression_level` (String) default:STORE
- `compression_type` (String) enum:/ZIP/JAR/TAR/GZIP, default:ZIP
- `condition` (String) default:''
- `condition_type` (String) enum:ALWAYS/EL, default:ALWAYS
- `file_filter_expression` (String) default:*
- `file_filter_expression_type` (String) default:GLOB
- `post_transformation_action_rename_as_expression` (String) default:''
- `single_archive_enabled` (Boolean) default:false
- `single_archive_name` (String) default:''
- `status` (String) default:ENABLED
- `type` (String) default:Compress
- `zip_password` (String) default:''

Read-Only:

- `id` (String)


<a id="nestedatt--steps--decompress"></a>
### Nested Schema for `steps.decompress`

Optional:

- `action_on_step_failure` (String) enum:/FAIL/PROCEED, default:FAIL
- `action_on_step_success` (String) enum:/PROCEED/STOP, default:PROCEED
- `condition` (String) default:''
- `condition_type` (String) enum:ALWAYS/EL, default:ALWAYS
- `file_filter_expression` (String) default:*
- `file_filter_expression_type` (String) default:GLOB
- `filename_collision_resolution_type` (String) enum:/RENAME_OLD/OVERWRITE/FAIL, default:OVERWRITE
- `post_transformation_action_rename_as_expression` (String) default:''
- `status` (String) default:ENABLED
- `type` (String) default:Decompress
- `zip_password` (String) default:''

Read-Only:

- `id` (String)


<a id="nestedatt--steps--execute_route"></a>
### Nested Schema for `steps.execute_route`

Required:

- `execute_route_id` (String)

Optional:

- `status` (String) default:ENABLED
- `type` (String) default:ExecuteRoute


<a id="nestedatt--steps--pluggable"></a>
### Nested Schema for `steps.pluggable`

Required:

- `custom_properties` (Map of String)

Optional:

- `action_on_step_failure` (String) enum:/FAIL/PROCEED, default:FAIL
- `action_on_step_success` (String) enum:/PROCEED/STOP, default:PROCEED
- `condition` (String) default:''
- `condition_type` (String) enum:ALWAYS/EL, default:ALWAYS
- `status` (String) default:ENABLED
- `type` (String) default:Pluggable

Read-Only:

- `id` (String)


<a id="nestedatt--steps--publish_to_account"></a>
### Nested Schema for `steps.publish_to_account`

Required:

- `target_account_expression` (String)

Optional:

- `action_on_step_failure` (String) enum:/FAIL/PROCEED, default:FAIL
- `action_on_step_success` (String) enum:/PROCEED/STOP, default:PROCEED
- `condition` (String) default:''
- `condition_type` (String) enum:ALWAYS/EL, default:ALWAYS
- `disable_auto_create_target_folder` (Boolean) default:false
- `file_filter_expression` (String) default:*
- `file_filter_expression_type` (String) enum:/GLOB/REGEXP/TEXT_FILES, default:GLOB
- `filename_collision_resolution_type` (String) enum:/RENAME_OLD/OVERWRITE/FAIL, default:OVERWRITE
- `post_routing_action_rename_expression` (String) default:''
- `post_routing_action_type` (String) default:d
- `publish_file_as` (String) default:''
- `status` (String) default:ENABLED
- `target_account_expression_type` (String) enum:/NAME/EXPRESSION, default:NAME
- `target_folder_expression` (String) default:/
- `target_folder_expression_type` (String) enum:/SIMPLE/EXPRESSION, default:SIMPLE
- `trigger_subscription` (Boolean) default:false
- `type` (String) default:Publish

Read-Only:

- `id` (String)


<a id="nestedatt--steps--pull_from_partner"></a>
### Nested Schema for `steps.pull_from_partner`

Required:

- `target_account_expression` (String)
- `transfer_site_expression` (String)

Optional:

- `action_on_step_failure` (String) enum:/FAIL/PROCEED, default:FAIL
- `action_on_step_success` (String) enum:/PROCEED/STOP, default:PROCEED
- `condition` (String) default:''
- `condition_type` (String) enum:ALWAYS/EL, default:ALWAYS
- `local_file_name_expression` (String) default:''
- `local_file_name_expression_type` (String) enum:/SIMPLE/EXPRESSION, default:SIMPLE
- `local_folder_path_expression` (String) default:''
- `local_folder_path_expression_type` (String) enum:/SIMPLE/EXPRESSION, default:SIMPLE
- `remote_file_name_expression` (String) default:''
- `remote_file_name_expression_type` (String) enum:/GLOB/REGEXP/TEXT_FILES, default:GLOB
- `remote_folder_path_expression` (String) default:''
- `remote_folder_path_expression_type` (String) enum:/SIMPLE/EXPRESSION, default:SIMPLE
- `status` (String) default:ENABLED
- `target_account_expression_type` (String) enum:/NAME/EXPRESSION, default:NAME
- `transfer_site_expression_type` (String) enum:/LIST, default:LIST
- `type` (String) default:PullFromPartner

Read-Only:

- `id` (String)


<a id="nestedatt--steps--rename"></a>
### Nested Schema for `steps.rename`

Required:

- `output_file_name` (String)

Optional:

- `action_on_step_failure` (String) enum:/FAIL/PROCEED, default:FAIL
- `action_on_step_success` (String) enum:/PROCEED/STOP, default:PROCEED
- `condition` (String) default:''
- `condition_type` (String) enum:ALWAYS/EL, default:ALWAYS
- `file_filter_expression` (String) default:*
- `file_filter_expression_type` (String) enum:/GLOB/REGEXP/TEXT_FILES, default:GLOB
- `status` (String) default:ENABLED
- `type` (String) default:Rename

Read-Only:

- `id` (String)


<a id="nestedatt--steps--send_to_partner"></a>
### Nested Schema for `steps.send_to_partner`

Optional:

- `action_on_step_failure` (String) enum:/FAIL/PROCEED, default:FAIL
- `action_on_step_success` (String) enum:/PROCEED/STOP, default:PROCEED
- `archive_policy_on_failure` (String)
- `archive_policy_on_success` (String)
- `condition` (String) default:''
- `condition_type` (String) enum:ALWAYS/EL, default:ALWAYS
- `data_encoding` (String) default:''
- `file_filter_expression` (String) default:*
- `file_filter_expression_type` (String) enum:/GLOB/REGEXP/TEXT_FILES, default:GLOB
- `file_label` (String) default:''
- `final_destination` (String) default:''
- `max_number_of_retries` (Number)
- `max_parallel_clients` (Number)
- `originator` (String) default:''
- `post_routing_action_rename_expression` (String)
- `post_routing_action_type` (String) enum:/DELETE/RENAME, default:''
- `record_format` (String) default:''
- `record_length` (String) default:''
- `route_file_as` (String) default:''
- `sleep_between_retries` (Number)
- `sleep_increment_between_retries` (Number)
- `status` (String) default:ENABLED
- `store_and_forward_mode` (String) default:''
- `target_account_expression` (String) default:''
- `target_account_expression_type` (String) enum:/NAME/EXPRESSION, default:''
- `transfer_profile_expression` (String) default:''
- `transfer_profile_expression_type` (String)
- `transfer_site_expression` (String) default:''
- `transfer_site_expression_type` (String) enum:/LIST/EXPRESSION_WILDCARD, default:LIST
- `trigger_file_content` (String)
- `trigger_file_for_each` (Boolean)
- `trigger_file_name` (String) default:''
- `trigger_target_account_expression` (String) default:''
- `trigger_target_account_expression_type` (String)
- `trigger_transfer_profile_expression` (String)
- `trigger_transfer_profile_expression_type` (String)
- `trigger_transfer_site_expression` (String) default:''
- `trigger_transfer_site_expression_type` (String) enum:/LIST/EXPRESSION_WILDCARD, default:''
- `trigger_upload_folder` (String)
- `type` (String) default:SendToPartner
- `upload_folder` (String) default:/
- `user_message` (String) default:''
- `virtual_filename` (String) default:''

Read-Only:

- `id` (String)
