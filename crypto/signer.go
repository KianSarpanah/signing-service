package crypto

type Signer interface {
	Sign(data string) (string, error)
	PublicKey() string
}
