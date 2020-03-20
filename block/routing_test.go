package block

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestRouting_Dump(t *testing.T) {
	routing := NewRouting([]byte("addr"), []byte("sub"))
	fmt.Println(routing)

	fmt.Println(routing.Address())
	fmt.Println(routing.Subress())

	buf, _ := routing.Dump()
	fmt.Println(hex.Dump(buf))

	routing = NewRoutingUnDump(buf)
	fmt.Println(hex.Dump(routing.Addr))
	fmt.Println(hex.Dump(routing.Sub))
}