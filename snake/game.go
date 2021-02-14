package snake

import (
	"github.com/nsf/termbox-go"
	"time"
)

var keyboardEventChan = make(chan keyboardEvent)

type (
	coord struct {
		x, y int
	}
	Game struct {
		plane *plane
		score int
		isOver bool
	}
)

func initialPlane() *plane {
	return newPlane(20, 30)
}

func NewGame() *Game {
	return &Game{plane: initialPlane()}
}

func (game *Game) moveInterval() time.Duration {
	ms := 100 - (game.score / 10)
	return time.Duration(ms) * time.Millisecond
}

func (game *Game) retry() {
	game.plane = initialPlane()
	game.isOver = false
}

func (game *Game) Start() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()
	go listenToKeyboard(keyboardEventChan)
	if err := game.render(); err != nil {
		panic(err)
	}
	mainloop:
		for {
			select {
			case ev := <- keyboardEventChan:
				switch ev.eventType {
				case MOVE:
					d := keyToDirection(ev.key)
					game.plane.snake.changeDirection(d)
				case RETRY:
					game.retry()
				case END:
					break mainloop
				}
			default:
				if !game.isOver {
					if err := game.render(); err != nil {
						panic(err)
					}
				}
				time.Sleep(game.moveInterval())
				}
		}
}