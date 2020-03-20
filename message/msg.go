package message

import (
	"encoding/json"
	"fmt"
)

const FWRD uint8 = 0x01
const ADDR uint8 = 0x02

type Message struct {
	cmd string
}

type MsgAddr struct {
	Addr   string `json:"addr"`
	PubKey string `json:"pubk"`
}

func NewMsgAddr(addr string, pubKey string) *MsgAddr {
	return &MsgAddr{Addr: addr, PubKey: pubKey}
}

func (self *MsgAddr) String() string {
	data, _ := json.Marshal(self)
	return fmt.Sprintf("[%s]", data)
}

func (self *MsgAddr) Marshal() (string, error) {
	data, error := json.Marshal(self)
	return string(data), error
}

type MsgAddrAck struct {
	Addr   string `json:"addr"`
	PubKey string `json:"pubk"`
	RanSer string `json:"random"`
}

type MsgAddrAuth struct {
	RanSer string `json:"random"`
}

type MsgFwrd struct {
	Routing   string
	Block string
}




