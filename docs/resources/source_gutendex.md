---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_gutendex Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceGutendex Resource
---

# airbyte_source_gutendex (Resource)

SourceGutendex Resource

## Example Usage

```terraform
resource "airbyte_source_gutendex" "my_source_gutendex" {
  configuration = {
    author_year_end   = -500
    author_year_start = 2020
    copyright         = false
    languages         = "en"
    search            = "dickens%20great%20expect"
    sort              = "ascending"
    topic             = "children"
  }
  definition_id = "741fa85f-790e-4a62-807d-c3d6d966a992"
  name          = "...my_name..."
  secret_id     = "...my_secret_id..."
  workspace_id  = "abfa20fd-2d77-4fbf-ace1-fcc26e5c4f89"
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

- `author_year_end` (String) (Optional) Defines the maximum birth year of the authors. Books by authors born after the end year will not be returned. Supports both positive (CE) or negative (BCE) integer values
- `author_year_start` (String) (Optional) Defines the minimum birth year of the authors. Books by authors born prior to the start year will not be returned. Supports both positive (CE) or negative (BCE) integer values
- `copyright` (String) (Optional) Use this to find books with a certain copyright status - true for books with existing copyrights, false for books in the public domain in the USA, or null for books with no available copyright information.
- `languages` (String) (Optional) Use this to find books in any of a list of languages. They must be comma-separated, two-character language codes.
- `search` (String) (Optional) Use this to search author names and book titles with given words. They must be separated by a space (i.e. %20 in URL-encoded format) and are case-insensitive.
- `sort` (String) (Optional) Use this to sort books - ascending for Project Gutenberg ID numbers from lowest to highest, descending for IDs highest to lowest, or popular (the default) for most popular to least popular by number of downloads.
- `topic` (String) (Optional) Use this to search for a case-insensitive key-phrase in books' bookshelves or subjects.

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_gutendex.my_airbyte_source_gutendex ""
```
