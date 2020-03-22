package block

import (
	"bytes"
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

type Block struct {
	Header  chunk.Header
	Routing chunk.Routing
	Content []chunk.Content
}

func NewBlock(header chunk.Header, routing chunk.Routing) *Block {
	content := []chunk.Content{}
	return &Block{Header: header, Routing: routing, Content: content}
}

func (self *Block) AddContent(content chunk.Content) {
	self.Content = append(self.Content, content)
	data := [][]byte{}
	for _, v := range self.Content {
		data = append(data, v.Marshal())
	}
	m := chunk.NewMerkleTree(data)
	copy(self.Header.HashMerkle[:], m.RootNode.Data)
	self.Header.Count = uint32(len(self.Content))
}

func (self *Block) Marshal() []byte {
	/*	helloWord := [][]byte{ helloBytes,worldBytes }
		fmt.Println( helloWord )
		helloWords := bytes.Join(helloWord,[]byte{})
	*/
	var b bytes.Buffer
	h := self.Header.Marshal()
	b.Write(h)
	r := self.Routing.Marshal()
	b.Write(r)
	for _, v := range self.Content {
		b.Write(v.Marshal())
	}
	return b.Bytes()
}
