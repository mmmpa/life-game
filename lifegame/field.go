package lifegame

import "log"

type Cell int

type Field struct {
	W     int
	H     int
	Cells []Cell
	L     int
}

type FieldPosition struct {
	X int
	Y int
}

func createFiled(w, h int, lifePositions []FieldPosition) Field {
	log.Printf("create\n")

	f := Field{}
	f.W = w
	f.H = h
	f.L = w * h
	f.Cells = make([]Cell, f.L)

	for _, p := range lifePositions {
		f.addLife(p.X, p.Y)
	}

	return f
}

func (f Field) blank() Field {
	newField := Field{}
	newField.W = f.W
	newField.H = f.H
	newField.L = f.W * f.H
	newField.Cells = make([]Cell, newField.L)

	return newField
}

func (f Field) addLife(x, y int) {
	f.Cells[y*f.W+x] = Cell(1)
}

func (f Field) isAlive(i int) bool {
	y := i / f.W
	x := i % f.W

	// log.Printf("x: %v, y: %v\n", x, y)
	// 周りの生命の数を数える
	lives := 0

	hasLeft := x > 0
	hasRight := x < f.W-1
	hasUp := y > 0
	hasDown := y < f.H-1

	if hasUp {
		if hasLeft {
			lives += f.life(i - (f.W + 1))
		}
		lives += f.life(i - f.W)
		if hasRight {
			lives += f.life(i - (f.W - 1))
		}
	}

	if hasLeft {
		lives += f.life(i - 1)
	}
	if hasRight {
		lives += f.life(i + 1)
	}

	if hasDown {
		if hasLeft {
			lives += f.life(i + (f.W - 1))
		}
		lives += f.life(i + f.W)
		if hasRight {
			lives += f.life(i + (f.W + 1))
		}
	}

	// log.Printf("position: %+v, my_life: %v, life: %v\n", i, f.life(i), lives)

	// 生死を判断する
	switch {
	case lives == 1:
		return false
	case lives == 2 && f.life(i) == 1:
		return true
	case lives == 3:
		return true
	default:
		return false
	}
}

func (f Field) life(i int) int {
	return int(f.Cells[i])
}

func (f Field) generateFieldString() string {
	str := ""

	for y := 0; y < f.H; y++ {
		for x := 0; x < f.W; x++ {
			if f.life(y*f.W+x) == 1 {
				str += "■ "
			} else {
				str += "□ "
			}

			if x == f.W-1 {
				str += "\n"
			}
		}
	}

	return str
}
