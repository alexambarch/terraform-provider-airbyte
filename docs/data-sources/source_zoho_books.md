---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_zoho_books Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceZohoBooks DataSource
---

# airbyte_source_zoho_books (Data Source)

SourceZohoBooks DataSource

## Example Usage

```terraform
data "airbyte_source_zoho_books" "my_source_zohobooks" {
  source_id = "...my_source_id..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `source_id` (String)

### Read-Only

- `configuration` (String) The values required to configure the source. Parsed as JSON.
- `created_at` (Number)
- `definition_id` (String)
- `name` (String)
- `source_type` (String)
- `workspace_id` (String)
