---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_gitbook Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceGitbook Resource
---

# airbyte_source_gitbook (Resource)

SourceGitbook Resource

## Example Usage

```terraform
resource "airbyte_source_gitbook" "my_source_gitbook" {
  configuration = {
    access_token = "...my_access_token..."
    space_id     = "...my_space_id..."
  }
  definition_id = "bd853e05-630d-4e28-852c-3ff9e84d219d"
  name          = "...my_name..."
  secret_id     = "...my_secret_id..."
  workspace_id  = "a7f7be58-dede-4a51-8b06-e957318ad0a2"
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

- `access_token` (String, Sensitive) Personal access token for authenticating with the GitBook API. You can view and manage your access tokens in the Developer settings of your GitBook user account.
- `space_id` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_gitbook.my_airbyte_source_gitbook ""
```
