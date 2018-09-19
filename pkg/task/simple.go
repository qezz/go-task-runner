package task

import (
	"context"
	"time"

	"github.com/qezz/go-task-runner/pkg/future"
)

type SimpleTask struct {
	future future.Future

	name     string
	interval time.Duration
	deadline time.Duration
}

func (self *SimpleTask) FromFuture(fut future.Future) *SimpleTask {
	self.future = fut
	return self
}

func (self *SimpleTask) WithName(name string) *SimpleTask {
	self.name = name
	return self
}

func (self *SimpleTask) Every(timespan time.Duration) *SimpleTask {
	self.interval = timespan
	return self
}

// func (self *SimpleTask) WithTimeout(timeout time.Duration) *SimpleTask {
// 	self.timeout = timeout
// 	return self
// }

func (self *SimpleTask) WithDeadline(deadline time.Duration) *SimpleTask {
	self.deadline = deadline
	return self
}

// impl Task

func (self *SimpleTask) Eval(ctx *context.Context) (interface{}, error) {
	return self.future.Poll(ctx)
}

func (self *SimpleTask) getName() string {
	return self.name
}
