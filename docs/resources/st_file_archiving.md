---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "xmft_st_file_archiving Resource - xmft"
subcategory: ""
description: |-
  
---

# xmft_st_file_archiving (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `archive_folder` (String)
- `delete_files_older_than` (Number) default:1
- `delete_files_older_than_unit` (String) default:days
- `encryption_certificate` (String) default:''
- `global_archiving_policy` (String) default:disabled
- `maximum_file_size_allowed_to_archive` (Number) default:0
- `name` (String) default:sentinel

### Read-Only

- `last_updated` (String)
