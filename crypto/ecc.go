package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/base64"
	"math/big"
)

// ECCSigner holds the ECC private key
type ECCSigner struct {
	privateKey *ecdsa.PrivateKey
}

// NewECCSigner creates a new ECCSigner with a generated private key
func NewECCSigner() (*ECCSigner, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	return &ECCSigner{privateKey: privateKey}, nil
}

// ecdsaSignature holds the R and S values of an ECDSA signature
type ecdsaSignature struct {
	R, S *big.Int
}

// Sign signs the given data and returns the signature as a base64 encoded string
func (signer *ECCSigner) Sign(data string) (string, error) {
	hash := sha256.Sum256([]byte(data))
	r, s, err := ecdsa.Sign(rand.Reader, signer.privateKey, hash[:])
	if err != nil {
		return "", err
	}

	// Correctly marshal r and s into the ecdsaSignature struct
	signature, err := asn1.Marshal(ecdsaSignature{R: r, S: s})
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}

// PublicKey returns the public key as a base64 encoded string
func (signer *ECCSigner) PublicKey() string {
	pubASN1, err := x509.MarshalPKIXPublicKey(&signer.privateKey.PublicKey)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(pubASN1)
}
