package main

import (
	"fmt"
	"log"
)

type Lawn struct {
	Width  int
	Length int
}

type Mower struct {
	X            int
	Y            int
	Orientation  string
	Instructions []string
}

func main() {
	// Read file
	fileLines, err := readFile("./testfiles/default_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File lines: ", fileLines)

	// Get lawn dimensions
	var lawn Lawn
	lawn.Width, lawn.Length, err = parseDimentions(fileLines[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Lawn: ", lawn)

	// Get mowers
	var mowers []Mower
	for i := 1; i < len(fileLines); i += 2 {

		// Get position
		x, y, orientation, err := parsePosition(fileLines[i], lawn)
		if err != nil {
			log.Fatal(err)
		}

		// Get instructions
		instructions, err := parseInstructions(fileLines[i+1])
		if err != nil {
			log.Fatal(err)
		}

		mowers = append(mowers, Mower{x, y, orientation, instructions})
	}
	fmt.Println("Mowers: ", mowers)
}
