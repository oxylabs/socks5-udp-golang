# Oxylabs SOCKS5 UDP examples

[![Oxylabs promo code](https://raw.githubusercontent.com/oxylabs/product-integrations/refs/heads/master/Affiliate-Universal-1090x275.png)](https://oxylabs.io/pages/gitoxy?utm_source=877&utm_medium=affiliate&groupid=877&utm_content=socks5-udp-golang-github&transaction_id=102f49063ab94276ae8f116d224b67)

[![](https://dcbadge.limes.pink/api/server/Pds3gBmKMH?style=for-the-badge&theme=discord)](https://discord.gg/Pds3gBmKMH) [![YouTube](https://img.shields.io/badge/YouTube-Oxylabs-red?style=for-the-badge&logo=youtube&logoColor=white)](https://www.youtube.com/@oxylabs)

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
