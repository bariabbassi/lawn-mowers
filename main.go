package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	// Read file
	fileLines, err := readFile("./testfiles/default_file.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Parse lawn dimensions
	var lawn Lawn
	lawn.Width, lawn.Length, err = parseDimensions(fileLines[0])
	if err != nil {
		log.Fatal(err)
	}

	// Parse mowers
	var wg sync.WaitGroup
	wg.Add(len(fileLines) / 2)
	for i := 1; i < len(fileLines); i += 2 {
		go func(index int) {
			// Parse mower position
			x, y, orientation, err := parsePosition(fileLines[index], lawn)
			if err != nil {
				log.Fatal(err)
			}

			// Parse mower instructions
			instructions, err := parseInstructions(fileLines[index+1])
			if err != nil {
				log.Fatal(err)
			}

			// Add mower to the lawn
			lawn.addMower(Mower{x, y, orientation, instructions})
			wg.Done()
		}(i)
	}
	// Wait for all mowers to be added to the lawn
	wg.Wait()

	// Move mowers
	wg.Add(len(lawn.Mowers))
	for i := 0; i < len(lawn.Mowers); i++ {
		go func(index int) {
			err := lawn.moveMower(index)
			if err != nil {
				log.Fatal(err)
			}
			wg.Done()
		}(i)
	}
	// Wait for all instructions to be executed
	wg.Wait()

	// Display result (one-indexed)
	fmt.Print(lawn)
}
