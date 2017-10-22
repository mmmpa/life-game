package lifegame

import (
	"fmt"
	"time"
)


func Run(wait, w, h int, cells []FieldPosition) {
	ch := start(w, h, cells)

	for field := range ch {
		fmt.Printf(field.generateFieldString())
		fmt.Printf("\x1b[%dA", h)
		time.Sleep(time.Duration(wait) * time.Millisecond)
	}
}

func start(w, h int, lifePositions []FieldPosition) chan Field {
	field := createFiled(w, h, lifePositions)

	return parallel(field)
}
