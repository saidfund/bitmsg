package main

import (
	"encoding/hex"
	"fmt"
	"github.com/fananchong/cstruct-go"
)

type mys struct {
	F5 []byte
}

func main() {
	m := &mys{
		F5: []byte{},
	}

	data, _ := cstruct.Marshal(m)
	fmt.Println(hex.Dump(data))

}
