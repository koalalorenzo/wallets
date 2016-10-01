package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"flag"

	"github.com/koalalorenzo/wallets/actions"
)

var coinsOption = flag.String("coins", "all", "the coins to use")
var generateOption = flag.Bool("gen", false, "generate new wallets")
var outputOption = flag.String("o", "", "path of the output (optional)")

var encryptOption = flag.Bool("e", false, "encrypt the private keys instantly")
var encryptAlgoOption = flag.String("ealgo", "AES", "Specify the encryption algorithm between AES and PGP")
var encryptMainOption = flag.String("eopt", "", "if -ealgo == AES: password, if -ealgo == PGP: path of pgp pub key")

var decryptOption = flag.Bool("d", false, "decrypt the private keys")
var decryptOptionPath = flag.String("dpath", "./wallet.json.aes", "path of the encrypted file")

func main() {
	flag.Parse()
	var output string
	var exitStatus int
	// From here Actions (Generate or Decrypt)

	// If the decryption is allowed, the decrypt the output
	if *generateOption == true {
		// Parsing the coins available
		var coins []string
		if *coinsOption != "all" {
			coins = strings.Split(*coinsOption, ",")
		} else {
			coins = []string{}
		}

		allCoins := actions.GenerateAddresses(coins)
		jsonCoins, err := json.Marshal(allCoins)
		if err != nil {
			log.Panic(err)
		}
		output = string(jsonCoins[:])
		exitStatus = 0
	}

	// If the decryption is allowed, then check the file and decrypt it
	if *decryptOption == true && *encryptOption == false {
		output = actions.DecryptOutput(*decryptOptionPath, *encryptAlgoOption, *encryptMainOption)
		exitStatus = 0
	}

	// From here: Output handling

	// If the encryptio is allowed, the encrypt the output
	if *encryptOption == true && *decryptOption == false {
		output = actions.EncryptOutput(output, *encryptAlgoOption, *encryptMainOption)
	}

	if output == "" {
		output = "Invalid options. Check --help "
		*outputOption = "" // Do not save to file
		exitStatus = 1
	}

	if *outputOption == "" {
		fmt.Println(output)
		os.Exit(exitStatus)
	} else {
		actions.SaveToFile(*outputOption, []byte(output))
		os.Exit(exitStatus)
	}

}
