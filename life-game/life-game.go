package life_game

import (
	"fmt"
	"time"
)

type Position struct {
	X int
	Y int
}

func Run(wait, w, h int, cells []Position) {
	field := createFiled(w, h)
	field.takeLife(cells)

	fmt.Printf(field.generateFieldString())

	for {
		time.Sleep(time.Duration(wait) * time.Millisecond)

		field = field.game()
		fmt.Printf("\x1b[%dA", h)
		fmt.Printf(field.generateFieldString())
	}
}
