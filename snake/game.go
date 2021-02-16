package snake

import (
	"github.com/nsf/termbox-go"
	"time"
)

var (
	keyboardEventChan = make(chan keyboardEvent)
	pointsChan = make(chan int)
)

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

func initialSnake() *snake {
	return newSnake(RIGHT, []coord{
		{x: 1, y: 1},
		{x: 1, y: 2},
		{x: 1, y: 3},
		{x: 1, y: 4},
	})
}
func initialPlane() *plane {
	return newPlane(initialSnake(), pointsChan,20, 30)
}

func NewGame() *Game {
	return &Game{plane: initialPlane()}
}

func (game *Game) addPoints(p int) {
	game.score += p
}

func (game *Game) end() {
	game.isOver = true
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
		case point := <- pointsChan:
			game.addPoints(point)
		default:
			if !game.isOver {
				if err := game.plane.moveSnake(); err != nil {
					game.end()
				}
			}
			if err := game.render(); err != nil {
				panic(err)
				}
			time.Sleep(game.moveInterval())
		}
	}
}