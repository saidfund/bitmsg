package block

import (
	"encoding/hex"
	"fmt"
	"github.com/davyxu/golog"
	"testing"
)

func TestRouting_Dump(t *testing.T) {
	var log *golog.Logger = golog.New("test")

	routing := NewRouting([]byte("addr"), []byte("sub"))
	log.Infof("%s\n", routing.String())

	fmt.Println(routing.Address())
	fmt.Println(routing.Subress())

	buf, _ := routing.Dump()
	fmt.Println(hex.Dump(buf))

	routing = NewRoutingUnDump(buf)
	fmt.Println(hex.Dump(routing.Addr))
	fmt.Println(hex.Dump(routing.Sub))
}
