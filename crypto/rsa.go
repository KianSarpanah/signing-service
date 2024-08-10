package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
)

type RSASigner struct {
	privateKey *rsa.PrivateKey
}

func NewRSASigner() (*RSASigner, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	return &RSASigner{privateKey: privateKey}, nil
}

func (s *RSASigner) Sign(data string) (string, error) {
	hashed := sha256.Sum256([]byte(data))
	signature, err := rsa.SignPKCS1v15(rand.Reader, s.privateKey, 0, hashed[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}

func (s *RSASigner) PublicKey() string {
	pubASN1, err := x509.MarshalPKIXPublicKey(&s.privateKey.PublicKey)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(pubASN1)
}
