---
page_title: "cloudflare_account_dns_settings_internal_views Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_account_dns_settings_internal_views (Data Source)



## Example Usage

```terraform
data "cloudflare_account_dns_settings_internal_views" "example_account_dns_settings_internal_views" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  name = {
    contains = "view"
    endswith = "ew"
    exact = "my view"
    startswith = "my"
  }
  order = "name"
  zone_id = "ae29bea30e2e427ba9cd8d78b628177b"
  zone_name = "www.example.com"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) Identifier

### Optional

- `direction` (String) Direction to order DNS views in.
Available values: "asc", "desc".
- `match` (String) Whether to match all search requirements or at least one (any). If set to `all`, acts like a logical AND between filters. If set to `any`, acts like a logical OR instead.
Available values: "any", "all".
- `max_items` (Number) Max items to fetch, default: 1000
- `name` (Attributes) (see [below for nested schema](#nestedatt--name))
- `order` (String) Field to order DNS views by.
Available values: "name", "created_on", "modified_on".
- `zone_id` (String) A zone ID that exists in the zones list for the view.
- `zone_name` (String) A zone name that exists in the zones list for the view.

### Read-Only

- `result` (Attributes List) The items returned by the data source (see [below for nested schema](#nestedatt--result))

<a id="nestedatt--name"></a>
### Nested Schema for `name`

Optional:

- `contains` (String) Substring of the DNS view name.
- `endswith` (String) Suffix of the DNS view name.
- `exact` (String) Exact value of the DNS view name.
- `startswith` (String) Prefix of the DNS view name.


<a id="nestedatt--result"></a>
### Nested Schema for `result`

Read-Only:

- `created_time` (String) When the view was created.
- `id` (String) Identifier
- `modified_time` (String) When the view was last modified.
- `name` (String) The name of the view.
- `zones` (List of String) The list of zones linked to this view.


