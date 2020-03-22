package block

import (
	"encoding/hex"
	"fmt"

	"github.com/saidfund/bitmsg/chunk"
	"testing"
)

//var log = logging.MustGetLogger("test")

func TestBlock(t *testing.T) {
	l := &Leading{
		Magic:    0x7425,
		LenBlock: 0,
	}

	fmt.Println(l)

	h := &chunk.Header{}
	fmt.Println(hex.Dump(h.Marshal()))

	caller := &chunk.Quater{
		Main: "1LVfRcj31E9mGujxUD3nTJjsUPtcczqJnX",
		Sub:  "send",
	}
	callee := &chunk.Quater{
		Main: "1LVfRcj31E9mGujxUD3nTJjsUPtcczqJnX",
		Sub:  "recv",
	}

	r := chunk.NewRouting(caller, callee)
	fmt.Println(hex.Dump(r.Marshal()))

	c := &chunk.Content{
		Sha:    chunk.GetHash('0'),
		Mime:   0xff,
		Cipher: []byte{},
		Body:   []byte("Hello bitmsg"),
	}
	//fmt.Println(c)
	c.Marshal()
	b := NewBlock(*h, *r)
	b.AddContent(*c)
	fmt.Println(hex.Dump(b.Header.ShaMerkle[:]))
	fmt.Println(hex.Dump(b.Marshal()))
}
