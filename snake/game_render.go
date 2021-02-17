package snake

import (
	"fmt"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

const (
	defaultColor = termbox.ColorDefault
	bgColor = termbox.ColorDefault
	snakeColor = termbox.ColorGreen
)

func (game *Game) render() error {
	err := termbox.Clear(defaultColor, defaultColor)
	if err != nil {
		panic(err)
	}
	var (
		w, h = termbox.Size()
		midY = h / 2
		left = (w - game.plane.width) / 2
		right = (w + game.plane.width) / 2
		top = midY - (game.plane.height / 2)
		bottom = midY + (game.plane.height / 2) + 1
	)
	renderTitle(left, top)
	renderPlane(game.plane, top, bottom, left)
	renderSnake(left, bottom, game.plane.snake)
	renderFood(left, bottom, game.plane.food)
	renderScore(left, bottom, game.score)
	renderQuitMsg(right, bottom)
	return termbox.Flush()
}

func renderSnake(left, bottom int, s *snake)  {
	for _, b := range s.body {
		termbox.SetCell(left + b.x, bottom - b.y, ' ', snakeColor, snakeColor)
	}
}

func renderFood(left, bottom int, f *food) {
	termbox.SetCell(left + f.x, bottom - f.y, f.emoji, defaultColor, bgColor)
}

func renderScore(left, bottom, s int) {
	score := fmt.Sprintf("Score %v", s)
	termboxPrint(left, bottom + 1, score)
}

func renderQuitMsg(right, bottom int) {
	msg := "Press ESC to quit"
	termboxPrint(right - len(msg), bottom + 1, msg)
}

func fill (x, y, w, h int, cell termbox.Cell){
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < w; lx++ {
			termbox.SetCell(x + lx, y + ly, cell.Ch, cell.Fg, cell.Bg)
		}
	}
}


func renderPlane(p *plane, top, bottom, left int) {
	for i := top; i < bottom; i++ {
		termbox.SetCell(left - 1, i, '|', defaultColor, bgColor)
		termbox.SetCell(left + p.width, i, '|', defaultColor, bgColor)
	}
	termbox.SetCell(left - 1, top, '\u250C', defaultColor, bgColor)
	termbox.SetCell(left - 1, bottom, '\u2514', defaultColor, bgColor)
	termbox.SetCell(left + p.width, top, '\u2510', defaultColor, bgColor)
	termbox.SetCell(left + p.width, bottom, '\u2518', defaultColor, bgColor)
	fill(left, top, p.width, 1, termbox.Cell{Ch: '─'})
	fill(left, bottom, p.width, 1, termbox.Cell{Ch: '─'})
	/* need test:
	fill(left, top, p.width, 1, termbox.Cell{Ch: '\u2014'})
	fill(left, bottom, p.width, 1, termbox.Cell{Ch: '\u2014'})
	and:
	fill(left, top, p.width, 1, termbox.Cell{Ch: '\u2013'})
	fill(left, bottom, p.width, 1, termbox.Cell{Ch: '\u2013'})
	 */
}

func renderTitle(left, top int) {
	termboxPrint(left, top - 1, "Snake Game")
}

func termboxPrint(x, y int, msg string)  {
	for _, c := range msg {
		termbox.SetCell(x, y, c, defaultColor, defaultColor)
		x += runewidth.RuneWidth(c)
	}
}