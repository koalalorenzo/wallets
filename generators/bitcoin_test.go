package generators

import (
	"testing"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

func TestGenerateBTC(t *testing.T) {
	privWIF, pubAddress, err := GenerateBTC()
	if err != nil {
		t.Error(err)
	}

	keyImported, err := btcutil.DecodeWIF(privWIF)
	keyImported.IsForNet(&chaincfg.MainNetParams)
	keyImported.CompressPubKey = false

	if keyImported.String() != privWIF {
		t.Error("WIFs imported, and generated is not matching")
	}

	pubImpKeySerial := keyImported.PrivKey.PubKey().SerializeUncompressed()
	pubKeyImpAddress, err := btcutil.NewAddressPubKey(pubImpKeySerial, &chaincfg.MainNetParams)
	if pubKeyImpAddress.EncodeAddress() != pubAddress {
		t.Error("Public keys are not matching")
	}

}
