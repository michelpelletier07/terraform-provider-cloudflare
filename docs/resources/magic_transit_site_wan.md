---
page_title: "cloudflare_magic_transit_site_wan Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_magic_transit_site_wan (Resource)



## Example Usage

```terraform
resource "cloudflare_magic_transit_site_wan" "example_magic_transit_site_wan" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  site_id = "023e105f4ecef8ad9ca31a8372d0c353"
  physport = 1
  name = "name"
  priority = 0
  static_addressing = {
    address = "192.0.2.0/24"
    gateway_address = "192.0.2.1"
    secondary_address = "192.0.2.0/24"
  }
  vlan_tag = 42
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) Identifier
- `physport` (Number)
- `site_id` (String) Identifier

### Optional

- `name` (String)
- `priority` (Number)
- `static_addressing` (Attributes) (optional) if omitted, use DHCP. Submit secondary_address when site is in high availability mode. (see [below for nested schema](#nestedatt--static_addressing))
- `vlan_tag` (Number) VLAN ID. Use zero for untagged.

### Read-Only

- `health_check_rate` (String) Magic WAN health check rate for tunnels created on this link. The default value is `mid`.
Available values: "low", "mid", "high".
- `id` (String) Identifier

<a id="nestedatt--static_addressing"></a>
### Nested Schema for `static_addressing`

Required:

- `address` (String) A valid CIDR notation representing an IP range.
- `gateway_address` (String) A valid IPv4 address.

Optional:

- `secondary_address` (String) A valid CIDR notation representing an IP range.

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_magic_transit_site_wan.example '<account_id>/<site_id>/<wan_id>'
```
