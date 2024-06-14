

## 0.0.5

FEATURES:
- *BREAKING* `st_route_simple.steps= [ send_to_partner = {} ]` 
- st_sentinel resource
- st_basic_application resource
- st_site_pesit resource
- st_business_unit resource
- add advanced routing steps types :
  - compress 
  - decompress
  - pluggable
  - rename
  - publish_to_account
  - pull_from_partner 
  - send_to_partner

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
