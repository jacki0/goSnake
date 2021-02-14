package snake

import "github.com/nsf/termbox-go"

const (
	defaultColor = termbox.ColorDefault
	bgColor = termbox.ColorDefault
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
		top = midY - (game.plane.height / 2)
		bottom = midY + (game.plane.height / 2) + 1
	)
	renderPlane(game.plane, top, bottom, left)
	return termbox.Flush()
}
func fill (x, y, w, h int, cell termbox.Cell){
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < w; lx++ {
			termbox.SetCell(x + lx, y + ly, cell.Ch, cell.Fg, cell.Bg)
		}
	}
}

/* need test:
func fill (x, y, w, h int, cell termbox.Cell){
	for ; h >= 0; h-- {
		for ; w >= 0; w-- {
			termbox.SetCell(x + w, y + h, cell.Ch, cell.Fg, cell.Bg)
		}
	}
}*/

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