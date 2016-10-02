package encryption

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/openpgp"
)

// PGPEncrypt will encrypt a string.
func PGPEncrypt(secretString, pubKeyPath string) (string, error) {
	keyringFileBuffer, _ := os.Open(pubKeyPath)
	defer keyringFileBuffer.Close()
	entityList, err := openpgp.ReadKeyRing(keyringFileBuffer)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	w, err := openpgp.Encrypt(buf, entityList, nil, nil, nil)
	if err != nil {
		return "", err
	}

	_, err = w.Write([]byte(secretString))
	if err != nil {
		return "", err
	}

	err = w.Close()
	if err != nil {
		return "", err
	}
	bytes, err := ioutil.ReadAll(buf)
	if err != nil {
		return "", err
	}

	base64EncStr := base64.StdEncoding.EncodeToString(bytes)
	return base64EncStr, nil
}
