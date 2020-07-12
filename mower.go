package main

type Mower struct {
	X            int
	Y            int
	Orientation  string
	Instructions []string
}

func (m *Mower) move(lawn Lawn) {
	for _, instruction := range m.Instructions {
		switch instruction {
		case "L":
			m.turnLeft()
		case "R":
			m.turnRight()
		case "F":
			m.moveForward(lawn)
		}
		m.Instructions = m.Instructions[1:]
	}
}

func (m *Mower) turnLeft() {
	switch m.Orientation {
	case "N":
		m.Orientation = "W"
	case "E":
		m.Orientation = "N"
	case "W":
		m.Orientation = "S"
	case "S":
		m.Orientation = "E"
	}
}

func (m *Mower) turnRight() {
	switch m.Orientation {
	case "N":
		m.Orientation = "E"
	case "E":
		m.Orientation = "S"
	case "W":
		m.Orientation = "N"
	case "S":
		m.Orientation = "W"
	}
}

func (m *Mower) moveForward(lawn Lawn) {
	switch m.Orientation {
	case "N":
		if m.Y+1 < lawn.Length {
			m.Y++
		}
	case "E":
		if m.X+1 < lawn.Width {
			m.X++
		}
	case "W":
		if m.X-1 >= 0 {
			m.X--
		}
	case "S":
		if m.Y-1 >= 0 {
			m.Y--
		}
	}
}
