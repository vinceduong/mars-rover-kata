package application

import (
	"kata-mars-rover/domain"
	"kata-mars-rover/persistence"
)

type CommandRoverDTO struct {
	id      int
	command domain.Command
}

func CommandRover(dto CommandRoverDTO) error {
	return persistence.Terrain.MoveRover(dto.id, dto.command)
}
