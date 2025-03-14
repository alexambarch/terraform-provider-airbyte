---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_kisi Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceKisi Resource
---

# airbyte_source_kisi (Resource)

SourceKisi Resource

## Example Usage

```terraform
resource "airbyte_source_kisi" "my_source_kisi" {
  configuration = {
    api_key = "...my_api_key..."
  }
  definition_id = "6a95d581-a8d4-4c54-8218-1af6e73e44ad"
  name          = "...my_name..."
  secret_id     = "...my_secret_id..."
  workspace_id  = "83b4e9a0-f0c6-4fff-9a55-474bf577a7b7"
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

- `api_key` (String) Your KISI API Key

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_kisi.my_airbyte_source_kisi ""
```
