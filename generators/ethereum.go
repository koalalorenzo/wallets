package generators

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
)

// GenerateETH returns a private and a public key
func GenerateETH() (string, string, error) {
	key, _ := crypto.GenerateKey()
	addr := crypto.PubkeyToAddress(key.PublicKey)

	prvKey := hex.EncodeToString(crypto.FromECDSA(key))
	pubKey := "0x" + hex.EncodeToString(addr[:])
	return prvKey, pubKey, nil
}
