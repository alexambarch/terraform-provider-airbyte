---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_bamboo_hr Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceBambooHr DataSource
---

# airbyte_source_bamboo_hr (Data Source)

SourceBambooHr DataSource

## Example Usage

```terraform
data "airbyte_source_bamboo_hr" "my_source_bamboohr" {
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
