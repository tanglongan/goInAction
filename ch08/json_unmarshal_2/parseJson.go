package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Address string `json:"address"`
	Age     int    `json:"age"`
	Name    string `json:"name"`
}

var jsonString = `{
    "address":"china",
    "age":23,
    "name":"程序猿"
}`

func main() {
	message := Message{}
	err := json.Unmarshal([]byte(jsonString), &message)
	if err != nil {
		fmt.Println("反序列化失败", err)
	}

	fmt.Printf("%+v\n", message)
	fmt.Println(message)
}
