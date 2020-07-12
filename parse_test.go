package main

import (
	"testing"
)

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
