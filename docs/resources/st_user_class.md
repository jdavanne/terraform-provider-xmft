---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "xmft_st_user_class Resource - xmft"
subcategory: ""
description: |-
  
---

# xmft_st_user_class (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) User class name.

### Optional

- `address` (String) default:*, User class host.
- `enabled` (Boolean) default:true, Is user class currently enabled.
- `expression` (String) default:'', User class expression.
- `group` (String) default:*, Group of user class
- `order` (Number) default:1, The order in which user class is saved in database
- `user_name` (String) default:*, User class username.
- `user_type` (String) enum:/*/real/virtual, default:*, User class type.

### Read-Only

- `id` (String) The id of the user class.
- `last_updated` (String)
