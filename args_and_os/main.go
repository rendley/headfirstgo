package main

import (
	"fmt"
	"os"
)

func main() {
	binaryFileName := os.Args[0]
	firstFlag := os.Args[1]
	strPath := os.Args[2]
	// go run guestbook.go -dir /opt/file/format
	fmt.Println(binaryFileName)
	fmt.Println(firstFlag)
	fmt.Println(strPath) // /opt/file/format
}

// ./main -dir /etc/dir/format
