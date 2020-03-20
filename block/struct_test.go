package block

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestStructs(t *testing.T) {
	s := new(Structs)
	s.Format = []string{"H", "H", "I", "16s", "I", "I", "I"}
	pack := s.Pack(uint16(1), uint16(1), uint32(1), "abc", uint32(1), uint32(0), uint32(12))
	fmt.Println(hex.Dump(pack))

	r := s.UnPack(pack)
	a := r[0].(uint16)

	fmt.Println(a)
}