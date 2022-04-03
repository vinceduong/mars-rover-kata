package application

import (
	"kata-mars-rover/domain"
	"kata-mars-rover/persistence"
)

type CreateTerrainDTO struct {
	height int
	width  int
}

func CreateTerrain(dto CreateTerrainDTO) error {
	var err error
	persistence.Terrain, err = domain.NewTerrain(dto.height, dto.width)

	if err != nil {
		return err
	}

	return nil
}
