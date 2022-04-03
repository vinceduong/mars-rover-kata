package domain

import "errors"

type Plateau struct {
	height   int
	width    int
	roverPos map[int]Position
}

func (d Direction) Validate() error {
	switch d {
	case North, South, West, East:
		{
			return nil
		}
	}
	return errors.New("invalid direction")
}

func NewPlateau(height, width int) *Plateau {
	if height <= 0 {
		panic("Plateau height cannot be negative")
	}
	if width <= 0 {
		panic("Plateau width cannot be negative")
	}

	roverPos := make(map[int]Position)

	return &Plateau{height: height, width: width, roverPos: roverPos}
}

func (plat *Plateau) SpawnRover(id int, pos Position) {
	if pos.x < 0 || pos.x >= plat.width {
		panic("Rover x position cannot be off-plateau")
	}

	if pos.y < 0 || pos.y >= plat.height {
		panic("Rover y position cannot be off-plateau")
	}

	if err := pos.d.Validate(); err != nil {
		panic("Rover direction needs to be valid")
	}

	plat.roverPos[id] = pos
}

// func (p *Plateau) MoveRover(c RoverCommand) error {
// 	pos, ok := p.roverPos[c.id]

// 	if !ok {
// 		return errors.New("Rover does not exist")
// 	}

// }
