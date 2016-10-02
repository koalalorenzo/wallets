package actions

import (
	"encoding/json"
	"strings"
)

// ShowWalletPubKeys shows wallet's list of pub keys
func ShowWalletPubKeys(walletPath string) string {
	decryptedWallet := DecryptOutput(walletPath, "AES", "")

	var wallet = []Address{}
	var pubKeys = []string{}

	json.Unmarshal([]byte(decryptedWallet), &wallet)
	for nthAddress := range wallet {
		address := wallet[nthAddress]
		pubKeys = append(pubKeys, address.Coin+": "+address.PubKey)
	}
	return strings.Join(pubKeys, "\n")
}

// ShowWalletIPrivKeys shows wallet's list of private keys
func ShowWalletIPrivKeys(walletPath string) string {
	decryptedWallet := DecryptOutput(walletPath, "AES", "")

	var wallet = []Address{}
	var privKeys = []string{}

	json.Unmarshal([]byte(decryptedWallet), &wallet)
	for nthAddress := range wallet {
		address := wallet[nthAddress]
		privKeys = append(privKeys, address.Coin+": "+address.PrivKey)
	}
	return strings.Join(privKeys, "\n")
}
