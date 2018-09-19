// package executor

// import (
// 	"github.com/qezz/go-task-runner/pkg/task"
// )

// // Executor ...
// type Executor struct {
// 	incomingTasks chan *task.Task
// 	pool          []*task.Task
// }

// // NewExecutor create a executor
// func NewExecutor() *Executor {
// 	return &Executor{
// 		incomingTasks: make(chan *task.Task, 10), // TODO: Fix magic number
// 		pool:          make([]*task.Task, 0),
// 	}
// }

// func (ex *Executor) WaitForExit() {
// 	for _, t := range ex.pool {
// 		// TODO: make async
// 		(*t).WaitStop()
// 	}
// }

// func (ex *Executor) Add(task *task.Task) {
// 	ex.incomingTasks <- task
// }
