package snake

import (
	"errors"
	"os"
	"strings"
)

type food struct {
	emoji rune
	points, x, y int
}

func newFood(x, y int) *food {
	e, p := getFoodEmoji()
	if p == 0 {
		panic(errors.New("error when getting emoji"))
	}
	return &food{
		points: p,
		emoji: e,
		x: x,
		y: y,
	}
}

func getFoodEmoji() (rune, int) {
	if hasUnicodeSupport() {
		return randomFoodEmoji()
	}
	return '@', 1
}

func randomFoodEmoji() (v1 rune, v2 int) {
	food := map[rune]int{
		'🍒': 1,
		'🍍': 1,
		'🍑': 1,
		'🍇': 1,
		'🍏': 1,
		'🍌': 1,
		'🍫': 1,
		'🍭': 1,
		'🍕': 1,
		'🍩': 1,
		'🍗': 1,
		'🍖': 1,
		'🍬': 1,
		'🍤': 1,
		'🍪': 10,
	}
	for {
		for v1, v2 = range food {
			return
		}
		return
	}
}

func hasUnicodeSupport() bool {
	return strings.Contains(os.Getenv("LANG"), "UTF-8")
}