---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_destination_duckdb Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  DestinationDuckdb DataSource
---

# airbyte_destination_duckdb (Data Source)

DestinationDuckdb DataSource

## Example Usage

```terraform
data "airbyte_destination_duckdb" "my_destination_duckdb" {
  destination_id = "...my_destination_id..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `destination_id` (String)

### Read-Only

- `configuration` (String) The values required to configure the destination. Parsed as JSON.
- `created_at` (Number)
- `definition_id` (String)
- `destination_type` (String)
- `name` (String)
- `workspace_id` (String)
