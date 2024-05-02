---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "xmft_st_route_composite Resource - xmft"
subcategory: ""
description: |-
  
---

# xmft_st_route_composite (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)
- `route_template` (String)

### Optional

- `account` (String)
- `additional_attributes` (Map of String)
- `condition` (String)
- `condition_type` (String)
- `description` (String)
- `failure_email_name` (String)
- `failure_email_notification` (Boolean)
- `failure_email_template` (String)
- `managed_by_cg` (Boolean)
- `steps` (Attributes List) (see [below for nested schema](#nestedatt--steps))
- `subscriptions` (List of String)
- `success_email_name` (String)
- `success_email_notification` (Boolean)
- `success_email_template` (String)
- `triggering_email_name` (String)
- `triggering_email_notification` (Boolean)
- `triggering_email_template` (String)
- `type` (String)

### Read-Only

- `id` (String) The ID of this resource.
- `last_updated` (String)

<a id="nestedatt--steps"></a>
### Nested Schema for `steps`

Required:

- `execute_route` (String)

Optional:

- `status` (String)
- `type` (String)