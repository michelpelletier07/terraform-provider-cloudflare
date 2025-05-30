---
page_title: "cloudflare_load_balancer_monitors Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_load_balancer_monitors (Data Source)



## Example Usage

```terraform
data "cloudflare_load_balancer_monitors" "example_load_balancer_monitors" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) Identifier

### Optional

- `max_items` (Number) Max items to fetch, default: 1000

### Read-Only

- `result` (Attributes List) The items returned by the data source (see [below for nested schema](#nestedatt--result))

<a id="nestedatt--result"></a>
### Nested Schema for `result`

Read-Only:

- `allow_insecure` (Boolean) Do not validate the certificate when monitor use HTTPS. This parameter is currently only valid for HTTP and HTTPS monitors.
- `consecutive_down` (Number) To be marked unhealthy the monitored origin must fail this healthcheck N consecutive times.
- `consecutive_up` (Number) To be marked healthy the monitored origin must pass this healthcheck N consecutive times.
- `created_on` (String)
- `description` (String) Object description.
- `expected_body` (String) A case-insensitive sub-string to look for in the response body. If this string is not found, the origin will be marked as unhealthy. This parameter is only valid for HTTP and HTTPS monitors.
- `expected_codes` (String) The expected HTTP response code or code range of the health check. This parameter is only valid for HTTP and HTTPS monitors.
- `follow_redirects` (Boolean) Follow redirects if returned by the origin. This parameter is only valid for HTTP and HTTPS monitors.
- `header` (Map of List of String) The HTTP request headers to send in the health check. It is recommended you set a Host header by default. The User-Agent header cannot be overridden. This parameter is only valid for HTTP and HTTPS monitors.
- `id` (String)
- `interval` (Number) The interval between each health check. Shorter intervals may improve failover time, but will increase load on the origins as we check from multiple locations.
- `method` (String) The method to use for the health check. This defaults to 'GET' for HTTP/HTTPS based checks and 'connection_established' for TCP based health checks.
- `modified_on` (String)
- `path` (String) The endpoint path you want to conduct a health check against. This parameter is only valid for HTTP and HTTPS monitors.
- `port` (Number) The port number to connect to for the health check. Required for TCP, UDP, and SMTP checks. HTTP and HTTPS checks should only define the port when using a non-standard port (HTTP: default 80, HTTPS: default 443).
- `probe_zone` (String) Assign this monitor to emulate the specified zone while probing. This parameter is only valid for HTTP and HTTPS monitors.
- `retries` (Number) The number of retries to attempt in case of a timeout before marking the origin as unhealthy. Retries are attempted immediately.
- `timeout` (Number) The timeout (in seconds) before marking the health check as failed.
- `type` (String) The protocol to use for the health check. Currently supported protocols are 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
Available values: "http", "https", "tcp", "udp_icmp", "icmp_ping", "smtp".


