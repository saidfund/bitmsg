package block

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/fananchong/cstruct-go"
	"github.com/saidfund/bitmsg/chunk"
)

type Package struct {
	Leading Leading
	Block   Block
}

type Leading struct {
	Magic    uint16 //魔法数字
	LenBlock uint32 //区块的长度
}

func LeadingUnMarshal(data []byte) (self *chunk.Header) {
	self = &chunk.Header{}
	cstruct.Unmarshal(data, self)
	return self
}

func (self *Leading) Marshal() []byte {
	data, _ := cstruct.Marshal(self)
	fmt.Println(hex.Dump(data))
	return data
}

func (self *Leading) UnMarshal(data []byte) {
	cstruct.Unmarshal(data, self)
}

func (self *Leading) String() string {
	data, _ := json.Marshal(self)
	return fmt.Sprintf("%s", data)
}
