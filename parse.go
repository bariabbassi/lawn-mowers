package main

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

func readFile(filePath string) ([]string, error) {
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	if len(fileData) < 11 {
		return nil, errors.New("File contains too litle data")
	}

	fileLines := strings.Split(string(fileData), "\n")
	if len(fileData) < 3 {
		return nil, errors.New("File should contain at least 3 lines")
	}

	// The number of lines can't be even:
	// 1st line for lawn dimensions and every mower needs 2 lines (position & instructions)
	if len(fileLines)%2 == 0 {
		return nil, errors.New("File lines number is even")
	}

	return fileLines, nil
}

func parseDimensions(firstLine string) (int, int, error) {
	dimensions := strings.Split(firstLine, " ")

	if len(dimensions) != 2 {
		return 0, 0, errors.New("First line should contain 2 numbers: width and length")
	}

	width, err := strconv.Atoi(dimensions[0])
	if err != nil {
		return 0, 0, err
	}
	if width < 1 {
		return 0, 0, errors.New("Lawn width is out of range")
	}

	length, err := strconv.Atoi(dimensions[1])
	if err != nil {
		return 0, 0, err
	}
	if length < 1 {
		return 0, 0, errors.New("Lawn length is out of range")
	}

	return width, length, nil
}

func parsePosition(fileLine string, lawn Lawn) (int, int, string, error) {
	position := strings.Split(fileLine, " ")

	if len(position) != 3 {
		return 0, 0, "", errors.New("Position line should contain 3 characters: X, Y, and O")
	}

	x, err := strconv.Atoi(position[0])
	if err != nil {
		return 0, 0, "", err
	}

	if x < 1 || x > lawn.Width {
		return 0, 0, "", errors.New("X coordinate out of bound")
	}

	y, err := strconv.Atoi(position[1])
	if err != nil {
		return 0, 0, "", err
	}

	if y < 1 || y > lawn.Length {
		return 0, 0, "", errors.New("Y coordinate out of bound")
	}

	orientation := position[2]
	if len(orientation) != 1 {
		return 0, 0, "", errors.New("O can't be longer than 1 lettre. O should be N, S, E, or W")
	}

	if orientation != "N" && orientation != "S" && orientation != "E" && orientation != "W" {
		return 0, 0, "", errors.New("O is not equal to one of the 4 orientations: N, S, E, or W")
	}

	// the file is 1 indexed, our program is 0 indexed
	return x - 1, y - 1, orientation, nil
}

func parseInstructions(fileLine string) ([]string, error) {
	var instructions []string
	for _, char := range fileLine {
		if char == 'L' || char == 'F' || char == 'R' {
			instructions = append(instructions, string(char))
		} else {
			return nil, errors.New("Instructions can only be L, F, or R")
		}
	}
	return instructions, nil
}
