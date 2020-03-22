package chunk

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/fananchong/cstruct-go"
)

type Content struct {
	Length uint32
	Mime   uint32   //单元内容类型
	Sha    [32]byte //单元哈希值
	Cipher []byte   //密码
	Body   []byte   //内容
}

func (self *Content) Marshal() []byte {
	self.Length = uint32(4 + 32 + 4 + len(self.Cipher) + 4 + len(self.Body))
	hash := sha256.Sum256(self.Body)
	copy(self.Sha[:], hash[:])
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
		TypeOf:   self.Mime,
		HashLeaf: string(self.Sha[:]),
		Cipher:   string(self.Cipher),
		Body:     string(self.Body),
	}

	data, _ := json.Marshal(content)
	return fmt.Sprintf("%s", data)
}
