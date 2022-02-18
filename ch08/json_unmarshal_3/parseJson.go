package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// JSON 字符串
var JSON = `{
	"name":"Gopher",
	"title":"programmer",
	"contact":{
		"home":"415.333.3333",
		"cell":"415.555.1234"
	}
}`

func main() {
	// 将JSON字符串反序列化到map变量
	var c map[string]interface{}

	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println("Name:", c["name"])
	fmt.Println("Title:", c["title"])
	fmt.Println("Contact:")
	fmt.Println("\tH:", c["contact"].(map[string]interface{})["home"])
	fmt.Println("\tC:", c["contact"].(map[string]interface{})["cell"])

}
