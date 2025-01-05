package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, err := os.Stat("/home/rendle/GolandProjects/github.com/rendley/headfirstgo/base/if_and_loop/file_info/my.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Size())
}
