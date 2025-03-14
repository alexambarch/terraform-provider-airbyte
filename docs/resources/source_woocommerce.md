---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_woocommerce Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceWoocommerce Resource
---

# airbyte_source_woocommerce (Resource)

SourceWoocommerce Resource

## Example Usage

```terraform
resource "airbyte_source_woocommerce" "my_source_woocommerce" {
  configuration = {
    api_key    = "...my_api_key..."
    api_secret = "...my_api_secret..."
    shop       = "...my_shop..."
    start_date = "2021-01-01"
  }
  definition_id = "5b87d40f-20e4-42b3-b267-1deb489c5b98"
  name          = "...my_name..."
  secret_id     = "...my_secret_id..."
  workspace_id  = "820ccea2-91c4-433d-bca2-ae97f9986afe"
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

- `api_key` (String, Sensitive) Customer Key for API in WooCommerce shop
- `api_secret` (String, Sensitive) Customer Secret for API in WooCommerce shop
- `shop` (String) The name of the store. For https://EXAMPLE.com, the shop name is 'EXAMPLE.com'.
- `start_date` (String) The date you would like to replicate data from. Format: YYYY-MM-DD

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_woocommerce.my_airbyte_source_woocommerce ""
```
