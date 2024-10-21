package game

import (
	"github.com/franciscolkdo/family-feud/internal/family"
	"github.com/franciscolkdo/family-feud/internal/table"
)

type Config struct {
	Table    table.Config    `json:"table"`
	Families []family.Config `json:"families"`
}
