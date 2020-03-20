package message

import (
	"fmt"
	"testing"
)

func TestPrivKeys(t *testing.T) {
	msg := NewMsgAddr("addr", "pubk")
	fmt.Println(msg) //[123, 11, 34]

}