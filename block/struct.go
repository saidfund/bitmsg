package block


//https://blog.csdn.net/qq_30638831/article/details/80421019

import (
	"bytes"
	"encoding/binary"
	"fmt"
	_ "os"
	"regexp"
	"strconv"
	"strings"
)

type Structs struct {
	Format []string
}


func (self *Structs) PackStrLen(src string) (num int) {
	reg := regexp.MustCompile(`^[1-9]\d*s$`)
	find := reg.FindString(src)
	if find != ""{
		find = strings.Trim(find, "s")
		num, err := strconv.Atoi("find")
		if err != nil{
			return 0
		}
		return num
	}
	return 0
}

//封包
func (self *Structs) Pack(args ...interface{}) []byte {
	lenArgs := len(args)
	lenFormat := len(self.Format)
	result := []byte{}
	if lenFormat > 0 && lenArgs > 0 && lenFormat == lenArgs {
		big := true
		for i := 0; i < lenFormat; i++ {
			if self.Format[i] == "H" {
				result = append(result, Uint16ToBytes(args[i].(uint16), big)...)
			} else if self.Format[i] == "I" {
				result = append(result, Uint32ToBytes(args[i].(uint32), big)...)
			} else if self.Format[i] == "q" {
				result = append(result, Int64ToBytes(args[i].(int64), big)...)
			} else if ok, _ := regexp.MatchString(`^[1-9]\d*s$`, self.Format[i]); ok {
				num, _ := strconv.Atoi(strings.TrimRight(self.Format[i], "s"))
				argstr := fmt.Sprintf("%s", args[i].(string))
				if len(argstr) < num{
					result = append(result, []byte(argstr)...)
					left := strings.Repeat("\x00", num - len(argstr))
					result = append(result, []byte(left)...)
				}else{
					result = append(result, []byte(argstr[:num])...)
				}
			}
		}
	}
	return result
}

//解包
func (self *Structs) UnPack(data []byte) []interface{} {
	if !self.IsSizeEqual(data) {
		return nil
	}
	lenFormat := len(self.Format)
	result := make([]interface{}, lenFormat)
	big := true
	if lenFormat > 0 {
		for i := 0; i < lenFormat; i++ {
			if self.Format[i] == "H" {
				result[i] = BytesTouint16(data[0:2], big)
				data = data[2:len(data)]
			} else if self.Format[i] == "I" {
				result[i] = BytesTouint32(data[0:4], big)
				data = data[4:len(data)]
			} else if self.Format[i] == "q" {
				result[i] = BytesToint64(data[0:8], big)
				data = data[8:len(data)]
			} else if ok, _ := regexp.MatchString(`^[1-9]\d*s$`, self.Format[i]); ok {
				num, _ := strconv.Atoi(strings.TrimRight(self.Format[i], "s"))
				result[i] = string(data[0:num])
				data = data[num:len(data)]
			}
		}
	}
	return result
}

func (self *Structs) IsSizeEqual(data []byte) bool {
	return self.Size() == len(data)
}

func (self *Structs) Size() int {
	size := 0
	ls := len(self.Format)
	if ls > 0 {
		for i := 0; i < ls; i++ {
			if self.Format[i] == "H" {
				size = size + 2
			} else if self.Format[i] == "I" {
				size = size + 4
			} else if self.Format[i] == "q" {
				size = size + 8
			} else if strings.Contains(self.Format[i], "s") {
				num, _ := strconv.Atoi(strings.TrimRight(self.Format[i], "s"))
				size = size + num
			}
		}
	}
	return size
}


//整数向字节转换
func Uint16ToBytes(num uint16, big bool) ([]byte){
	buf := new(bytes.Buffer)
	if big{
		binary.Write(buf, binary.BigEndian, &num)
	}else{
		binary.Write(buf, binary.LittleEndian, &num)
	}
	return buf.Bytes()
}

func Uint32ToBytes(num uint32, big bool) ([]byte){
	buf := new(bytes.Buffer)
	if big{
		binary.Write(buf, binary.BigEndian, &num)
	}else{
		binary.Write(buf, binary.LittleEndian, &num)
	}
	return buf.Bytes()
}

func Int16ToBytes(num int16, big bool) ([]byte){
	buf := new(bytes.Buffer)
	if big{
		binary.Write(buf, binary.BigEndian, &num)
	}else{
		binary.Write(buf, binary.LittleEndian, &num)
	}
	return buf.Bytes()
}

func Int32ToBytes(num int32, big bool) ([]byte){
	buf := new(bytes.Buffer)
	if big{
		binary.Write(buf, binary.BigEndian, &num)
	}else{
		binary.Write(buf, binary.LittleEndian, &num)
	}
	return buf.Bytes()
}

func Int64ToBytes(num int64, big bool) ([]byte){
	buf := new(bytes.Buffer)
	if big{
		binary.Write(buf, binary.BigEndian, &num)
	}else{
		binary.Write(buf, binary.LittleEndian, &num)
	}
	return buf.Bytes()
}


//字节向整数转换
func BytesTouint16(data []byte, big bool)(num uint16) {
	buf := bytes.NewBuffer(data)
	if big{
		binary.Read(buf, binary.BigEndian, &num)
	}else{
		binary.Read(buf, binary.BigEndian, &num)
	}
	return
}

func BytesTouint32(data []byte, big bool)(num uint32) {
	buf := bytes.NewBuffer(data)
	if big{
		binary.Read(buf, binary.BigEndian, &num)
	}else{
		binary.Read(buf, binary.BigEndian, &num)
	}
	return
}


func BytesToint64(data []byte, big bool)(num int64) {
	buf := bytes.NewBuffer(data)
	if big{
		binary.Read(buf, binary.BigEndian, &num)
	}else{
		binary.Read(buf, binary.BigEndian, &num)
	}
	return
}