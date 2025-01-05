package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	files, err := os.ReadDir("my_directory")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("Direcroty:", file.Name())
		} else {
			fmt.Println("File:", file.Name())
		}
	}
}
