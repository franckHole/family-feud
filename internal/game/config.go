package game

import (
	"github.com/franciscolkdo/family-feud/internal/family"
	"github.com/franciscolkdo/family-feud/internal/table"
)

type Config struct {
	Table      table.Config  `json:"table"`
	BlueFamily family.Config `json:"blueFamily"`
	RedFamily  family.Config `json:"redFamily"`
}
