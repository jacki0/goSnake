package snake

import "github.com/nsf/termbox-go"

var keyboardEventChan = make(chan keyboardEvent)

type Game struct {
	plane *plane
}

func initialPlane() *plane {
	return newPlane(20, 30)
}

func NewGame() *Game {
	return &Game{plane: initialPlane()}
}
func (game *Game) Start() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()
	go listenToKeyboard(keyboardEventChan)
}