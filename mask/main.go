package mask

import (
	"crypto/sha256"
	"fmt"
	"net"
)

// MaskIP takes an IP address and returns a masked version of it.
func MaskIP(ip string) (string, error) {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return "", fmt.Errorf("invalid IP address: %s", ip)
	}
	hashedIP := sha256.Sum256([]byte(parsedIP.String()))
	return fmt.Sprintf("%x", hashedIP[:]), nil
}

// MaskDNS takes a DNS address and returns a masked version of it.
func MaskDNS(dns string) (string, error) {
	ips, err := net.LookupIP(dns)
	if err != nil {
		return "", fmt.Errorf("failed to resolve DNS: %s", dns)
	}

	maskedIPs := make([]string, len(ips))
	for i, ip := range ips {
		maskedIP, err := MaskIP(ip.String())
		if err != nil {
			return "", err
		}
		maskedIPs[i] = maskedIP
	}

	return fmt.Sprintf("%x", sha256.Sum256([]byte(dns))), nil
}
