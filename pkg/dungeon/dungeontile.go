package dungeon

//Tile represents a dungeon tile
type Tile int

const (
	dungeonTileEmpty     = 0
	dungeonTileDirtWall  = 1
	dungeonTileDirtFloor = 2
	dungeonTileStoneWall = 3
	dungeonTileCorridor  = 4
	dungeonTileDoor      = 5
	dungeonTileMud       = 9
	dungeonTileForrest   = 10
)
