# Mask

## Example

```
package main

import (
  "fmt"
  "log"

  "github.com/jtclarkjr/sysutils/mask"
)

func main() {
  ip := "192.168.1.1"
  maskedIP, err := mask.MaskIP(ip)
  if err != nil {
    log.Fatalf("Error masking IP: %v", err)
  }
  fmt.Printf("Original IP: %s, Masked IP: %s\n", ip, maskedIP)

  dns := "example.com"
  maskedDNS, err := mask.MaskDNS(dns)
  if err != nil {
    log.Fatalf("Error masking DNS: %v", err)
  }
  fmt.Printf("Original DNS: %s, Masked DNS: %s\n", dns, maskedDNS)
}
```
