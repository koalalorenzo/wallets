package actions

import (
	"io/ioutil"
	"log"

	"github.com/koalalorenzo/wallets/encryption"
)

// EncryptOutput is encrypting the input (plainContent string), returning the
// corresponding encrypted value. The options are based on the encryption aglo
// used.
//
// AES: When using AES encryption, you can sepcify the encryptMainOption as the
//      key that you want to use to encrypt the text (16 or 32 bit keys).
//
// PGP: Not yet implemented, but the idea is that encryptMainOption will be
//      used as path for the gpg pub key.
//
func EncryptOutput(plainContent, encryptAlgoOption, encryptMainOption string) string {
	switch encryptAlgoOption {

	case "AES":
		// ToDo Move this into a function?
		key := []byte(encryptMainOption)
		encryptedContent, err := encryption.AESEncrypt(key, plainContent)
		if err != nil {
			log.Panicln(err)
		}
		return encryptedContent

	case "PGP":
	case "OpenPGP":
	case "GPG":
	case "GnuPG":
		log.Fatal("GPG/PGP not supported yet")
		break
	default:
		log.Fatalln("Error: Unknown encryption algorithm ", encryptAlgoOption)
		break
	}
	return ""
}

// DecryptOutput is able to decrypt the content of a file, based on the
// algorithm options. See the comment for EncryptOutput() to understand the
// meaning of encryptAlgoOption and encryptMainOption
func DecryptOutput(decryptOptionPath, encryptAlgoOption, encryptMainOption string) string {
	// Reading the content of the encrypted file
	encryptedFile, err := ioutil.ReadFile(decryptOptionPath)
	if err != nil {
		log.Panic(err)
	}

	switch encryptAlgoOption {

	case "AES":
		// ToDo Move this into a function?
		key := []byte(encryptMainOption)
		content := string(encryptedFile[:])
		encryptedContent, err := encryption.AESDecrypt(key, content)
		if err != nil {
			log.Panic(err)
		}
		return encryptedContent

	case "PGP":
	case "OpenPGP":
	case "GPG":
	case "GnuPG":
		log.Fatal("GPG/PGP not supported yet")
		break
	default:
		log.Fatalln("Error: Unknown encryption algorithm ", encryptAlgoOption)
		break
	}
	return ""
}
