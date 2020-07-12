package main

import (
	"errors"
)

type Lawn struct {
	Width  int
	Length int
	Mowers []Mower
}

func (lawn *Lawn) addMower(mower Mower) {
	lawn.Mowers = append(lawn.Mowers, mower)
}

func (lawn *Lawn) moveMowerForward(index int) error {
	if index < 0 || index > len(lawn.Mowers)-1 {
		return errors.New("This mower doesn't exist in the lawn")
	}
	mower := lawn.Mowers[index]
	// Find next position
	switch mower.Orientation {
	case "N":
		if mower.Y+1 < lawn.Length {
			mower.Y++
		}
	case "E":
		if mower.X+1 < lawn.Width {
			mower.X++
		}
	case "W":
		if mower.X-1 >= 0 {
			mower.X--
		}
	case "S":
		if mower.Y-1 >= 0 {
			mower.Y--
		}
	}
	// Check if another mower is in this position
	for i := 0; i < len(lawn.Mowers); i++ {
		if lawn.Mowers[i].X == mower.X && lawn.Mowers[i].Y == mower.Y {
			return nil
		}
	}
	// Move mower if the position is empty
	lawn.Mowers[index].X = mower.X
	lawn.Mowers[index].Y = mower.Y
	return nil
}

func (lawn *Lawn) moveMower(index int) error {
	if index < 0 || index > len(lawn.Mowers)-1 {
		return errors.New("This mower doesn't exist in the lawn")
	}
	for _, instruction := range lawn.Mowers[index].Instructions {
		switch instruction {
		case "L":
			lawn.Mowers[index].turnLeft()
		case "R":
			lawn.Mowers[index].turnRight()
		case "F":
			err := lawn.moveMowerForward(index)
			if err != nil {
				return err
			}
		}
		lawn.Mowers[index].Instructions = lawn.Mowers[index].Instructions[1:]
	}
	return nil
}
