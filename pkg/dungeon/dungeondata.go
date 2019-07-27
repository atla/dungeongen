package dungeon

//Dimension - holds width and height of a dungeon
type Dimension struct {
	Width  int
	Height int
}

//Area -  Returns area of the Dimension struct
func (d *Dimension) Area() int {
	return d.Width * d.Height
}

//Data - contains all dungeon related data
type Data struct {
	Size Dimension
	Data []Tile
}

//NewData creates a new dungeon data object
func NewData(width, height int) *Data {
	dat := &Data{
		Size: Dimension{
			Width:  width,
			Height: height,
		},
		Data: make([]Tile, width*height),
	}
	return dat
}

func (d *Data) empty() {
	for i := 0; i < len(d.Data); i++ {
		d.Data[i] = dungeonTileEmpty
	}
}

//GetTile returns Tile information for x, y coords
func (d *Data) GetTile(x int, y int) Tile {
	return d.Data[y*d.Size.Width+x]
}

//GetTileByIndex returns Tile information for index
func (d *Data) GetTileByIndex(index int) Tile {
	return d.Data[index]
}

//Copy creates a copy of the data
func (d *Data) Copy() *Data {
	dat := NewData(d.Size.Width, d.Size.Height)
	dat.Data = append(d.Data[:0:0], d.Data...)
	return dat
}
