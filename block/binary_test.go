package block

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestBytes_ToUint32(t *testing.T) {
	ui := uint32(65535)
	bs := FromUint32(ui)
	fmt.Printf(hex.Dump(bs))
}

func TestFromUint32(t *testing.T) {
	bs := []byte{0x00, 0x00, 0xff, 0xff, 0x34}
	b := Bytes(bs)
	ui, size := b.ToUint32()
	bs = bs[size:]
	fmt.Printf("%d\n", ui)
	fmt.Printf(hex.Dump(bs))
}

func TestBytesToInt64(t *testing.T) {
	i := int64(1531293019)
	b := FromInt64(i)
	fmt.Printf(hex.Dump(b))
}

func TestFromInt64(t *testing.T) {
	bs := []byte{0x00, 0x00, 0x00, 0x00, 0x5b, 0x45, 0xad, 0x5b}
	b := Bytes(bs)
	i, size := b.ToInt64()
	bs = bs[size:]
	fmt.Printf("%d\n", i)
	fmt.Printf(hex.Dump(bs))

}
