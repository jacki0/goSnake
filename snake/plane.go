package snake

type plane struct {
    height int
    width  int
    snake *snake
}

func newPlane(h, w int) *plane {
    pln := &plane{
        height: h,
        width:  w,
    }
    return pln
}

func (p *plane) snakeLeftPlane() bool {
    h := p.snake.head()
    return h.x > p.width || h.y > p.height || h.x < 0 || h.y < 0
}

func (p *plane) moveSnake() error {
	if err := p.snake.move(); err != nil {
	    return err
    }
    if p.snakeLeftPlane() {
        return p.snake.die()
    }
    return nil
}