package addr

import (
	"bytes"
	"compress/gzip"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"golang.org/x/crypto/ripemd160"
	"log"
	"math/big"
	"strings"
)
//VERSION = uint8(0x89) x
//VERSION = uint8(0x12) 8
//VERSION = uint8(0x32) M

const (
	VERSION = uint8(0x12)
	LenAddChechsum = 4
	LenPrivKey     = 32
)

type SayKey struct {
	privateKey *ecdsa.PrivateKey
	PublicKey  ecdsa.PublicKey
}

// 生成密钥对
func MakeNewKeyP256() (*SayKey, error) {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	sayKey := SayKey{private, private.PublicKey}
	return &sayKey, nil
}

// 生成密钥对
func MakeNewKeyP256Str(randReader string) (*SayKey, error) {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, strings.NewReader(randReader))
	if err != nil {
		log.Panic(err)
	}
	sayKey := SayKey{private, private.PublicKey}
	return &sayKey, nil
}


func (self *SayKey) GetPrivKey() []byte {
	buf := make([]byte, 0, LenPrivKey)
	priKey := Padded(LenPrivKey, buf, self.privateKey.D.Bytes()) // []bytes type
	return priKey
}

func (self *SayKey) GetPubKey() []byte {
	pubKey := append(self.PublicKey.X.Bytes(), self.privateKey.Y.Bytes()...) // []bytes type
	return pubKey
}

// 得到地址
/* See https://en.bitcoin.it/wiki/Technical_background_of_Bitcoin_addresses */
func (self *SayKey) GetAddressStep() (address string) {
	pubKey := self.GetPubKey()

	/* SHA256 HASH */
	fmt.Println("1 - Perform SHA-256 hashing on the public key")
	sumSha256 := sha256.Sum256(pubKey)

	/* RIPEMD-160 HASH */
	fmt.Println("2 - Perform RIPEMD-160 hashing on the result of SHA-256")
	hasher := ripemd160.New()
	hasher.Reset()
	hasher.Write(sumSha256[:])
	sumRipemd160 := hasher.Sum(nil) // 对公钥hash进行ripemd160运算

	/* Convert hash bytes to base58 chech encoded sequence */
	address = B58EncodeCheck(VERSION, sumRipemd160)
	return address
}

func (self *SayKey) GetAddress() (address string) {
	sum := self.DHashPubKey(self.GetPubKey())
	address = B58EncodeCheck(VERSION, sum)
	return
}

func (self *SayKey) LeftAddress(right int) (string) {
	result := make([]byte, 0, 100)
	buf := []byte(self.GetAddress())
	fmt.Println(hex.Dump(buf))

	result = append(result, buf[:1]...)
	left := buf[len(buf) - right:]

	buf = buf[1:len(buf) - right]
	fmt.Println(hex.Dump(left))
	fmt.Println(hex.Dump(buf))

	l := len(buf)
	for i:= 0; i < l; i++{
		if i % 6 == 0{
			result = append(result, '-')
			j := i + 6
			if j >= l{
				j = l
				result = append(result, buf[i:j]...)
				break
			} else{
				result = append(result, buf[i:j]...)
			}
		}
	}
	result = append(result, '-')
	result = append(result, left...)
	return string(result)
}

// 把pubkey进行二次哈希运算
func (self *SayKey) DHashPubKey(pubKey []byte) []byte {
	sumSHA256 := sha256.Sum256(pubKey) // 第一次哈希运算
	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(sumSHA256[:])
	if err != nil {
		log.Panic(err)
	}
	sumRIPEMD160 := RIPEMD160Hasher.Sum(nil) // 第二次哈希运算
	return sumRIPEMD160
}

// b58chechencode encodes version ver and byte slice b into a base-58 chech encoded string.
func (self *SayKey)B58EncodeCheck(ver uint8, sumRIPEMD160 []byte) (result string) {
	// Prepend version
	buf := append([]byte{ver}, sumRIPEMD160...)
	// Create a new SHA256 context
	sumSHA256 := sha256.Sum256(buf) // 第一次哈希运算
	sumSHA256 = sha256.Sum256(sumSHA256[:]) // 第二次哈希运算
	//Add the 4 checksum bytes at the end of extended PIPEMD-160 hash
	buf = append(buf, sumSHA256[0:4]...)
	// Encode base58 string
	result = B58Encode(buf)
	/* For number  of leading 0'result in bytes, prepend 1 */
	for _, v := range buf {
		if v != 0 {
			break
		}
		result = "1" + result
	}
	return result
}


/*
对text签名
返回加密结果，结果为数字证书r、s的序列化后拼接，然后用hex转换为string
*/
func (self *SayKey) Sign(text []byte) ([]byte, error) {
	r, s, err := ecdsa.Sign(rand.Reader, self.privateKey, text)
	if err != nil {
		return nil, err
	}
	rt, err := r.MarshalText()
	if err != nil {
		return nil, err
	}
	st, err := s.MarshalText()
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	defer w.Close()
	_, err = w.Write([]byte(string(rt) + "+" + string(st)))
	if err != nil {
		return nil, err
	}
	w.Flush()
	return buf.Bytes(), nil
}


/*
校验文本内容是否与签名一致
使用公钥校验签名和文本内容
*/
func (self *SayKey)Verify(text []byte, signature []byte, pubKey *ecdsa.PublicKey) (bool, error) {
	rint, sint, err := self.getSign(signature)
	if err != nil {
		return false, err
	}
	result := ecdsa.Verify(pubKey, text, &rint, &sint)
	return result, nil
}

/*
证书分解
通过hex解码，分割成数字证书r，s
*/
func (self *SayKey)getSign(signature []byte) (rint, sint big.Int, err error) {
	reader, err := gzip.NewReader(bytes.NewBuffer(signature))
	if err != nil {
		err = errors.New("decode error," + err.Error())
		return
	}
	defer reader.Close()
	buf := make([]byte, 1024)
	count, err := reader.Read(buf)
	if err != nil {
		fmt.Println("decode = ", err)
		err = errors.New("decode read error," + err.Error())
		return
	}
	rs := bytes.Split(buf[:count], []byte("+"))
	if len(rs) != 2 {
		err = errors.New("decode fail")
		return
	}
	err = rint.UnmarshalText([]byte(rs[0]))
	if err != nil {
		err = errors.New("decrypt rint fail, " + err.Error())
		return
	}
	err = sint.UnmarshalText([]byte(rs[1]))
	if err != nil {
		err = errors.New("decrypt sint fail, " + err.Error())
		return
	}
	return
}

