package generators

import (
	"crypto/rand"

	"github.com/ehmry/monero"
)

// GenerateXMR returns a Monero private and a public key
func GenerateXMR() (string, string, error) {
	account, err := monero.GenerateAccount(rand.Reader)
	if err != nil {
		return "", "", err
	}
	return account.Address().String(), account.String(), nil
}
