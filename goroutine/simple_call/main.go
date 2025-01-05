package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	response, err := http.Get("https://example.com")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
