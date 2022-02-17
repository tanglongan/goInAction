package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type (
	gResult struct {
		GsearchResultClass string `json:"GsearchResultClass"`
		UnescapedUrl       string `json:"unescapedUrl"`
		URL                string `json:"url"`
		VisibleURL         string `json:"visibleUrl"`
		CacheURL           string `json:"cacheURL"`
		Title              string `json:"title"`
		TitleNoFormatting  string `json:"titleNoFormatting"`
		Content            string `json:"content"`
	}

	gResponse struct {
		ResponseData struct {
			Results []gResult `json:"results"`
		} `json:"responseData"`
	}
)

// OpenFile 打开文件
func OpenFile(filename string) (*os.File, error) {
	fmt.Println("Opening", filename)
	return os.Open(filename)
}

// CloseFile 关闭文件
func CloseFile(file *os.File) {
	fmt.Println("Closing file")
	file.Close()
}

func ReadFile(filepath string) (string, error) {
	file, err := OpenFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer CloseFile(file)

	content := make([]string, 1000)
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else {
			content = append(content, str)
		}
	}

	return strings.Join(content, ""), nil
}

func main() {
	content, err := ReadFile("D:\\Code\\goInAction\\ch08\\json_unmarshal_1\\data.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(content)

	response := gResponse{}
	err = json.Unmarshal([]byte(content), &response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)

}
