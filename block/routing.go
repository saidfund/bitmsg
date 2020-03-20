package block

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

const ADDRLEN int = 34
const SUBLEN int = 42


type Routing struct {
	AddrLen uint16		//地址长度，最长34
	Addr []byte			//地址
	SubLen uint16		//子地址长度, 最长42
	Sub  []byte			//子地址
}

func NewRoutingUnDump(data []byte) *Routing{
	structs := new(Structs)
	structs.Format = []string{"H", "34s", "H", "42s"}
	result := structs.UnPack(data)
	addLen := result[0].(uint16)
	addr :=[]byte(result[1].(string))
	subLen := result[2].(uint16)
	sub :=[]byte(result[3].(string))
	return &Routing{AddrLen: addLen, Addr: addr, SubLen: subLen, Sub: sub}
}


func NewRouting(addr []byte, sub []byte) *Routing {
	routing := &Routing{Addr: addr, Sub: sub}

	routing.AddrLen = uint16(len(addr))
	routing.SubLen = uint16(len(sub))

	left := bytes.Repeat([]byte{0x00}, ADDRLEN - len(addr))
	addr = bytes.Join([][]byte{addr, left}, []byte{})
	left = bytes.Repeat([]byte{0x00}, SUBLEN - len(sub))
	sub = bytes.Join([][]byte{sub, left}, []byte{})

	routing.Addr = addr
	routing.Sub = sub
	return routing
}

func (self *Routing) Address() (string) {
	return string(self.Addr[:self.AddrLen])
}

func (self *Routing) Subress() (string) {
	return string(self.Sub[:self.SubLen])
}

func (self *Routing) Dump() ([]byte, error) {
	bigEndian := true
	buf := bytes.NewBuffer([]byte{})
	buf.Write(Uint16ToBytes(self.AddrLen, bigEndian))
	buf.Write(self.Addr)
	buf.Write(Uint16ToBytes(self.SubLen, bigEndian))
	buf.Write(self.Sub)
	return buf.Bytes(), nil
}



func (self *Routing) String() string {
	data, _ := json.Marshal(self)
	return fmt.Sprintf("%s", data)
}

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

