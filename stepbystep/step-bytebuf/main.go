package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

func main() {
	buf := bytes.NewBuffer([]byte{})
	buf.Write([]byte{0xef})
	fmt.Printf(hex.Dump(buf.Bytes()))


	Sign := "ABC"
	buf.Write([]byte(Sign))
	fmt.Printf(hex.Dump(buf.Bytes()))
}
