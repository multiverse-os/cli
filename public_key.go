package cli

type KeyType int

const (
	RSA KeyType = iota
	ECDSA
	ED25199
)

type PublicKey struct {
	Type      KeyType
	PublicKey string
}
