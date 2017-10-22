package lifegame

import "runtime"

type Task struct {
	Start int
	End   int
	Field Field
}

type Cells struct {
	Task  Task
	Cells []Cell
}

type ResultOut chan Field
type WorkerIn chan Field
type WorkerOut chan Cells

// boss を作成する
// worker を作成する
func parallel(field Field) chan Field {
	ch := make(chan Field)

	go func() {
		// 分割を決定
		cpus := runtime.NumCPU()

		tasks := split(field, cpus)
		bossOut, workerIn, workerOut := boss(field, cpus)

		worker(workerIn, workerOut, tasks)

		for result := range bossOut {
			ch <- result
		}
	}()

	return ch
}

func split(field Field, workersNum int) []Task {
	l := len(field.Cells)
	base := l / workersNum
	rest := l % workersNum

	tasks := make([]Task, workersNum)
	tail := 0
	for i := 0; i < workersNum; i++ {
		end := tail + base

		if rest > 0 {
			rest -= 1
			end += 1
		}

		tasks[i] = Task{
			Start: tail,
			End:   end,
			Field: field,
		}
	}

	return tasks
}

// 並列化: worker に field を供給する
// 同期: worker から演算済みの Cells を得て新しい field に反映する
func boss(field Field, workersNum int) (ResultOut, WorkerIn, WorkerOut) {
	resultOut := make(ResultOut)
	workerIn := make(WorkerIn)
	workerOut := make(WorkerOut)

	go func() {
		base := field

		for {
			for i := 0; i < workersNum; i++ {
				workerIn <- base
			}

			nextField := field.blank()

			// write 作業なので直列で行う
			for i := 0; i < workersNum; i++ {
				cells := <-workerOut
				// field への適用
				for p, cell := range cells.Cells {
					nextField.Cells[cells.Task.Start+p] = cell
				}
			}

			base = nextField
			resultOut <- nextField
		}
	}()

	return resultOut, workerIn, workerOut
}

// 開始: boss から field を受け取る
// 終了: 自分の担当範囲の演算が終わる
func worker(in WorkerIn, out WorkerOut, tasks []Task) {
	for _, task := range tasks {
		go func(task Task) {
			for {
				field := <-in
				cells := make([]Cell, task.End-task.Start)
				for i := task.Start; i < task.End; i++ {
					y := i / task.Field.W
					x := 0
					if y != 0 {
						x = i % y
					}

					if field.isAlive(x, y) {
						cells[i] = Cell(1)
					}
				}
				out <- Cells{
					Task:  task,
					Cells: cells,
				}
			}
		}(task)
	}
}
