package snake

import "github.com/nsf/termbox-go"

type keyboardEventType int

type keyboardEvent struct {
	eventType keyboardEventType
	key 	  termbox.Key
}

func listenToKeyboard(evChan chan keyboardEvent) {
	termbox.SetInputMode(termbox.InputEsc)
}