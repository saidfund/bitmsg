package chunk

import (
	"encoding/json"
	"fmt"
	"github.com/fananchong/cstruct-go"
)

type Content struct {
	TypeOf   uint32   //单元内容类型
	HashLeaf [32]byte //单元哈希值
	Cipher   []byte   //密码
	Body     []byte   //内容
}

func (self *Content) Marshal() []byte {
	data, _ := cstruct.Marshal(self)
	//fmt.Println(hex.Dump(data))
	return data
}

func (self *Content) UnMarshal(data []byte) {
	cstruct.Unmarshal(data, self)
}

func (self *Content) String() string {
	var content = struct {
		TypeOf   uint32 //单元内容类型
		HashLeaf string //单元哈希值
		Cipher   string //密码
		Body     string //内容
	}{
		TypeOf:   self.TypeOf,
		HashLeaf: string(self.HashLeaf[:]),
		Cipher:   string(self.Cipher),
		Body:     string(self.Body),
	}

	data, _ := json.Marshal(content)
	return fmt.Sprintf("%s", data)
}
