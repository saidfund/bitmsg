package addr

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
)

func TestKey(t *testing.T) {
	addrs := make([]string, 0, 1024)
	for i := 0; i < 5; i++ {
		key, err := MakeNewKeyP256()
		if err != nil {
			fmt.Println(err)
			return
		}
		address := key.GetAddress()
		addrs = append(addrs, address)
		fmt.Println("My address is:", address)
	}
	fmt.Println("-----------------------")
	fmt.Println(addrs)
}

func TestKeyAddress(t *testing.T) {
	randread := strings.Repeat("1", 41)
	key, err := MakeNewKeyP256Str(randread)
	if err != nil {
		fmt.Println(err)
		return
	}
	address := key.GetAddress()
	address_2 := key.GetAddressStep()
	fmt.Println(address)
	fmt.Println(address_2)
	fmt.Println(key.LeftAddress(6))
}


func TestSayKey_Verify(t *testing.T) {
	key, err := MakeNewKeyP256()
	if err != nil {
		fmt.Println(err)
		return
	}
	privKey := key.GetPrivKey()
	fmt.Println("My PrivateKey is :")
	fmt.Println(hex.Dump(privKey))
	pubKey := key.GetPubKey()
	fmt.Println("My publickKey is :")
	fmt.Println(hex.Dump(pubKey))
	address := key.GetAddress()
	fmt.Println("My address is:", address)
	text := []byte("hahahaha~!")
	signature, _ := key.Sign(text)
	fmt.Println("Signature is :", signature)
	isSuccess, _ := key.Verify(text, signature, &key.PublicKey)
	if isSuccess == true {
		fmt.Println("Verify success")
	}

}