package cell

import (
	"encoding/json"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
	"reflect"
)

const (
	HAND int = 0x10
	SBLK int = 0x11
	RBLK int = 0x12
	SACK int = 0x13
	OVER int = 0xFF
)

/////////////////////////////////////////////////////////////////////////////////////
type HandMsg struct {
	Addr string `json:"addr"`
	Info string `json:"info"`
}

func (self *HandMsg) String() string {
	buf, _ := json.Marshal(self)
	return string(buf)
}

/////////////////////////////////////////////////////////////////////////////////////
type SblkMsg struct {
	Mkt  string `json:"mkt"`
	Info string `json:"info"`
}

func (self *SblkMsg) String() string {
	buf, _ := json.Marshal(self)
	return string(buf)
}

/////////////////////////////////////////////////////////////////////////////////////
type RblkMsg struct {
	Mkt  string `json:"mkt"`
	Info string `json:"info"`
}

func (self *RblkMsg) String() string {
	buf, _ := json.Marshal(self)
	return string(buf)
}

/////////////////////////////////////////////////////////////////////////////////////
type OverMsg struct {
	Mkt  string `json:"mkt"`
	Info string `json:"info"`
}

func (self *OverMsg) String() string {
	buf, _ := json.Marshal(self)
	return string(buf)
}

/////////////////////////////////////////////////////////////////////////////////////
type SackMsg struct {
	Code string `json:"code"`
	Info string `json:"info"`
}

func (self *SackMsg) String() string {
	buf, _ := json.Marshal(self)
	return string(buf)
}

// 将消息注册到系统
func init() {
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("binary"),
		Type:  reflect.TypeOf((*HandMsg)(nil)).Elem(),
		ID:    HAND,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("binary"),
		Type:  reflect.TypeOf((*SblkMsg)(nil)).Elem(),
		ID:    SBLK,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("binary"),
		Type:  reflect.TypeOf((*OverMsg)(nil)).Elem(),
		ID:    OVER,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("binary"),
		Type:  reflect.TypeOf((*SackMsg)(nil)).Elem(),
		ID:    SACK,
	})
}
