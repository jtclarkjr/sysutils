package ssl

import (
	"crypto/tls"
	"fmt"
	"time"
)

// CertInfo holds information about the SSL certificate.
type CertInfo struct {
	Domain     string
	ExpiryDate time.Time
	DaysLeft   int
	Issuer     string
}

// CheckSSL takes a domain and returns information about its SSL certificate.
func CheckSSL(domain string) (*CertInfo, error) {
	conn, err := tls.Dial("tcp", domain+":443", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s: %v", domain, err)
	}
	defer conn.Close()

	cert := conn.ConnectionState().PeerCertificates[0]
	daysLeft := int(time.Until(cert.NotAfter).Hours() / 24)

	return &CertInfo{
		Domain:     domain,
		ExpiryDate: cert.NotAfter,
		DaysLeft:   daysLeft,
		Issuer:     cert.Issuer.CommonName,
	}, nil
}
