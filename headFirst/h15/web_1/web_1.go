package main

import (
	"log"
	"net/http"
)

// viewHandler 处理请求
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("hello Web!")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/hello", viewHandler)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
