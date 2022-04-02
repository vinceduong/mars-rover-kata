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
