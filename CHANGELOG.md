
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
