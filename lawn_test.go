package main

import (
	"testing"
)

/////////////////////// Add Move ///////////////////////

func TestAddMower(t *testing.T) {
	expectedMower := Mower{0, 1, "N", nil}

	lawn := Lawn{5, 5, nil}
	lawn.addMower(Mower{0, 1, "N", nil})
	if len(lawn.Mowers) != 1 {
		t.Error("Mower was not added")
	}
	if lawn.Mowers[0].Orientation != expectedMower.Orientation {
		t.Error("Orientation is not correct")
	}
	if lawn.Mowers[0].X != expectedMower.X {
		t.Error("X is not correct")
	}
	if lawn.Mowers[0].Y != expectedMower.Y {
		t.Error("Y is not correct")
	}
}

/////////////////////// Move Mower Forward ///////////////////////

func TestMoveMowerForward1(t *testing.T) {
	expectedLawn := Lawn{5, 5, nil}
	expectedLawn.addMower(Mower{0, 2, "N", nil})

	lawn := Lawn{5, 5, nil}
	lawn.addMower(Mower{0, 1, "N", []string{"L", "F", "L", "F", "L", "F", "L", "F", "F"}})
	err := lawn.moveMowerForward(0)
	if err != nil {
		t.Error(err)
	}
	if len(lawn.Mowers) != len(expectedLawn.Mowers) {
		t.Error("Mowers slice was modified")
	}
	if lawn.Mowers[0].Orientation != expectedLawn.Mowers[0].Orientation {
		t.Error("Orientation shouldn't change while moving forward")
	}
	if lawn.Mowers[0].X != expectedLawn.Mowers[0].X {
		t.Error("X is not correct")
	}
	if lawn.Mowers[0].Y != expectedLawn.Mowers[0].Y {
		t.Error("Y is not correct")
	}
}

func TestMoveMowerForward2(t *testing.T) {
	expectedLawn := Lawn{5, 5, nil}
	expectedLawn.addMower(Mower{0, 1, "W", nil})

	lawn := Lawn{5, 5, nil}
	lawn.addMower(Mower{0, 1, "W", []string{"L", "F", "L", "F", "L", "F", "L", "F", "F"}})
	err := lawn.moveMowerForward(0)
	if err != nil {
		t.Error(err)
	}
	if len(lawn.Mowers) != len(expectedLawn.Mowers) {
		t.Error("Mowers slice was modified")
	}
	if lawn.Mowers[0].Orientation != expectedLawn.Mowers[0].Orientation {
		t.Error("Orientation shouldn't change while moving forward")
	}
	if lawn.Mowers[0].X != expectedLawn.Mowers[0].X {
		t.Error("X is not correct")
	}
	if lawn.Mowers[0].Y != expectedLawn.Mowers[0].Y {
		t.Error("Y is not correct")
	}
}

func TestMoveMowerForward3(t *testing.T) {
	expectedLawn := Lawn{5, 5, nil}
	expectedLawn.addMower(Mower{4, 4, "E", nil})

	lawn := Lawn{5, 5, nil}
	lawn.addMower(Mower{4, 4, "E", []string{"L", "F", "L", "F", "L", "F", "L", "F", "F"}})
	err := lawn.moveMowerForward(0)
	if err != nil {
		t.Error(err)
	}
	if len(lawn.Mowers) != len(expectedLawn.Mowers) {
		t.Error("Mowers slice was modified")
	}
	if lawn.Mowers[0].Orientation != expectedLawn.Mowers[0].Orientation {
		t.Error("Orientation shouldn't change while moving forward")
	}
	if lawn.Mowers[0].X != expectedLawn.Mowers[0].X {
		t.Error("X is not correct")
	}
	if lawn.Mowers[0].Y != expectedLawn.Mowers[0].Y {
		t.Error("Y is not correct")
	}
}

/////////////////////// Move Mower ///////////////////////

func TestMoveMower(t *testing.T) {
	expectedLawn := Lawn{5, 5, nil}
	expectedLawn.addMower(Mower{1, 2, "N", []string{}})

	lawn := Lawn{5, 5, nil}
	lawn.addMower(Mower{0, 1, "N", []string{"L", "F", "L", "F", "L", "F", "L", "F", "F"}})
	err := lawn.moveMower(0)
	if err != nil {
		t.Error(err)
	}
	if len(lawn.Mowers) != len(expectedLawn.Mowers) {
		t.Error("Mowers slice was modified")
	}
	if lawn.Mowers[0].Orientation != expectedLawn.Mowers[0].Orientation {
		t.Error("Mower orientation was not turned to the left correctly")
	}
	if lawn.Mowers[0].X != expectedLawn.Mowers[0].X {
		t.Error("X shouldn't change while turning")
	}
	if lawn.Mowers[0].Y != expectedLawn.Mowers[0].Y {
		t.Error("Y shouldn't change while turning")
	}
}

func TestMoveMower2(t *testing.T) {
	expectedLawn := Lawn{5, 5, nil}
	expectedLawn.addMower(Mower{4, 0, "E", []string{}})

	lawn := Lawn{5, 5, nil}
	lawn.addMower(Mower{2, 2, "E", []string{"F", "F", "R", "F", "F", "R", "F", "R", "R", "F"}})
	err := lawn.moveMower(0)
	if err != nil {
		t.Error(err)
	}
	if len(lawn.Mowers) != len(expectedLawn.Mowers) {
		t.Error("Mowers slice was modified")
	}
	if lawn.Mowers[0].Orientation != expectedLawn.Mowers[0].Orientation {
		t.Error("Mower orientation was not turned to the left correctly")
	}
	if lawn.Mowers[0].X != expectedLawn.Mowers[0].X {
		t.Error("X shouldn't change while turning")
	}
	if lawn.Mowers[0].Y != expectedLawn.Mowers[0].Y {
		t.Error("Y shouldn't change while turning", lawn.Mowers[0])
	}
}

func TestMoveMower3(t *testing.T) {
	expectedLawn := Lawn{5, 5, nil}
	expectedLawn.addMower(Mower{2, 3, "S", []string{}})

	lawn := Lawn{5, 5, nil}
	lawn.addMower(Mower{2, 3, "S", []string{}})
	err := lawn.moveMower(0)
	if err != nil {
		t.Error(err)
	}
	if len(lawn.Mowers) != len(expectedLawn.Mowers) {
		t.Error("Mowers slice was modified")
	}
	if lawn.Mowers[0].Orientation != expectedLawn.Mowers[0].Orientation {
		t.Error("Mower orientation was not turned to the left correctly")
	}
	if lawn.Mowers[0].X != expectedLawn.Mowers[0].X {
		t.Error("X shouldn't change while turning")
	}
	if lawn.Mowers[0].Y != expectedLawn.Mowers[0].Y {
		t.Error("Y shouldn't change while turning", lawn.Mowers[0])
	}
}
