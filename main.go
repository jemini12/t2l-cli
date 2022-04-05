package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

var HEROKU_URL string

func init() {
	HEROKU_URL = "https://t2l-project.herokuapp.com/"
}

func readStringWithRune() []rune {

	reader := bufio.NewReaderSize(os.Stdin, 40960)

	var output []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	return output

}

func main() {

	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	defer os.Stdin.Close()

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: fortune | ./t2l")
		return
	}

	output := readStringWithRune()
	data := make(map[string]interface{})
	data["raw"] = string(output)
	pbytes, _ := json.Marshal(data)
	buff := bytes.NewBuffer(pbytes)
	resp, err := http.Post(HEROKU_URL+"documents", "application/json", buff)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(string(output))
	// Response 체크.
	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println("T2L >> " + str)
	}
}
