package block

import (
	"encoding/json"
	"fmt"
)

/*self.magic = magic  # 魔数，封包的标识
self.pkg_len = pkg_len  # 封包长度
self.sender_sign = sender_sign  # 发送方签名
*/

type Leading struct {
	Magic uint16	//魔法数字
	LenBlock uint32	//区块的长度
	LenSign uint16	//签名的长度
	Sign []byte		//签名
}

func NewLeading(magic uint16, lenBlock uint32, lenSign uint16, sign []byte) *Leading {
	return &Leading{Magic: magic, LenBlock: lenBlock, LenSign: lenSign, Sign: sign}
}

func NewLeadingUnDump(data []byte) (*Leading, error) {
	structs := new(Structs)
	structs.Format = []string{"I", "H", "I", "104s"}
	result := structs.UnPack(data)
	magic := result[0].(uint16)
	lenBlock := result[1].(uint32)
	lenSign := result[2].(uint16)
	sign := []byte(result[3].(string))
	return &Leading{Magic: magic, LenBlock: lenBlock, LenSign: lenSign, Sign: sign}, nil
}

func (self *Leading) String() string{
	data, _ := json.Marshal(self)
	return fmt.Sprintf("%s", data)
}
