---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_coin_api Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceCoinAPI Resource
---

# airbyte_source_coin_api (Resource)

SourceCoinAPI Resource

## Example Usage

```terraform
resource "airbyte_source_coin_api" "my_source_coinapi" {
  configuration = {
    api_key     = "...my_api_key..."
    end_date    = "2019-01-01T00:00:00"
    environment = "sandbox"
    limit       = 39021
    period      = "5SEC"
    start_date  = "2019-01-01T00:00:00"
    symbol_id   = "...my_symbol_id..."
  }
  definition_id = "bddfa872-3b12-4feb-a665-3aa5aa88bcb2"
  name          = "...my_name..."
  secret_id     = "...my_secret_id..."
  workspace_id  = "0929de6e-00e3-4d8c-81a1-849a2ab68cdb"
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

- `api_key` (String, Sensitive) API Key
- `period` (String) The period to use. See the documentation for a list. https://docs.coinapi.io/#list-all-periods-get
- `start_date` (String) The start date in ISO 8601 format.
- `symbol_id` (String) The symbol ID to use. See the documentation for a list.
https://docs.coinapi.io/#list-all-symbols-get

Optional:

- `end_date` (String) The end date in ISO 8601 format. If not supplied, data will be returned
from the start date to the current time, or when the count of result
elements reaches its limit.
- `environment` (String) The environment to use. Either sandbox or production. Default: "sandbox"; must be one of ["sandbox", "production"]
- `limit` (Number) The maximum number of elements to return. If not supplied, the default
is 100. For numbers larger than 100, each 100 items is counted as one
request for pricing purposes. Maximum value is 100000.
Default: 100

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_coin_api.my_airbyte_source_coin_api ""
```
