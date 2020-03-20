package store

import (
	"fmt"
	"testing"
)

func TestPrivKeys(t *testing.T) {
	ok, err := IsExists("d:\\dump.txt")
	fmt.Println(ok)
	if err != nil{
		fmt.Println(err.Error())
	}

	err = CreateFileTrunc("d:\\dump.txt")
	if err != nil{
		fmt.Println(err.Error())
	}

}