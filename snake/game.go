package snake

type Game struct {
	plane *plane
}

func initialPlane() *plane {
	return newPlane(20, 30)
}

func NewGame() *Game {
	return &Game{plane: initialPlane()}
}