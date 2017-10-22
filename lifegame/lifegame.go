package lifegame

import (
	"fmt"
	"time"
)


func Run(wait, w, h int, cells []FieldPosition) {
	ch := start(w, h, cells)

	for field := range ch {
		time.Sleep(time.Duration(wait) * time.Millisecond)
		fmt.Printf("\x1b[%dA", h+1)
		fmt.Printf(field.generateFieldString())
	}
}

func start(w, h int, lifePositions []FieldPosition) chan Field {
	field := createFiled(w, h, lifePositions)

	return parallel(field)
}
