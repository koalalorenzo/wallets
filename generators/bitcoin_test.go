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

func TestSignBTCHexTransaction(t *testing.T) {
	hexTx := "01000000013ba3edfd7a7b12b27ac72c3e67768f617fc81bc3888a51323a9fb8aa4b1e5e4a0000000000ffffffff0100f2052a010000001976a91462e907b15cbf27d5425399ebf6f0fb50ebb88f1888ac00000000"
	privKeyWIF := "L4BGmD3qJAhUWqTY1x9oVWhGoWF1m7i6Yo6F6mvBdchro9jxkUbk"
	expectedHexTx := "01000000013ba3edfd7a7b12b27ac72c3e67768f617fc81bc3888a51323a9fb8aa4b1e5e4a000000006b483045022100e4e93e630ac4891b60d1373618748a5f9eb7fe9645f92652fecb6e5f25b4aa8e022035696b89e9509ddbd504b6fbfefe87c465ddc9a1e4027a6d49998fc4609735440121031f7a9d57bde61d47996a39b6d79b67015058a05ccb9a13250c27a13ebf5f181cffffffff0100f2052a010000001976a91462e907b15cbf27d5425399ebf6f0fb50ebb88f1888ac00000000"

	signedTx, err := SignBTCHexTransaction(hexTx, privKeyWIF)
	if err != nil {
		t.Error(err)
	}

	if expectedHexTx != signedTx {
		t.Error("Signatures are not matching", signedTx, expectedHexTx)
	}

}
