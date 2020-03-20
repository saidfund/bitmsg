package block

import (
	"bytes"
	"encoding/json"
	"fmt"
)

/*self.time = time  # 时间戳，整数，固定长度4个字节
self.option = option
self.version = version  # 版本号，固定长度2个字节
self.hash_prev_block = hash_prev_block  # 上一个块的hash，固定长度32个字节
self.hash_merkle_root = hash_merkle_root  # 梅克尔树的根hash，固定长度32个字节
*/

type Header struct {
	TimeStamp  int64  //时间戳，整数，固定长度4个字节
	Option     uint32 //区块参数
	Version    uint32 //版本号
	HashPrev   []byte //前一个区块的哈希,长度32
	HashMerkle []byte //梅克尔树哈希,长度32
	Count      uint32 //消息单元个数
	Extension  uint32 //扩展头部长度
}

func NewBlockUnDump(data []byte) (*Header, error) {
	structs := new(Structs)
	structs.Format = []string{"q", "I", "I", "32s", "32s", "I", "I"}
	result := structs.UnPack(data)
	timeStamp := result[0].(int64)
	option := result[1].(uint32)
	version := result[2].(uint32)
	hashPrev := []byte(result[3].(string))
	hashMerkle := []byte(result[4].(string))
	count := result[5].(uint32)
	extension := result[6].(uint32)
	return &Header{TimeStamp: timeStamp, Option: option, Version: version,
		HashPrev: hashPrev, HashMerkle: hashMerkle, Count: count, Extension: extension}, nil
}

func NewHeader(timeStamp int64, option uint32, version uint32, hashPrev []byte, hashMerkle []byte, count uint32, extension uint32) *Header {
	return &Header{TimeStamp: timeStamp, Option: option, Version: version,
		HashPrev: hashPrev, HashMerkle: hashMerkle, Count: count, Extension: extension}
}

func (self *Header) String() string {
	data, _ := json.Marshal(self)
	return fmt.Sprintf("%s", data)
}

func (self *Header) Dump() ([]byte, error) {
	bigEndian := true
	buf := bytes.NewBuffer([]byte{})
	buf.Write(Int64ToBytes(self.TimeStamp, bigEndian))
	buf.Write(Uint32ToBytes(self.Option, bigEndian))
	buf.Write(Uint32ToBytes(self.Version, bigEndian))
	buf.Write(self.HashPrev)
	buf.Write(self.HashMerkle)
	buf.Write(Uint32ToBytes(self.Count, bigEndian))
	buf.Write(Uint32ToBytes(self.Extension, bigEndian))

	return buf.Bytes(), nil
}

