package dungeon


type creationStrategy interface {
	create (data *data.Data)
}