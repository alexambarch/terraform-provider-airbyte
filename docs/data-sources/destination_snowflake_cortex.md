---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_destination_snowflake_cortex Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  DestinationSnowflakeCortex DataSource
---

# airbyte_destination_snowflake_cortex (Data Source)

DestinationSnowflakeCortex DataSource

## Example Usage

```terraform
data "airbyte_destination_snowflake_cortex" "my_destination_snowflakecortex" {
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
