package dungeon


import (
	"github.com/atla/dungeongen/pkg/data"
)

//Generator ...
type Generator struct {
	dimension        data.Dimension
	creationStrategy (func(d *data.Data))
}

//NewGenerator ...
func NewGenerator() *Generator {
	return &Generator{}

}

//Dimension adds Dimension to generators
func (g *Generator) Dimension(width, height int) *Generator {
	g.dimension = Dimension{
		Width:  width,
		Height: height,
	}
	return g

}

func (g* Generator) build () Dungeon* {

	
}  