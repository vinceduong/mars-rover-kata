package domain

import "errors"

type Terrain struct {
	height   int
	width    int
	roverPos map[int]Position
}

func NewTerrain(height, width int) (*Terrain, error) {
	if height <= 0 {
		return nil, errors.New("teau height cannot be negative")
	}
	if width <= 0 {
		return nil, errors.New("teau width cannot be negative")
	}

	roverPos := make(map[int]Position)

	return &Terrain{height: height, width: width, roverPos: roverPos}, nil
}

func (t *Terrain) ValidatePos(pos Position) error {
	if err := pos.d.Validate(); err != nil {
		return err
	}

	if pos.x < 0 || pos.x >= t.width {
		return errors.New("err: rover x position cannot be off-teau")
	}

	if pos.y < 0 || pos.y >= t.height {
		return errors.New("err: rover y position cannot be off-teau")
	}

	return nil
}

func (t *Terrain) SpawnRover(id int, pos Position) error {
	if err := t.ValidatePos(pos); err != nil {
		return err
	}
	t.roverPos[id] = pos
	return nil
}

func (t *Terrain) MoveRover(id int, c Command) error {
	pos, ok := t.roverPos[id]

	if !ok {
		return errors.New("rover does not exist")
	}

	pos = pos.ApplyCommand(c)
	t.ValidatePos(pos)

	t.roverPos[id] = pos

	return nil
}
