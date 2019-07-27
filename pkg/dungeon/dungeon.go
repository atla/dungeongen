package dungeon

// Dungeon structure
type Dungeon struct {
	Data *Data
	
}

// New creates a new dungeon
func New() *Dungeon {
	return &Dungeon{}
}
