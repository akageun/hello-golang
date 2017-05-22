package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
)

func main() {
	var url string = "http://blog.geun.kr";

	resp, err := http.Get(url) // GET
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	// 결과 출력
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%s\n", string(data))
}
