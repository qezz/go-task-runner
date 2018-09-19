package task

import (
	"context"
)

type TaskState int

const (
	ACTIVE   TaskState = iota
	INACTIVE TaskState = iota
	NOTFOUND TaskState = iota
)

type TaskInfo struct {
	state TaskState

	name   string
	ctx    *context.Context
	cancel context.CancelFunc
}

func NewTaskInfoWith(name string, ctx *context.Context, cancel context.CancelFunc) *TaskInfo {
	return &TaskInfo{
		state:  INACTIVE,
		name:   name,
		ctx:    ctx,
		cancel: cancel,
	}
}
