package domain

import "errors"

type Plateau struct {
	height   int
	width    int
	roverPos map[int]Position
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

func (plat *Plateau) ValidatePos(pos Position) {
	if err := pos.d.Validate(); err != nil {
		panic("Rover direction needs to be valid")
	}

	if pos.x < 0 || pos.x >= plat.width {
		panic("Rover x position cannot be off-plateau")
	}

	if pos.y < 0 || pos.y >= plat.height {
		panic("Rover y position cannot be off-plateau")
	}
}

func (plat *Plateau) SpawnRover(id int, pos Position) {
	plat.ValidatePos(pos)

	plat.roverPos[id] = pos
}

func (p *Plateau) MoveRover(id int, c Command) error {
	pos, ok := p.roverPos[id]

	if !ok {
		return errors.New("rover does not exist")
	}

	pos = pos.ApplyCommand(c)
	p.ValidatePos(pos)

	p.roverPos[id] = pos

	return nil
}
