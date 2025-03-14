---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_chargedesk Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceChargedesk Resource
---

# airbyte_source_chargedesk (Resource)

SourceChargedesk Resource

## Example Usage

```terraform
resource "airbyte_source_chargedesk" "my_source_chargedesk" {
  configuration = {
    password   = "...my_password..."
    start_date = 2
    username   = "...my_username..."
  }
  definition_id = "5b19adc8-7cc9-4582-909b-da4c3635bcf7"
  name          = "...my_name..."
  secret_id     = "...my_secret_id..."
  workspace_id  = "72e1cba6-e614-4114-a01d-10c04d110961"
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

- `username` (String)

Optional:

- `password` (String, Sensitive)
- `start_date` (Number) Date from when the sync should start in epoch Unix timestamp

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_chargedesk.my_airbyte_source_chargedesk ""
```
