package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"os"
	"time"
)

func main() {
	// Generate a 2048-bit RSA key pair
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// Create a self-signed certificate
	serialNumber, _ := rand.Int(rand.Reader, big.NewInt(999999))
	subject := pkix.Name{
		CommonName:   "example.com",
		Organization: []string{"Example Inc."},
		Country:      []string{"US"},
	}
	template := x509.Certificate{
		SerialNumber:          serialNumber,
		Subject:               subject,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}
	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privKey.PublicKey, privKey)
	if err != nil {
		panic(err)
	}

	// Write the private key and certificate to files
	privKeyFile, err := os.Create("privkey.pem")
	if err != nil {
		panic(err)
	}
	defer privKeyFile.Close()
	err = pem.Encode(privKeyFile, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privKey),
	})
	if err != nil {
		panic(err)
	}
	certFile, err := os.Create("cert.pem")
	if err != nil {
		panic(err)
	}
	defer certFile.Close()
	err = pem.Encode(certFile, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: derBytes,
	})
	if err != nil {
		panic(err)
	}
}