package main

import (
	"time"

	"github.com/qezz/go-task-runner/pkg/heart"
)

type Fib struct {
	f1 int
	f2 int
}

func NewFib() *Fib {
	return &Fib{
		f1: 0,
		f2: 1,
	}
}

func (fib *Fib) Poll() (interface{}, error) {
	res := fib.f1 + fib.f2
	fib.f1 = fib.f2
	fib.f2 = res

	return res, nil
}

func main() {
	h := heart.NewHeart()

	fib := NewFib()
	t := task.FromFuture(fib).
		WithName("fib").
		WithDesc("Fibonacci delayed generator").
		Every(2 * time.Second).
		WithDeadline(3 * time.Second)

	h.AddTask(t)

	h.Run()

	// ---

	time.Sleep(5 * time.Second)

	// ---

	fmt.Println("state of task:", h.StateOf("fib"))
	h.Stop("fib")
	fmt.Println("state of task:", h.StateOf("fib"))

	time.Sleep(2 * time.Second)

	fmt.Println("state of task:", h.StateOf("fib"))
}
