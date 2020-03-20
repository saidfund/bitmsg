package block

import (
	"bytes"
	"encoding/binary"
)


type Bytes []byte

func (self Bytes) ToUint16()(num uint16, size int) {
	size = 2
	buf := bytes.NewBuffer(self)
	binary.Read(buf, binary.BigEndian, &num)
	return
}

func FromUint16(num uint16) (Bytes){
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, &num)
	return buf.Bytes()
}

func (self Bytes) ToUint32()(num uint32, size int) {
	size = 4
	buf := bytes.NewBuffer(self)
	binary.Read(buf, binary.BigEndian, &num)
	return
}

func FromUint32(num uint32) (Bytes){
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, &num)
	return buf.Bytes()
}

func (self Bytes) ToInt64()(num int64, size int) {
	size = 8
	buf := bytes.NewBuffer(self)
	binary.Read(buf, binary.BigEndian, &num)
	return
}

func FromInt64(num int64) (Bytes){
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, &num)
	return buf.Bytes()
}

