---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_auth0 Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceAuth0 Resource
---

# airbyte_source_auth0 (Resource)

SourceAuth0 Resource

## Example Usage

```terraform
resource "airbyte_source_auth0" "my_source_auth0" {
  configuration = {
    base_url = "https://dev-yourOrg.us.auth0.com/"
    credentials = {
      o_auth2_access_token = {
        access_token = "...my_access_token..."
      }
      o_auth2_confidential_application = {
        audience      = "https://dev-yourOrg.us.auth0.com/api/v2/"
        client_id     = "Client_ID"
        client_secret = "Client_Secret"
      }
    }
    start_date = "2023-08-05T00:43:59.244Z"
  }
  definition_id = "c81522f4-2b1b-4d41-9891-5186609bc62d"
  name          = "...my_name..."
  secret_id     = "...my_secret_id..."
  workspace_id  = "0f188999-803b-497f-bcdb-fe3868611adc"
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

- `base_url` (String) The Authentication API is served over HTTPS. All URLs referenced in the documentation have the following base `https://YOUR_DOMAIN`
- `credentials` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials))

Optional:

- `start_date` (String, Sensitive) UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated. Default: "2023-08-05T00:43:59.244Z"

<a id="nestedatt--configuration--credentials"></a>
### Nested Schema for `configuration.credentials`

Optional:

- `o_auth2_access_token` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--o_auth2_access_token))
- `o_auth2_confidential_application` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--o_auth2_confidential_application))

<a id="nestedatt--configuration--credentials--o_auth2_access_token"></a>
### Nested Schema for `configuration.credentials.o_auth2_access_token`

Required:

- `access_token` (String, Sensitive) Also called <a href="https://auth0.com/docs/secure/tokens/access-tokens/get-management-api-access-tokens-for-testing">API Access Token </a> The access token used to call the Auth0 Management API Token. It's a JWT that contains specific grant permissions knowns as scopes.


<a id="nestedatt--configuration--credentials--o_auth2_confidential_application"></a>
### Nested Schema for `configuration.credentials.o_auth2_confidential_application`

Required:

- `audience` (String) The audience for the token, which is your API. You can find this in the Identifier field on your  <a href="https://manage.auth0.com/#/apis">API's settings tab</a>
- `client_id` (String) Your application's Client ID. You can find this value on the <a href="https://manage.auth0.com/#/applications">application's settings tab</a> after you login the admin portal.
- `client_secret` (String, Sensitive) Your application's Client Secret. You can find this value on the <a href="https://manage.auth0.com/#/applications">application's settings tab</a> after you login the admin portal.

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_auth0.my_airbyte_source_auth0 ""
```
