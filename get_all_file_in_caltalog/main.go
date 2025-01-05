package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func reportPanic() {
	p := recover()
	if p == nil {
		return
	}
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		panic(p)
	}
}

func scanDirectory(path string) error {
	fmt.Println(path)
	files, err := os.ReadDir(path)
	if err != nil {
		//fmt.Printf("Returning error from scanDirecroty(\"%s\") call\n", path)
		//return err
		panic(err)
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			err := scanDirectory(filePath)
			if err != nil {
				fmt.Printf("Returning error from scanDirecroty(\"%s\") call\n", path)
				return err
			}
		} else {
			fmt.Println(filePath)
			panic("oh, no, we get a file!")
		}
	}
	return nil
}

func main() {
	//err := scanDirectory("headfirstgo")
	//if err != nil {
	//	log.Fatal(err)
	//}
	defer reportPanic()
	scanDirectory("headfirstgo")
}
