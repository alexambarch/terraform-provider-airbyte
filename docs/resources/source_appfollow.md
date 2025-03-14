---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_appfollow Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceAppfollow Resource
---

# airbyte_source_appfollow (Resource)

SourceAppfollow Resource

## Example Usage

```terraform
resource "airbyte_source_appfollow" "my_source_appfollow" {
  configuration = {
    api_secret = "...my_api_secret..."
  }
  definition_id = "23ac9822-9f43-4e31-a31d-6a6109f207ae"
  name          = "...my_name..."
  secret_id     = "...my_secret_id..."
  workspace_id  = "fdaa89a3-6845-4ed8-abc8-d5d54ff37be9"
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

Optional:

- `api_secret` (String, Sensitive) API Key provided by Appfollow

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_appfollow.my_airbyte_source_appfollow ""
```
