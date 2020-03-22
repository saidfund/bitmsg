package block

import (
	"encoding/hex"
	"fmt"
	"github.com/saidfund/bitmsg/chunk"
	"log"
	"testing"
)

//var log = logging.MustGetLogger("test")

func TestBlock(t *testing.T) {
	l := &Leading{
		Magic:    0x7425,
		LenBlock: 0,
	}

	fmt.Println(l)

	h := &chunk.Header{
		TimeStamp:  1,
		Option:     2,
		Version:    3,
		HashPrev:   chunk.GetHash('a'),
		HashMerkle: chunk.GetHash('b'),
		Count:      12,
		Extension:  0,
	}
	//fmt.Println(h)

	caller := &chunk.Quater{
		Main: "1LVfRcj31E9mGujxUD3nTJjsUPtcczqJnX",
		Sub:  "send",
	}
	callee := &chunk.Quater{
		Main: "1LVfRcj31E9mGujxUD3nTJjsUPtcczqJnX",
		Sub:  "recv",
	}

	r := chunk.NewRouting(caller, callee)

	//fmt.Println(r)

	c := &chunk.Content{
		HashLeaf: chunk.GetHash('3'),
		Cipher:   []byte{},
		Body:     []byte("Hello bitmsg"),
	}
	//fmt.Println(c)
	c.Marshal()
	b := NewBlock(*h, *r)
	b.AddContent(*c)

	log.Println(hex.Dump(b.Marshal()))
}
