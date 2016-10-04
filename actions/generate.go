package actions

import (
	"log"
	"strings"

	"github.com/koalalorenzo/wallets/generators"
)

// Address will store the addresses generated
type Address struct {
	Coin    string
	PubKey  string
	PrivKey string
}

// GenerateAddresses will simply returns a list of Addresses generated based
// on the list of coins passed. If none is passed, then all of them will be
// generated. Currently supporting ETH and BTC only.
func GenerateAddresses(coins []string) []Address {
	if len(coins) == 0 {
		coins = []string{"ETH", "BTC", "XMR"}
	}

	generatedCoins := []Address{}

	for i := range coins {
		crypto := coins[i]
		crypto = strings.ToUpper(crypto)

		switch crypto {

		case "ETC":
		case "ETH":
			pub, priv, err := generators.GenerateETH()
			if err != nil {
				log.Fatalf("Generating ETH: %s\n", err)
			}
			newAddress := Address{crypto, priv, pub}
			generatedCoins = append(generatedCoins, newAddress)
			break

		case "BTC":
			pub, priv, err := generators.GenerateBTC()
			if err != nil {
				log.Fatalf("Generating BTC: %s\n", err)
			}
			newAddress := Address{crypto, priv, pub}
			generatedCoins = append(generatedCoins, newAddress)
			break

		case "XMR":
			pub, priv, err := generators.GenerateXMR()
			if err != nil {
				log.Fatalf("Generating XMR: %s\n", err)
			}
			newAddress := Address{crypto, priv, pub}
			generatedCoins = append(generatedCoins, newAddress)
			break

		default:
			log.Fatalln("Error: Unknown ", crypto, "... would you like to implement it?")
			break
		}

	}
	return generatedCoins
}
