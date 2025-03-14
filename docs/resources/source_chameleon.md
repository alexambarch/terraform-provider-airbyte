---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_chameleon Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceChameleon Resource
---

# airbyte_source_chameleon (Resource)

SourceChameleon Resource

## Example Usage

```terraform
resource "airbyte_source_chameleon" "my_source_chameleon" {
  configuration = {
    api_key    = "...my_api_key..."
    end_date   = "2022-10-06T08:02:50.190Z"
    filter     = "launcher"
    limit      = "...my_limit..."
    start_date = "2020-01-18T12:12:48.535Z"
  }
  definition_id = "5cbe19ca-55b1-46e7-9f3a-468d4b26f900"
  name          = "...my_name..."
  secret_id     = "...my_secret_id..."
  workspace_id  = "a1dd99e9-8adf-428a-ac07-dc3083920806"
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

- `api_key` (String, Sensitive)
- `start_date` (String)

Optional:

- `end_date` (String) End date for incremental sync
- `filter` (String) Filter for using in the `segments_experiences` stream. Default: "tour"; must be one of ["tour", "survey", "launcher"]
- `limit` (String) Max records per page limit. Default: "50"

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_chameleon.my_airbyte_source_chameleon ""
```
