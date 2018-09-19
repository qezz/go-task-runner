package heart

import (
	"context"
	"fmt"
	"time"

	"github.com/qezz/go-task-runner/pkg/task"
)

type ExecutorState int

const (
	NOTYETSTARTED ExecutorState = iota
	RUNNING       ExecutorState = iota
	STOPPING      ExecutorState = iota
	STOPPED       ExecutorState = iota
)

type Executor struct {
	// incomingTasks chan *task.Task
	// pool          []*task.Task

	state ExecutorState
}

func NewExecutor() *Executor {
	return &Executor{
		// incomingTasks: make(chan *task.Task, 10), // TODO: Fix magic number
		// pool:          make([]*task.Task, 0),

		state: NOTYETSTARTED,
	}
}

// func (ex *Executor) Add(task *task.Task) error {
// 	// if (ex.state != RUNNING) || (ex.state != NOTYETSTARTED) {
// 	// 	return fmt.Errorf("Executor is stopping, unable to add a task")
// 	// }
// 	// ex.incomingTasks <- task

// 	// panic("Not implemented, yet")

// 	return nil
// }

func (ex *Executor) Execute(ctx *context.Context, t *task.Task, now bool) error {
	// ctx, cancel := context.WithCancel(context.Background())

	interval := time.NewTicker(time.Duration((*t).GetInterval()))

	for {
		// create new deadline after successful evaluation of the task
		deadline := time.NewTicker(time.Duration((*t).GetDeadline()))

		select {
		case <-interval.C:
			// valueCh, err := task.Poll
			select {
			case val := <-(*t).Eval():
				fmt.Println("value:", val)
			case <-deadline.C: // TODO: Add custom event on deadline?
				// cancel the current goroutined task
				// cancel()
				fmt.Println("deadline")
				continue
			}
		case <-(*ctx).Done():
			return fmt.Errorf("task %v finished early")
		}
	}
}
