package main

type Plateau struct {
	height int
	width  int
}

func NewPlateau(height, width int) *Plateau {
	if height <= 0 {
		panic("Plateau height cannot be negative")
	}
	if width <= 0 {
		panic("Plateau width cannot be negative")
	}

	return &Plateau{height: height, width: width}
}
