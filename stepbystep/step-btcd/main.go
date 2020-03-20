package main

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

func main() {
	encodedStr := "c28a9f80738f770d527803a566cf6fc3edf6cea586c4fc4a5223a5ad797e1ac3"
	privKeyBytes, _ := hex.DecodeString(encodedStr)
	privKey, pubKey := btcec.PrivKeyFromBytes(btcec.S256(), privKeyBytes)
	//privKey, _ := btcec.NewPrivateKey(btcec.S256())
	fmt.Println(hex.Dump(privKey.Serialize()))
	//pubKey := privKey.PubKey()
	fmt.Println(hex.Dump(pubKey.SerializeUncompressed()))
	fmt.Println(hex.Dump(pubKey.SerializeCompressed()))
	pubKeyHash := btcutil.Hash160(pubKey.SerializeCompressed())
	addr, err := btcutil.NewAddressPubKeyHash(pubKeyHash, &chaincfg.MainNetParams)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(addr)
	wif, err := btcutil.NewWIF(privKey, &chaincfg.MainNetParams, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(wif)
}
