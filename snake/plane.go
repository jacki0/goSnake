package snake

import (
    "math/rand"
    "time"
)

type plane struct {
    food *food
    hasFood func(*plane, coord) bool
    height int
    width  int
    snake *snake
    pointsChan chan int
}

func newPlane(s *snake, p chan int, h, w int) *plane {
    rand.Seed(time.Now().UnixNano())
    pln := &plane{
        snake: s,
        height: h,
        width:  w,
        hasFood: hasFood,
        pointsChan: p,
    }
    pln.placeFood()
    return pln
}

func (p *plane) snakeLeftPlane() bool {
    h := p.snake.head()
    return h.x > p.width || h.y > p.height || h.x < 0 || h.y < 0
}

func (p *plane) addPoints(pp int) {
    p.pointsChan <- pp
}

func (p *plane) placeFood() {
    var x, y int
    for {
        x = rand.Intn(p.width)
        y = rand.Intn(p.height)
        if !p.isOccupied(coord{x: x, y: y}) {
            break
        }
    }
    p.food = newFood(x, y)
}

func (p *plane) moveSnake() error {
	if err := p.snake.move(); err != nil {
	    return err
    }
    if p.snakeLeftPlane() {
        return p.snake.die()
    }
    if p.hasFood(p, p.snake.head()) {
        go p.addPoints(p.food.points)
        p.snake.length++
        p.placeFood()
    }
    return nil
}

func hasFood(p *plane, c coord) bool {
    return c.x == p.food.x && c.y == p.food.y
}

func (p *plane) isOccupied(c coord) bool {
    return p.snake.isOnPosition(c)
}