package chunk

import (
	"encoding/json"
	"fmt"
	"github.com/fananchong/cstruct-go"
	"strings"
)

const ADDRLEN int = 34
const SUBLEN int = 42

type Routing struct {
	Caller [48]byte
	Callee [48]byte
}

func NewRouting(caller *Quater, callee *Quater) *Routing {
	cr := caller.ToBytes()
	ce := callee.ToBytes()
	return &Routing{Caller: cr, Callee: ce}
}

func RoutingUnMarshal(data []byte) (self *Routing) {
	self = &Routing{}
	cstruct.Unmarshal(data, self)
	return self
}

func (self *Routing) Marshal() []byte {
	data, _ := cstruct.Marshal(self)
	//fmt.Println(hex.Dump(data))
	return data
}

func (self *Routing) String() string {
	caller := &Quater{}
	caller.FromStr(string(self.Caller[:]))
	cr, _ := json.Marshal(caller)
	callee := &Quater{}
	callee.FromStr(string(self.Callee[:]))
	ce, _ := json.Marshal(callee)
	return fmt.Sprintf("Caller :%s,Callee :%s", cr, ce)
}

type Quater struct {
	Main string
	Sub  string
}

func (self *Quater) FromStr(data string) {
	result := strings.Split(data, ".")
	if len(result) != 2 {
		return
	}
	self.Main = result[0]
	self.Sub = strings.Trim(result[1], string(0x0))
	return
}

func (self *Quater) ToBytes() (result [48]byte) {
	copy(result[:], []byte(self.Main+"."+self.Sub))
	return
}
