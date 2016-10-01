package actions

import (
	"log"

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
		coins = []string{"ETH", "BTC"}
	}

	generatedCoins := []Address{}

	for i := range coins {
		crypto := coins[i]

		switch crypto {

		case "ETH":
			pub, priv, err := generators.GenerateETH()
			if err != nil {
				log.Fatalf("Generating ETH: %s\n", err)
			}
			newAddress := Address{"ETH", priv, pub}
			generatedCoins = append(generatedCoins, newAddress)
			break

		case "BTC":
			pub, priv, err := generators.GenerateBTC()
			if err != nil {
				log.Fatalf("Generating BTC: %s\n", err)
			}
			newAddress := Address{"BTC", priv, pub}
			generatedCoins = append(generatedCoins, newAddress)
			break

		default:
			log.Fatalln("Error: Unknown ", crypto, "... would you like to implement it?")
			break
		}

	}
	return generatedCoins
}
