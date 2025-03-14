---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_aircall Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceAircall Resource
---

# airbyte_source_aircall (Resource)

SourceAircall Resource

## Example Usage

```terraform
resource "airbyte_source_aircall" "my_source_aircall" {
  configuration = {
    api_id     = "...my_api_id..."
    api_token  = "...my_api_token..."
    start_date = "2022-03-01T00:00:00.000Z"
  }
  definition_id = "f2d6765e-6635-44d4-ae8c-8380f4c57950"
  name          = "...my_name..."
  secret_id     = "...my_secret_id..."
  workspace_id  = "e53fe6b3-92c0-44af-8136-cbb3bc59d66b"
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

- `api_id` (String, Sensitive) App ID found at settings https://dashboard.aircall.io/integrations/api-keys
- `api_token` (String, Sensitive) App token found at settings (Ref- https://dashboard.aircall.io/integrations/api-keys)
- `start_date` (String) Date time filter for incremental filter, Specify which date to extract from.

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_aircall.my_airbyte_source_aircall ""
```
