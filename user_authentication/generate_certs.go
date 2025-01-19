//generated code, check it

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"log"
	"math/big"
	"os"
	"time"
)

func main() {
	// Generate a private key for the CA
	caPrivateKey, err := rsa.GenerateKey(rand.Reader, 3072)
	if err != nil {
		log.Fatalf("Failed to generate private key: %v", err)
	}

	// Create the CA certificate template
	caTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization:  []string{"Secure Communication CA"},
			Country:       []string{"Botswana"},
			Province:      []string{"Kveneng"},
			Locality:      []string{"Molepolole"},
			StreetAddress: []string{"123 Secure St"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0), // 10 years validity
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
	}

	// Self-sign the certificate
	caCert, err := x509.CreateCertificate(rand.Reader, caTemplate, caTemplate, &caPrivateKey.PublicKey, caPrivateKey)
	if err != nil {
		log.Fatalf("Failed to create CA certificate: %v", err)
	}

	// Save the private key to a file
	caKeyFile, err := os.Create("ca-key.pem")
	if err != nil {
		log.Fatalf("Failed to create key file: %v", err)
	}
	defer caKeyFile.Close()
	pem.Encode(caKeyFile, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(caPrivateKey)})

	// Save the certificate to a file
	caCertFile, err := os.Create("ca-cert.pem")
	if err != nil {
		log.Fatalf("Failed to create certificate file: %v", err)
	}
	defer caCertFile.Close()
	pem.Encode(caCertFile, &pem.Block{Type: "CERTIFICATE", Bytes: caCert})

	log.Println("CA certificate and private key generated successfully!")
}