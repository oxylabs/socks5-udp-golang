# Oxylabs SOCKS5 UDP examples

[![Oxylabs promo code](https://raw.githubusercontent.com/oxylabs/product-integrations/refs/heads/master/Affiliate-Universal-1090x275.png)](https://oxylabs.go2cloud.org/aff_c?offer_id=7&aff_id=877&url_id=112)

This an example of using SOCKS5 UDP protocol with Oxylabs and Go.

You can modify `main.go` constants for default values to run without params 

```go
const defaultProxyURL = "socks.pr.oxylabs.io:7777"
const defaultProxyUsername = "customer-USERNAME"
const defaultProxyPassword = "PASSWORD"
const defaultTarget = "TARGET:PORT"
const defaultContent = "Hello UDP"
```

or build and run it as is by providing auth and target paramters:

```
  -c string
        content to send to target (default "Hello UDP")
  -p string
        proxy password (default "PASSWORD")
  -proxy string
        proxy URL (hostname:port (default "socks.pr.oxylabs.io:7777")
  -t string
        target (hostname:port) (default "TARGET:PORT")
  -u string
        proxy username (default "customer-USERNAME")
```
