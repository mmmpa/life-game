package life_game

type Field struct {
	w     int
	h     int
	cells []int
}

func createFiled(w, h int) Field {
	f := Field{}
	f.w = w
	f.h = h
	f.cells = make([]int, w * h)

	return f;
}

func (f Field) cell(x, y int) int {
	if (x < 0 || y < 0 || x > f.w - 1 || y > f.h - 1) {
		return 0
	}

	return f.cells[y * f.w + x]
}

func (f Field) addLife(x, y int) {
	f.cells[y * f.w + x] = 1
}

func (f *Field) takeLife(cells []Position) int {
	for _, p := range cells {
		f.addLife(p.X, p.Y)
	}

	return 0;
}

func (f Field) game() Field {
	field := createFiled(f.w, f.h)

	for y := 0; y < f.h; y++ {
		for x := 0; x < f.w; x++ {
			if f.isAlive(x, y) {
				field.addLife(x, y)
			}
		}
	}

	return field
}

func (f Field) isAlive(x, y int) bool {
	// 周りの生命の数を数える
	lives := 0;

	lives += f.cell(x - 1, y - 1)
	lives += f.cell(x + 0, y - 1)
	lives += f.cell(x + 1, y - 1)
	lives += f.cell(x - 1, y)
	lives += f.cell(x + 1, y)
	lives += f.cell(x - 1, y + 1)
	lives += f.cell(x + 0, y + 1)
	lives += f.cell(x + 1, y + 1)

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

	for y := 0; y < f.h; y++ {
		for x := 0; x < f.w; x++ {
			if f.cell(x, y) == 1 {
				str += "■ "
			} else {
				str += "□ "
			}

			if x == f.w - 1 {
				str += "\n"
			}
		}
	}

	return str
}
