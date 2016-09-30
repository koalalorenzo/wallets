package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"flag"

	"github.com/koalalorenzo/wallets/generators"
)

var coins = flag.String("coins", "all", "the coins to use")
var help = flag.Bool("help", false, "Help message")

func main() {
	flag.Parse()

	if *help != false {
		fmt.Println("HELP")
		os.Exit(0)
	}

	leftovers := strings.Split(*coins, ",")
	privates := [][]string{}

	for i := range leftovers {
		crypto := leftovers[i]

		switch crypto {

		case "ETH":
			pub, priv, err := generators.GenerateETH()
			if err != nil {
				log.Fatalf("Generating ETH: %s\n", err)
			}
			privates = append(privates, []string{"ETH", pub, priv})

		case "BTC":
			pub, priv, err := generators.GenerateBTC()
			if err != nil {
				log.Fatalf("Generating BTC: %s\n", err)
			}
			privates = append(privates, []string{"BTC", pub, priv})

		default:
			break
		}

	}

	for i := range privates {
		fmt.Println(privates[i])
	}

}
