# Oxylabs SOCKS5 UDP examples

[![Oxylabs promo code](https://raw.githubusercontent.com/oxylabs/product-integrations/refs/heads/master/Affiliate-Universal-1090x275.png)](https://oxylabs.io/pages/gitoxy?utm_source=github&utm_medium=referral&utm_content=socks5-udp-golang&utm_campaign=banner&groupid=877)

<a href="https://oxylabs.io/?utm_source=github&utm_medium=referral&utm_content=socks5-udp-golang&utm_campaign=banner&groupid=877"><img src="https://img.shields.io/badge/🚀_Get_Started-FF6B35?style=for-the-badge" alt="Get started"></a>
<a href="https://developers.oxylabs.io"><img src="https://img.shields.io/badge/📚_Documentation-4A90E2?style=for-the-badge" alt="Documentation"></a>
<a href="https://dashboard.oxylabs.io"><img src="https://img.shields.io/badge/📊_Dashboard-2EA44F?style=for-the-badge" alt="Dashboard"></a>
<a href="https://discord.gg/Pds3gBmKMH"><img src="https://img.shields.io/badge/💬_Discord-5865F2?style=for-the-badge" alt="Discord"></a>
<a href="https://www.youtube.com/@oxylabs"><img src="https://img.shields.io/badge/▶️_YouTube-FF0000?style=for-the-badge" alt="YouTube"></a>

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
