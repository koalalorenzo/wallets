package generators

import (
	"encoding/hex"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

// GenerateBTC returns a Bitcoin private and a public key
func GenerateBTC() (string, string, error) {
	privKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return "", "", err
	}

	privKeyWif, err := btcutil.NewWIF(privKey, &chaincfg.MainNetParams, false)
	if err != nil {
		return "", "", err
	}
	pubKeySerial := privKey.PubKey().SerializeUncompressed()

	pubKeyAddress, err := btcutil.NewAddressPubKey(pubKeySerial, &chaincfg.MainNetParams)
	if err != nil {
		return "", "", err
	}

	return privKeyWif.String(), pubKeyAddress.EncodeAddress(), nil
}

//SignBTCHexTransaction will sign a transaction (hex format) and return its hex
// WIP
func SignBTCHexTransaction(hexTx, privWIFKey string) (string, error) {

	txHexBytes, err := hex.DecodeString(hexTx)
	if err != nil {
		return "", err
	}

	wifKey, err := btcutil.DecodeWIF(privWIFKey)
	wifKey.IsForNet(&chaincfg.MainNetParams)
	wifKey.CompressPubKey = false

	signature, err := wifKey.PrivKey.Sign(txHexBytes)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(signature.Serialize()), nil
}
