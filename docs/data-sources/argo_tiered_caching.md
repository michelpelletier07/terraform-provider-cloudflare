---
page_title: "cloudflare_argo_tiered_caching Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_argo_tiered_caching (Data Source)



## Example Usage

```terraform
data "cloudflare_argo_tiered_caching" "example_argo_tiered_caching" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `zone_id` (String) Identifier

### Read-Only

- `editable` (Boolean) Whether the setting is editable
- `id` (String) ID of the zone setting.
Available values: "tiered_caching".
- `modified_on` (String) Last time this setting was modified.
- `value` (String) The value of the feature
Available values: "on", "off".


