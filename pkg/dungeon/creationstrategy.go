package dungeon

import (
	"github.com/atla/dungeongen/pkg/data"
)

type creationStrategy interface {
	create (data *data.Data)
}