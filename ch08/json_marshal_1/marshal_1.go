package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	c := make(map[string]interface{})
	c["name"] = "Gopher"
	c["title"] = "programmer"
	c["contact"] = map[string]interface{}{
		"home": "415.123.12345",
		"cell": "435.577.75655",
	}

	// MarshalIndent 带有缩进格式，适合开发人员查看
	// Marshal 没有缩进，适合网络传输
	//data, err := json.MarshalIndent(c, "", "  ")
	data, err := json.Marshal(c)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	fmt.Println(string(data))

}
