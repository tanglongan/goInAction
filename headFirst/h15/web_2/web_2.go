package main

import (
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func englishHandler(writer http.ResponseWriter, resuest *http.Request) {
	write(writer, "english")
}

func frenchHandler(writer http.ResponseWriter, resuest *http.Request) {
	write(writer, "french")
}

func chineseHandler(writer http.ResponseWriter, resuest *http.Request) {
	write(writer, "chinese")
}

func main() {
	http.HandleFunc("/english", englishHandler)
	http.HandleFunc("/french", frenchHandler)
	http.HandleFunc("/chinese", chineseHandler)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
