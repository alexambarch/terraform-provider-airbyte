---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_vitally Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceVitally Resource
---

# airbyte_source_vitally (Resource)

SourceVitally Resource

## Example Usage

```terraform
resource "airbyte_source_vitally" "my_source_vitally" {
  configuration = {
    basic_auth_header = "...my_basic_auth_header..."
    domain            = "...my_domain..."
    secret_token      = "...my_secret_token..."
    status            = "active"
  }
  definition_id = "4ccbdee7-a61d-46e5-a8de-2826c945689b"
  name          = "...my_name..."
  secret_id     = "...my_secret_id..."
  workspace_id  = "7acaf484-b8ee-458f-8670-2faa57d0d719"
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

- `domain` (String) Provide only the domain part, like https://{your-domain}.rest.vitally.io/.  Keep empty if you don't have a subdomain.
- `secret_token` (String) sk_live_secret_token
- `status` (String) Status of the Vitally accounts. One of the following values; active, churned, activeOrChurned. must be one of ["active", "churned", "activeOrChurned"]

Optional:

- `basic_auth_header` (String, Sensitive) Basic Auth Header

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_vitally.my_airbyte_source_vitally ""
```
