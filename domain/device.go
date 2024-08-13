package domain

import (
	"signaturesign/crypto"
)

type Device struct {
	ID               string        `json:"ID"`
	Algorithm        string        `json:"Algorithm"`
	Label            string        `json:"Label"`
	PublicKey        string        `json:"PublicKey"`
	SignatureCounter int           `json:"SignatureCounter"`
	LastSignature    string        `json:"LastSignature"`
	Signer           crypto.Signer `json:"Signer"`
}
