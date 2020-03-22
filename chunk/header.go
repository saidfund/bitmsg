package chunk

import (
	"encoding/json"
	"fmt"
	"github.com/fananchong/cstruct-go"
	"time"
)

const VER_HEADER uint32 = 0x10

type Header struct {
	TimeStamp int64    //时间戳，整数，固定长度4个字节
	Option    uint32   //区块参数, 暂时不用
	Version   uint32   //版本号
	ShaPrev   [32]byte //前一个区块的哈希,长度32
	ShaMerkle [32]byte //梅克尔树哈希,长度32
	Count     uint32   //消息单元个数
	Extension uint32   //扩展头部长度， 暂时不用
}

func UnMarshal(data []byte) (self *Header) {
	self = &Header{}
	cstruct.Unmarshal(data, self)
	return self
}

func (self *Header) Marshal() []byte {
	self.TimeStamp = time.Now().Unix()
	self.Version = VER_HEADER
	self.ShaPrev = [32]byte{}
	data, _ := cstruct.Marshal(self)
	//fmt.Println(hex.Dump(data))
	return data
}

func (self *Header) String() string {
	var header = struct {
		TimeStamp  int64  //时间戳，整数，固定长度4个字节
		Option     uint32 //区块参数, 暂时不用
		Version    uint32 //版本号
		HashPrev   string //前一个区块的哈希,长度32
		HashMerkle string //梅克尔树哈希,长度32
		Count      uint32 //消息单元个数
		Extension  uint32 //扩展头部长度， 暂时不用
	}{
		TimeStamp:  self.TimeStamp,
		Option:     self.Option,
		Version:    self.Version,
		HashPrev:   string(self.ShaPrev[:]),
		HashMerkle: string(self.ShaMerkle[:]),
		Count:      self.Count,
		Extension:  self.Extension,
	}
	data, _ := json.Marshal(header)
	return fmt.Sprintf("%s", data)
}
