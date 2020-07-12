package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

const (
	filePath = "./file_short.txt"
)

func main() {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	if len(fileBytes) < 8 {
		log.Fatalf("File %s is too short", filePath)
		return
	}
	fmt.Println(len(fileBytes))
	fmt.Println(string(fileBytes))
}
