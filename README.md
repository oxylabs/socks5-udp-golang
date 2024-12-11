# Oxylabs SOCKS5 UDP examples

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
