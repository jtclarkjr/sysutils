package cacert

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"
)

// CertManager manages CA certificates in a local directory.
type CertManager struct {
	dir string
}

// NewCertManager creates a new CertManager.
func NewCertManager(dir string) (*CertManager, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create directory: %v", err)
		}
	}
	return &CertManager{dir: dir}, nil
}

// AddCert adds a new CA certificate to the store.
func (cm *CertManager) AddCert(certPEM []byte, filename string) error {
	path := filepath.Join(cm.dir, filename)
	return os.WriteFile(path, certPEM, 0644)
}

// RemoveCert removes a CA certificate from the store.
func (cm *CertManager) RemoveCert(filename string) error {
	path := filepath.Join(cm.dir, filename)
	return os.Remove(path)
}

// ListCerts lists all CA certificates in the store.
func (cm *CertManager) ListCerts() ([]*x509.Certificate, error) {
	files, err := os.ReadDir(cm.dir)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %v", err)
	}

	var certs []*x509.Certificate
	for _, file := range files {
		path := filepath.Join(cm.dir, file.Name())
		certPEM, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("failed to read certificate: %v", err)
		}
		block, _ := pem.Decode(certPEM)
		if block == nil || block.Type != "CERTIFICATE" {
			return nil, fmt.Errorf("failed to decode PEM block: %s", file.Name())
		}
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse certificate: %v", err)
		}
		certs = append(certs, cert)
	}
	return certs, nil
}
