package main

import (
	"fmt"
	"log"
)

type Lawn struct {
	Width  int
	Length int
}

func main() {
	// Read file
	fileLines, err := readFile("./testfiles/default_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File lines: ", fileLines)

	// Parse lawn dimensions
	var lawn Lawn
	lawn.Width, lawn.Length, err = parseDimensions(fileLines[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Lawn: ", lawn)

	// Parse mowers
	var mowers []Mower
	for i := 1; i < len(fileLines); i += 2 {

		// Parse position
		x, y, orientation, err := parsePosition(fileLines[i], lawn)
		if err != nil {
			log.Fatal(err)
		}

		// Parse instructions
		instructions, err := parseInstructions(fileLines[i+1])
		if err != nil {
			log.Fatal(err)
		}

		mower := Mower{x, y, orientation, instructions}
		mower.move(lawn)
		mowers = append(mowers, mower)
	}
	fmt.Println("Mowers: ", mowers)
}
