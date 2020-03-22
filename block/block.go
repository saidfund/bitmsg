package block

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/fananchong/cstruct-go"
	"github.com/saidfund/bitmsg/chunk"
)

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
	copy(self.Header.ShaMerkle[:], m.RootNode.Data)
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
