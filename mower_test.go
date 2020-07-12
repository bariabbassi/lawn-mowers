package main

import (
	"testing"
)

/////////////////////// Turn Left ///////////////////////

func TestTurnLeftN(t *testing.T) {
	expectedMower := Mower{0, 1, "W", nil}

	mower := Mower{0, 1, "N", nil}
	mower.turnLeft()
	if mower.Orientation != expectedMower.Orientation {
		t.Error("Mower orientation was not turned to the left correctly")
	}
	if mower.X != expectedMower.X {
		t.Error("X shouldn't change while turning")
	}
	if mower.Y != expectedMower.Y {
		t.Error("Y shouldn't change while turning")
	}
	if len(mower.Instructions) != len(expectedMower.Instructions) {
		t.Error("Instructions shouldn't change while turning")
	}
}

func TestTurnLeftE(t *testing.T) {
	expectedMower := Mower{0, 1, "N", nil}

	mower := Mower{0, 1, "E", nil}
	mower.turnLeft()
	if mower.Orientation != expectedMower.Orientation {
		t.Error("Mower orientation was not turned to the left correctly")
	}
	if mower.X != expectedMower.X {
		t.Error("X shouldn't change while turning")
	}
	if mower.Y != expectedMower.Y {
		t.Error("Y shouldn't change while turning")
	}
	if len(mower.Instructions) != len(expectedMower.Instructions) {
		t.Error("Instructions shouldn't change while turning")
	}
}

func TestTurnLeftW(t *testing.T) {
	expectedMower := Mower{0, 1, "S", nil}

	mower := Mower{0, 1, "W", nil}
	mower.turnLeft()
	if mower.Orientation != expectedMower.Orientation {
		t.Error("Mower orientation was not turned to the left correctly")
	}
	if mower.X != expectedMower.X {
		t.Error("X shouldn't change while turning")
	}
	if mower.Y != expectedMower.Y {
		t.Error("Y shouldn't change while turning")
	}
	if len(mower.Instructions) != len(expectedMower.Instructions) {
		t.Error("Instructions shouldn't change while turning")
	}
}

func TestTurnLeftS(t *testing.T) {
	expectedMower := Mower{0, 1, "E", nil}

	mower := Mower{0, 1, "S", nil}
	mower.turnLeft()
	if mower.Orientation != expectedMower.Orientation {
		t.Error("Mower orientation was not turned to the left correctly")
	}
	if mower.X != expectedMower.X {
		t.Error("X shouldn't change while turning")
	}
	if mower.Y != expectedMower.Y {
		t.Error("Y shouldn't change while turning")
	}
	if len(mower.Instructions) != len(expectedMower.Instructions) {
		t.Error("Instructions shouldn't change while turning")
	}
}

/////////////////////// Turn Right ///////////////////////

func TestTurnRightN(t *testing.T) {
	expectedMower := Mower{0, 1, "E", nil}

	mower := Mower{0, 1, "N", nil}
	mower.turnRight()
	if mower.Orientation != expectedMower.Orientation {
		t.Error("Mower orientation was not turned to the right correctly")
	}
	if mower.X != expectedMower.X {
		t.Error("X shouldn't change while turning")
	}
	if mower.Y != expectedMower.Y {
		t.Error("Y shouldn't change while turning")
	}
	if len(mower.Instructions) != len(expectedMower.Instructions) {
		t.Error("Instructions shouldn't change while turning")
	}
}

func TestTurnRightE(t *testing.T) {
	expectedMower := Mower{0, 1, "S", nil}

	mower := Mower{0, 1, "E", nil}
	mower.turnRight()
	if mower.Orientation != expectedMower.Orientation {
		t.Error("Mower orientation was not turned to the right correctly")
	}
	if mower.X != expectedMower.X {
		t.Error("X shouldn't change while turning")
	}
	if mower.Y != expectedMower.Y {
		t.Error("Y shouldn't change while turning")
	}
	if len(mower.Instructions) != len(expectedMower.Instructions) {
		t.Error("Instructions shouldn't change while turning")
	}
}

func TestTurnRightW(t *testing.T) {
	expectedMower := Mower{0, 1, "N", nil}

	mower := Mower{0, 1, "W", nil}
	mower.turnRight()
	if mower.Orientation != expectedMower.Orientation {
		t.Error("Mower orientation was not turned to the right correctly")
	}
	if mower.X != expectedMower.X {
		t.Error("X shouldn't change while turning")
	}
	if mower.Y != expectedMower.Y {
		t.Error("Y shouldn't change while turning")
	}
	if len(mower.Instructions) != len(expectedMower.Instructions) {
		t.Error("Instructions shouldn't change while turning")
	}
}

func TestTurnRightS(t *testing.T) {
	expectedMower := Mower{0, 1, "W", nil}

	mower := Mower{0, 1, "S", nil}
	mower.turnRight()
	if mower.Orientation != expectedMower.Orientation {
		t.Error("Mower orientation was not turned to the right correctly")
	}
	if mower.X != expectedMower.X {
		t.Error("X shouldn't change while turning")
	}
	if mower.Y != expectedMower.Y {
		t.Error("Y shouldn't change while turning")
	}
	if len(mower.Instructions) != len(expectedMower.Instructions) {
		t.Error("Instructions shouldn't change while turning")
	}
}
