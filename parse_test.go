package main

import (
	"testing"
)

/////////////////////// Read File ///////////////////////

func TestReadDefaultFile(t *testing.T) {
	expected := []string{"5 5", "1 2 N", "LFLFLFLFF", "3 3 E", "FFRFFRFRRF"}

	fileLines, err := readFile("./testfiles/default_file.txt")
	if err != nil {
		t.Error(err)
	}
	if len(fileLines) != len(expected) {
		t.Error("File was not read intierly")
	}
	for i, line := range fileLines {
		if line != expected[i] {
			t.Error("File read incorrectly")
		}
	}
}

func TestReadEmptyFile(t *testing.T) {
	fileLines, err := readFile("./testfiles/empty_file.txt")
	if err.Error() != "File contains too litle data" {
		t.Error(err)
	}
	_ = fileLines
}

func TestRead1LineFile(t *testing.T) {
	fileLines, err := readFile("./testfiles/1_line_file.txt")
	if err.Error() != "File contains too litle data" {
		t.Error(err)
	}
	_ = fileLines
}

func TestRead2LineFile(t *testing.T) {
	fileLines, err := readFile("./testfiles/2_lines_file.txt")
	if err.Error() != "File contains too litle data" {
		t.Error(err)
	}
	_ = fileLines
}

func TestRead3LineFile(t *testing.T) {
	expected := []string{"5 5", "1 2 N", "L"}

	fileLines, err := readFile("./testfiles/3_lines_file.txt")
	if err != nil {
		t.Error(err)
	}
	if len(fileLines) != len(expected) {
		t.Error("File was not read intierly")
	}
	for i, line := range fileLines {
		if line != expected[i] {
			t.Error("File read incorrectly")
		}
	}
}

func TestRead4LineFile(t *testing.T) {
	fileLines, err := readFile("./testfiles/4_lines_file.txt")
	if err.Error() != "File lines number is even" {
		t.Error(err)
	}
	_ = fileLines
}

/////////////////////// Parse Dimensions ///////////////////////

func TestParseDefaultDimensions1(t *testing.T) {
	expectedWidth, expectedLength := 25, 6

	width, length, err := parseDimensions("25 6")
	if err != nil {
		t.Error(err)
	}
	if width != expectedWidth || length != expectedLength {
		t.Error("Lawn dimensions parsed incorrectly")
	}
}

func TestParseDefaultDimensions2(t *testing.T) {
	expectedWidth, expectedLength := 2019, 2020

	width, length, err := parseDimensions("2019 2020")
	if err != nil {
		t.Error(err)
	}
	if width != expectedWidth || length != expectedLength {
		t.Error("Lawn dimensions parsed incorrectly")
	}
}

func TestParseEmptyDimensions(t *testing.T) {
	width, length, err := parseDimensions("")
	if err.Error() != "First line should contain 2 numbers: width and length" {
		t.Error(err)
	}
	_, _ = width, length
}

func TestParse1Dimensions(t *testing.T) {
	width, length, err := parseDimensions("10")
	if err.Error() != "First line should contain 2 numbers: width and length" {
		t.Error(err)
	}
	_, _ = width, length
}

func TestParse3Dimensions(t *testing.T) {
	width, length, err := parseDimensions("")
	if err.Error() != "First line should contain 2 numbers: width and length" {
		t.Error(err)
	}
	_, _ = width, length
}

func TestParseAlphaDimensions(t *testing.T) {
	width, length, err := parseDimensions("27 a")
	if err.Error() != `strconv.Atoi: parsing "a": invalid syntax` {
		t.Error(err)
	}
	_, _ = width, length
}

func TestParseSpacedDimensions(t *testing.T) {
	width, length, err := parseDimensions("36  5")
	if err.Error() != "First line should contain 2 numbers: width and length" {
		t.Error(err)
	}
	_, _ = width, length
}

func TestParseZeroDimensions(t *testing.T) {
	width, length, err := parseDimensions("0 98")
	if err.Error() != "Lawn width is out of range" {
		t.Error(err)
	}
	_, _ = width, length
}

func TestParseNegativeDimensions(t *testing.T) {
	width, length, err := parseDimensions("40 -2")
	if err.Error() != "Lawn length is out of range" {
		t.Error(err)
	}
	_, _ = width, length
}

func TestParseOverFlowDimensions(t *testing.T) {
	width, length, err := parseDimensions("1000000000000000000000 7")
	if err.Error() != `strconv.Atoi: parsing "1000000000000000000000": value out of range` {
		t.Error(err)
	}
	_, _ = width, length
}

/////////////////////// Parse Position ///////////////////////

func TestParseDefaultPosition(t *testing.T) {
	expectedX, expectedY, expectedO := 0, 1, "N"

	x, y, o, err := parsePosition("1 2 N", Lawn{5, 5, nil})
	if err != nil {
		t.Error(err)
	}
	if x != expectedX || y != expectedY || o != expectedO {
		t.Error("Lawn dimensions parsed incorrectly", x, y, o)
	}
}

func TestParseDefaultPosition2(t *testing.T) {
	expectedX, expectedY, expectedO := 99, 0, "S"

	x, y, o, err := parsePosition("100 1 S", Lawn{100, 1, nil})
	if err != nil {
		t.Error(err)
	}
	if x != expectedX || y != expectedY || o != expectedO {
		t.Error("Lawn dimensions parsed incorrectly", x, y, o)
	}
}

func TestParse4FieldsPosition(t *testing.T) {
	x, y, o, err := parsePosition("34 @ 7 E", Lawn{100, 1, nil})
	if err.Error() != "Position line should contain 3 characters: X, Y, and O" {
		t.Error(err)
	}
	_, _, _ = x, y, o
}

func TestParseZeroPosition(t *testing.T) {
	x, y, o, err := parsePosition("3 0 N", Lawn{5, 5, nil})
	if err.Error() != "Y coordinate out of bound" {
		t.Error(err)
	}
	_, _, _ = x, y, o
}

func TestParseBoundPosition(t *testing.T) {
	x, y, o, err := parsePosition("3 6 W", Lawn{2, 15, nil})
	if err.Error() != "X coordinate out of bound" {
		t.Error(err)
	}
	_, _, _ = x, y, o
}

/////////////////////// Parse Instructions ///////////////////////

func TestParseDefaultInstructions(t *testing.T) {
	expected := []string{"L", "F", "L", "F", "L", "F", "L", "F", "F"}

	instructions, err := parseInstructions("LFLFLFLFF")
	if err != nil {
		t.Error(err)
	}
	for i, instruction := range instructions {
		if instruction != expected[i] {
			t.Error("File read incorrectly")
		}
	}
}

func TestParseSpacedInstructions(t *testing.T) {
	instructions, err := parseInstructions("LF R")
	if err.Error() != "Instructions can only be L, F, or R" {
		t.Error(err)
	}
	_ = instructions
}

func TestParseWrongInstructions(t *testing.T) {
	instructions, err := parseInstructions("RzzF")
	if err.Error() != "Instructions can only be L, F, or R" {
		t.Error(err)
	}
	_ = instructions
}
