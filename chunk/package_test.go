package chunk

import (
	"fmt"
	"testing"
)

func TestPackage(t *testing.T) {
	l := &Leading{
		Magic:    0x7425,
		LenBlock: 0,
	}

	fmt.Println(l)

	h := &Header{
		TimeStamp:  1,
		Option:     2,
		Version:    3,
		HashPrev:   getHash('a'),
		HashMerkle: getHash('b'),
		Count:      12,
		Extension:  0,
	}
	fmt.Println(h)

	caller := &Quater{
		Main: "1LVfRcj31E9mGujxUD3nTJjsUPtcczqJnX",
		Sub:  "send",
	}
	callee := &Quater{
		Main: "1LVfRcj31E9mGujxUD3nTJjsUPtcczqJnX",
		Sub:  "recv",
	}

	r := NewRouting(caller, callee)

	fmt.Println(r)

	c := &Content{
		HashLeaf: getHash('3'),
		Cipher:   []byte{},
		Body:     []byte("Hello bitmsg"),
	}
	fmt.Println(c)
	c.Marshal()

}
