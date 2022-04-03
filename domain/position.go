package domain

type Direction string

const (
	North Direction = "north"
	South Direction = "south"
	West  Direction = "west"
	East  Direction = "east"
)

type Position struct {
	x int
	y int
	d Direction
}

type Delta struct {
	x int
	y int
}

var directionsDelta = map[Direction]Delta{
	North: {0, 1},
	South: {0, -1},
	East:  {1, 0},
	West:  {-1, 0},
}

func (d Direction) TurnTo(r Rotation) Direction {
	switch d {
	case North:
		if r == TurnLeft {
			return West
		}
		return East
	case East:
		if r == TurnLeft {
			return North
		}
		return South
	case South:
		if r == TurnLeft {
			return East
		}
		return West
	case West:
		if r == TurnLeft {
			return South
		}
		return North
	default:
		return ""
	}
}

func (p Position) ApplyDelta(d Delta) Position {
	return Position{x: p.x + d.x, y: p.y + d.y, d: p.d}
}

func (p Position) MoveForward() Position {
	delta := directionsDelta[p.d]

	return p.ApplyDelta(delta)
}

func (p Position) ApplyCommand(c Command) Position {
	if c.mov {
		return p.MoveForward()
	}

	direction := p.d.TurnTo(c.rot)
	return Position{p.x, p.y, direction}
}
