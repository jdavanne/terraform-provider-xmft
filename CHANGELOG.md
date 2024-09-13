
# 0.0.10

 (!) means barely tested
FEATURES:
- full st_account_* support:
  - add st_account_service resource (!)
  - add st_account_template resource (!)
- full st_application_* support:
  - add st_application_account_ttl resource (!)
  - add st_application_archive_maintenance resource (!)
  - add st_application_basic resource (!)
  - add st_application_file_maintenance resource (!)
  - add st_application_human_to_system resource (!)
  - add st_application_log_entry_maintenance resource (!)
  - add st_application_login_threshold_maintenance resource (!)
  - add st_application_mbft resource (!)
  - add st_application_package_retention_maintenance resource (!)
  - add st_application_sentinel_link_data resource (!)
  - add st_application_shared_folder resource (!)
  - add st_application_site_mailbox resource (!)
  - add st_application_standard_router resource (!)
  - add st_application_synchrony_transfer resource (!)
  - add st_application_transfer_log_maintenance resource (!)
  - add st_application_unlicenced_account_maintenance resource (!)
- full st_subscription_* support:
  - add st_subscription_basic resource
  - add st_subscription_human_to_system resource (!)
  - add st_subscription_mbft resource (!)
  - add st_subscription_shared_folder resource (!)
  - add st_subscription_site_mailbox resource (!)
  - add st_subscription_standard_router resource (!)
  - add st_subscription_synchrony_transfer resource (!)
- full st_site_* support
  - add st_site_adhoc resource (!)
  - add st_site_as2 resource (!)
  - add st_site_connect_direct resource (!)
  - add st_site_ftp resource
  - add st_site_http resource
  - add st_site_synchrony_transfer resource (!)
- misc
  - add st_address_book resource
  - add st_zone (!)
  - add st_icap_server (!)
  - add st_ldap_domain (!)
  - add st_login_restriction_policy (!)

# 0.0.9

FEATURES:

- add st_conf_option resource (single value only)

# 0.0.8

FEATURES:
- add advanced routing steps types (still missing line_ending, line_padding):
  - line_folding
  - line_truncating
  - pgp_decryption
  - pgp_encryption
  - pluggable
  - characters_replace
  - encoding_conversion
  - external_script
- add st_site_custom resource

CAVEATS:
- line_ending and line_padding are not yet supported
- st_site_custom.custom_properties uses CamelCase instead of snake_case
- st_site_custom.custom_properties is write only, no change with be observed from server, however st_site_custom.custom_properties with get the latest field

# 0.0.7

FEATURES:
- support arm64 architecture
- add st_site_folder_monitoring resource
- add st_admin_role resource
- add st_user_class resource
- add st_admin resource
- add st_file_archiving resource

## 0.0.6

FEATURES:
- st_certificate support
- add better description in documentation 

## 0.0.5

FEATURES:
- *BREAKING* `st_route_simple.steps= [ { send_to_partner = {} } ]` 
- *BREAKING* `st_route_composite.steps= [ { execute_route_id = xmft_st_route_simple.xxxx.id } ]` 
- st_sentinel resource
- st_basic_application resource
- st_site_pesit resource
- st_transfer_profile
- st_business_unit resource
- add advanced routing steps types :
  - compress 
  - decompress
  - pluggable
  - rename
  - publish_to_account
  - pull_from_partner 
  - send_to_partner
  - execute_route

## 0.0.4

FEATURES:
- basic ST resources : 
  - st_account resource
  - st_advanced_routing application resource
  - st_route_template resource
  - st_route_composite resource (AKA route package)
  - st_route_simple resource (with basic send-to-partner steps)
  - st_site_ssh resource
- basic CFT resources : 
  - cft_cftpart resource
  - cft_cftsend resource
  - cft_cftrecv resource
  - cft_cfttcp (deprecated - see cfpart) resource
