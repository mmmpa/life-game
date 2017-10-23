package lifegame

import (
	"runtime"
)

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
		resultOut, workerIn, workerOut := boss(field, cpus)

		worker(workerIn, workerOut, tasks)

		for result := range resultOut {
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
	tail := 1
	for i := 0; i < workersNum; i++ {
		end := tail + (base - 1)

		if rest > 0 {
			rest -= 1
			end += 1
		}

		tasks[i] = Task{
			Start: tail - 1,
			End:   end - 1,
			Field: field,
		}

		tail = end + 1
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

		for i := 0; true; i++ {
			resultOut <- base

			for i := 0; i < workersNum; i++ {
				workerIn <- base
			}

			nextField := field.blank()

			// write 作業なので直列で行う
			for i := 0; i < workersNum; i++ {
				cells := <-workerOut
				for p, cell := range cells.Cells {
					nextField.Cells[cells.Task.Start+p] = cell
				}
			}

			base = nextField
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
				l := (task.End - task.Start) + 1
				cells := make([]Cell, l)
				for i := 0; i < l; i++ {
					if field.isAlive(task.Start + i) {
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
