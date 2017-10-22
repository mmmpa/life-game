package lifegame

type Cell int

type Field struct {
	W     int
	H     int
	Cells []Cell
}

type FieldPosition struct {
	X int
	Y int
}

func createFiled(w, h int, lifePositions []FieldPosition) Field {
	f := Field{}
	f.W = w
	f.H = h
	f.Cells = make([]Cell, w*h)

	for _, p := range lifePositions {
		f.addLife(p.X, p.Y)
	}

	return f
}

func (f Field) blank() Field {
	newField := Field{}
	newField.W = f.W
	newField.H = f.H
	newField.Cells = make([]Cell, f.W*f.H)

	return newField
}

func (f Field) cell(x, y int) Cell {
	if (x < 0 || y < 0 || x > f.W-1 || y > f.H-1) {
		return 0
	}

	return f.Cells[y*f.W+x]
}

func (f Field) addLife(x, y int) {
	f.Cells[y*f.W+x] = 1
}

func (f Field) isAlive(x, y int) bool {
	// 周りの生命の数を数える
	lives := Cell(0)

	lives += f.cell(x-1, y-1)
	lives += f.cell(x+0, y-1)
	lives += f.cell(x+1, y-1)
	lives += f.cell(x-1, y)
	lives += f.cell(x+1, y)
	lives += f.cell(x-1, y+1)
	lives += f.cell(x+0, y+1)
	lives += f.cell(x+1, y+1)

	// 生死を判断する
	switch {
	case lives == 1:
		return false
	case lives == 2 && f.cell(x, y) == 1:
		return true
	case lives == 3:
		return true
	default:
		return false
	}
}

func (f Field) generateFieldString() string {
	str := ""

	for y := 0; y < f.H; y++ {
		for x := 0; x < f.W; x++ {
			if f.cell(x, y) == 1 {
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
