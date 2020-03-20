package main

import (
	"encoding/json"
	"fmt"
)

type Users struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	//结构体解析成json
	user1 := Users{"1", "user1", 22}
	s, err := json.Marshal(user1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(s))
	}
	//json解析到结构体
	var user2 Users
	err = json.Unmarshal(s, &user2)
	if err == nil {
		fmt.Println(user2)
	}
	//map解析成json
	m := make(map[string]interface{}, 2)
	m["id"] = "a"
	m["name"] = "bb"
	var data []byte
	if data, err = json.Marshal(m); err == nil {
		fmt.Println(string(data))
	}
	//json解析成map
	if err = json.Unmarshal(data, &m); err == nil {
		fmt.Println(m)
	}

}
