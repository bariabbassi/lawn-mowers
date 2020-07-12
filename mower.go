package main

// Mower position and instructions
type Mower struct {
	X            int
	Y            int
	Orientation  string
	Instructions []string
}

func (mower *Mower) turnLeft() {
	switch mower.Orientation {
	case "N":
		mower.Orientation = "W"
	case "E":
		mower.Orientation = "N"
	case "W":
		mower.Orientation = "S"
	case "S":
		mower.Orientation = "E"
	}
}

func (mower *Mower) turnRight() {
	switch mower.Orientation {
	case "N":
		mower.Orientation = "E"
	case "E":
		mower.Orientation = "S"
	case "W":
		mower.Orientation = "N"
	case "S":
		mower.Orientation = "W"
	}
}
