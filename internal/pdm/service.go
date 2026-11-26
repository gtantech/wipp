package pdm

import (
	"uuid"

	"github.com/gtantech/pdm"
)

type Service struct {
	projects map[uuid.UUID]pdm.PDM[task]
}
