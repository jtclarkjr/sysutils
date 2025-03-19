# CACERT

## Example

```go
package main

import (
  "crypto/x509"
  "encoding/pem"
  "fmt"
  "log"

  "github.com/jtclarkjr/sysutils/cacert"
)

func main() {
  cm, err := cacert.NewCertManager("certs")
  if err != nil {
      log.Fatalf("Error creating CertManager: %v", err)
  }

  // Sample CA certificate (replace with a real certificate for actual use)
  sampleCert := `-----BEGIN CERTIFICATE-----
  MIIDXTCCAkWgAwIBAgIJAL7Ab7QKRxnZMA0GCSqGSIb3DQEBCwUAMEUxCzAJBgNV
  BAYTAkFVMQswCQYDVQQIDAJOU1cxFTATBgNVBAoMDEV4YW1wbGUgSW5jLjEQMA4G
  A1UEAwwHZXhhbXBsZTAeFw0xNzA5MTMxNzUyMDBaFw0yNzA5MTE3NTIwMDBaMEUx
  CzAJBgNVBAYTAkFVMQswCQYDVQQIDAJOU1cxFTATBgNVBAoMDEV4YW1wbGUgSW5j
  LjEQMA4GA1UEAwwHZXhhbXBsZTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoC
  ggEBAJZl27ovJxrChxy1tIRvE7ivZ/N5pt7FdFGu16BaQPV9cd9bD/tLRLXuUs7y
  wOPfovcttKlGnOBt8jWkPmNE1A+cBLJdVrTTrp0L8lXvFAhxT6RGdKRf4E8pP8lK
  +XACD2Ulv91sp3n1K6V9nYc9MWvNH9WlRdAKR3EZh4fFTVL6B4KgT8Cx0MD9pKtO
  BzKx2P8Naa5Tqv4vzmV2AalNdPU1zObg6Mne1ePx8dJ2KxykkupE+chVqWKh0aWm
  71ErgWkN/b9Ude3F+tvDG3/LgW/B7z5lN6zKhbL9HDGThM4jOGmWycANkQhI6mhA
  pVnWfFyj/a6O9CDEEy81VJbVvMsCAwEAAaNQME4wHQYDVR0OBBYEFJImnPlwS0lP
  hPH07sl39nmZz48pMB8GA1UdIwQYMBaAFJImnPlwS0lPhPH07sl39nmZz48pMAwG
  A1UdEwQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEBAKLwSbV5C8P0LqFRs5mAN6Bd
  7lvUBpUz6zkNCG+xklZVov3/CiO6YboYwKDKKpzxFoGGA5UmKcuzL45EwzQe/rsS
  fgC6ILztTce6L7hT8D3OBWjK2ZQ6R/EKrhOEv/OBBQb4E6PbVzOveAXDRl0BA9ik
  RzhjBF8K4Xeog2ldXFFPhA2orZ/qGwLt5V1pRfjgXMJ5vghs6y7nIDo+XEqU1fKy
  ZHqLRySOkHdRgrD+M8EapHkT2QJlm2E8es1FW5BzErq1TC3EvwHvBz/5koySn3OE
  dsz3+bQeGwY7wD3n3Z18Ej7Zxy6T8MY1l1FQ3H+kdlR9B+xggWAhzXP4uZ4=
  -----END CERTIFICATE-----`

  block, _ := pem.Decode([]byte(sampleCert))
  if block == nil {
    log.Fatal("Failed to decode PEM block")
  }
  cert, err := x509.ParseCertificate(block.Bytes)
  if err != nil {
    log.Fatalf("Failed to parse certificate: %v", err)
  }

  // Add the certificate
  err = cm.AddCert([]byte(sampleCert), "example_cert.pem")
  if err != nil {
    log.Fatalf("Error adding certificate: %v", err)
  }
  fmt.Println("Certificate added.")

  // List certificates
  certs, err := cm.ListCerts()
  if err != nil {
    log.Fatalf("Error listing certificates: %v", err)
  }
  for _, c := range certs {
    fmt.Printf("Certificate: %s, Issuer: %s, Expiry: %s\n", c.Subject.CommonName, c.Issuer.CommonName, c.NotAfter)
  }

  // Remove the certificate
  err = cm.RemoveCert("example_cert.pem")
  if err != nil {
    log.Fatalf("Error removing certificate: %v", err)
  }
  fmt.Println("Certificate removed.")
}
```
