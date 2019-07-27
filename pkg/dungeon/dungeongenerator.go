package dungeon

//Generator ...
type Generator struct {
	dimension        Dimension
	creationStrategy (func(d *Data))
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
