package snake

const (
	RIGHT = 1 + iota
	LEFT
	UP
	DOWN
)

type (
	direction int
	snake struct {
		body []coord
		direction direction
		length int
	}
)

func (s *snake) changeDirection(d direction) {
	opposites := map[direction]direction{
		RIGHT: LEFT,
		LEFT:  RIGHT,
		UP:    DOWN,
		DOWN:  UP,
	}
	if o := opposites[d]; o != 0 && 0 != s.direction {
		s.direction = d
	}
}