package domain

type Movement string
type Rotation string

const (
	TurnLeft  Rotation = "left"
	TurnRight Rotation = "right"
)

type Command struct {
	mov bool
	rot Rotation
}

type RoverCommand struct {
	id string
	c  Command
}
