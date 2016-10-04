package generators

import (
	"crypto/rand"
	"strings"

	"github.com/ehmry/monero"
)

// GenerateXMR returns a Monero private and a public key
func GenerateXMR() (string, string, error) {
	account, err := monero.GenerateAccount(rand.Reader)
	if err != nil {
		return "", "", err
	}
	mnemonic, err := account.Mnemonic()
	if err != nil {
		return "", "", err
	}
	mnemonicStr := strings.Join(mnemonic, " ")
	return mnemonicStr, account.Address().String(), nil
}
