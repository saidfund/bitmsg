package main

import (
	"encoding/json"
	"fmt"
)

type MessageFwrd struct {
	Addr string		`json:"addr"`
	PubKey string	`json:"pubk"`
}

func NewMessageFwrd(addr string, pubKey string) *MessageFwrd {
	return &MessageFwrd{Addr: addr, PubKey: pubKey}
}

func (self *MessageFwrd) String() string{
	data, _ := json.Marshal(self)
	return fmt.Sprintf("[%s]", data)
}

func main() {
	msg := NewMessageFwrd("addr", "pubk")
	fmt.Println(msg) //[123, 11, 34]
}