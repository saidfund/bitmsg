package chunk

import (
	"fmt"
	"testing"
)

func TestCell(t *testing.T) {
	h := &Header{
		TimeStamp:  1,
		Option:     2,
		Version:    3,
		HashPrev:   GetHash(0),
		HashMerkle: GetHash(1),
		Count:      12,
		Extension:  0,
	}
	fmt.Println(h)
	data := h.Marshal()
	i := UnMarshal(data)
	fmt.Println(i)
}
