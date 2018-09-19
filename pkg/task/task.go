package task

import (
	"time"
)

type Task interface {
	Eval() chan interface{}
	Poll() (chan interface{}, error)

	// WaitStop()
	// ReadInterruptChan() <-chan bool

	// FIXME: Remove interval, introduce timetable
	GetInterval() time.Duration
	GetDeadline() time.Duration

	GetName() string
	WaitStop()
}
