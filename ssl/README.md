# SSL

## Example

```
package main

import (
  "fmt"
  "log"

  "github.com/jtclarkjr/sysutils/ssl"
)

func main() {
  domain := "example.com"
  certInfo, err := ssl.CheckSSL(domain)
  if err != nil {
    log.Fatalf("Error checking SSL certificate: %v", err)
  }

  fmt.Printf("Domain: %s\n", certInfo.Domain)
  fmt.Printf("Expiry Date: %s\n", certInfo.ExpiryDate)
  fmt.Printf("Days Left: %d\n", certInfo.DaysLeft)
  fmt.Printf("Issuer: %s\n", certInfo.Issuer)
}

```
