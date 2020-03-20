package addr

import (
	"encoding/hex"
	"github.com/btcsuite/btcd/btcec"
)

type BtcKey struct {
	PrivateKey *btcec.PrivateKey
	PublicKey  *btcec.PublicKey
}

func NewBtcKey(privateKey *btcec.PrivateKey, publicKey *btcec.PublicKey) *BtcKey {
	return &BtcKey{PrivateKey: privateKey, PublicKey: publicKey}
}


// 生成密钥对
func NewKeyFromRand() (*BtcKey, error) {
	privateKey, _ := btcec.NewPrivateKey(btcec.S256())
	publicKey := privateKey.PubKey()
	return &BtcKey{PrivateKey: privateKey, PublicKey: publicKey}, nil
}

// 生成密钥对
func NewKeyFromHex(privKeyHex string) (*BtcKey, error) {
	//encodedStr := "50cde1c490bcd8c546a23f5ece0c637071e702c7b76f9296944f6381a5a2dcf3"
	privKeyBytes, _ := hex.DecodeString(privKeyHex)
	privateKey, publicKey := btcec.PrivKeyFromBytes(btcec.S256(), privKeyBytes)
	return &BtcKey{PrivateKey: privateKey, PublicKey: publicKey}, nil
}

func (slef BtcKey)PubKeyCompressed() ([]byte) {
	return slef.PublicKey.SerializeCompressed()
}

func (slef BtcKey)WifCompressed() ([]byte) {
	return slef.PublicKey.SerializeCompressed()
}
