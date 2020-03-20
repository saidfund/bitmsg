package block

type Content struct {
	TypeOf   uint32 //单元内容类型
	HashLeaf []byte //单元哈希值
	Cipher   []byte //单元密码，为空则表示未加密
	Len      uint32 //单元内容长度
	Body     []byte //内容
}

func NewCell(typeOf uint32, hashLeaf []byte, cipher []byte, len uint32, body []byte) *Content {
	return &Content{TypeOf: typeOf, HashLeaf: hashLeaf, Cipher: cipher, Len: len, Body: body}
}
