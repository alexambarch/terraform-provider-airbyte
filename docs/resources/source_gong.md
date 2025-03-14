---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_gong Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceGong Resource
---

# airbyte_source_gong (Resource)

SourceGong Resource

## Example Usage

```terraform
resource "airbyte_source_gong" "my_source_gong" {
  configuration = {
    access_key        = "...my_access_key..."
    access_key_secret = "...my_access_key_secret..."
    start_date        = "2018-02-18T08:00:00Z"
  }
  definition_id = "0b702d4e-9a91-4582-b5bf-1ff51cd4b92d"
  name          = "...my_name..."
  secret_id     = "...my_secret_id..."
  workspace_id  = "66fc102e-2c63-4350-beb2-94c7945f2740"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String) Name of the source e.g. dev-mysql-instance.
- `workspace_id` (String)

### Optional

- `definition_id` (String) The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
- `secret_id` (String) Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.

### Read-Only

- `created_at` (Number)
- `source_id` (String)
- `source_type` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Required:

- `access_key` (String, Sensitive) Gong Access Key
- `access_key_secret` (String, Sensitive) Gong Access Key Secret

Optional:

- `start_date` (String) The date from which to list calls, in the ISO-8601 format; if not specified, the calls start with the earliest recorded call. For web-conference calls recorded by Gong, the date denotes its scheduled time, otherwise, it denotes its actual start time.

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_gong.my_airbyte_source_gong ""
```
