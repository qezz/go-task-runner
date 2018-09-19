package heart

import (
	"context"
	"fmt"

	"github.com/qezz/go-task-runner/pkg/task"
)

type Heart struct {
	executor *Executor
	tasks    map[string]*task.TaskInfo
}

func NewHeart() *Heart {
	return &Heart{
		executor: NewExecutor(), // NewScheduler(),
		tasks:    make(map[string]*task.TaskInfo),
	}
}

func (self *Heart) AddTask(t *task.Task) error {
	taskName := (*t).GetName()

	_, ok := self.tasks[taskName]
	if ok {
		return fmt.Errorf("Name '%v' is already taken", taskName)
	}

	ctx, cancel := context.WithCancel(context.Background())
	self.tasks[taskName] = task.NewTaskInfoWith(taskName, &ctx, cancel)

	go self.executor.Execute(&ctx, t, true)

	return nil
}

func (h *Heart) Run() {
	sched.Poll()
}

func (h *Heart) StateOf(name string) TaskState {
	if t, ok := tasks[name]; ok {
		return t.State()
	}

	return task.NOTFOUND
}
