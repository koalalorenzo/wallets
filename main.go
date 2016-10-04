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

// Creation of the wallet action
var generateOption = flag.Bool("new", false, "create a new wallets")
var coinsOption = flag.String("coins", "all", "the coins to use [BTC, ETH, XMR]")
var outputOptionPath = flag.String("save-to", "", "define where (path) to save the wallet file")
var walletOptionPath = flag.String("wallet-file", "", "define where (path) to read the encrypted wallet file")

// Existing wallet info action
var infoOption = flag.Bool("show-info", false, "show the wallet info")

// Decription action
var exportOption = flag.Bool("export-json", false, "decrypt the private keys and print the JSON structure")

// Encryption options
var encryptOption = flag.Bool("encrypt", false, "encrypt the private keys instantly")
var encryptMainOption = flag.String("password", "", "(optional) the password used to encrypt the wallet")

func main() {
	flag.Parse()
	var output string
	var exitStatus int
	// From here Actions (Generate or Decrypt)

	// If the decryption is allowed, the decrypt the output
	if *generateOption == true && len(output) == 0 {
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

	// If the decryption is allowed, the decrypt the output
	if *infoOption == true && len(output) == 0 {
		// Parsing the coins available
		if len(*walletOptionPath) == 0 {
			output = "Please specify a path with --wallet-file"
			exitStatus = 1
		} else {
			output = actions.ShowWalletPubKeys(*walletOptionPath)
			exitStatus = 0
		}
	}

	// If the decryption is allowed, then check the file and decrypt it
	if *exportOption == true && *encryptOption == false && len(output) == 0 {
		output = actions.DecryptOutput(*walletOptionPath, "AES", *encryptMainOption)
		exitStatus = 0
	}

	// From here: Output handling

	// If the encryptio is allowed, the encrypt the output
	if *encryptOption == true && *exportOption == false {
		output = actions.EncryptOutput(output, "AES", *encryptMainOption)
	}

	if output == "" {
		output = "Invalid options/flags. Check --help "
		exitStatus = 1
	}

	if *outputOptionPath == "" || exitStatus != 0 {
		fmt.Println(output)
		os.Exit(exitStatus)
	} else {
		actions.SaveToFile(*outputOptionPath, []byte(output))
		os.Exit(exitStatus)
	}

}
