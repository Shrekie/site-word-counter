package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	resp, err := http.Get("http://mozilla.org")
	if err != nil {
		log.Fatal(err)
		// handle errors?
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		// handle errors?
	}
	wordCount := strings.Count(string(body), "firefox")
	fmt.Printf("%d \n", wordCount)
}
