---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_coassemble Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceCoassemble Resource
---

# airbyte_source_coassemble (Resource)

SourceCoassemble Resource

## Example Usage

```terraform
resource "airbyte_source_coassemble" "my_source_coassemble" {
  configuration = {
    user_id    = "...my_user_id..."
    user_token = "...my_user_token..."
  }
  definition_id = "44258de0-0776-4cf5-abd4-d67cfa889dbc"
  name          = "...my_name..."
  secret_id     = "...my_secret_id..."
  workspace_id  = "06f5b03f-2fe5-44ce-9f09-f2c22df1af8a"
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

- `user_id` (String, Sensitive)
- `user_token` (String, Sensitive)

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_coassemble.my_airbyte_source_coassemble ""
```
