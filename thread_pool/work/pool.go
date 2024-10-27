package work

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

// Executor is an interface that defines
// the methods that must be implemented by
// a type that will be executed by the pool.
type Executor interface {
	Execute() error
	OnError(error)
}

// Pool is a struct that represents a pool
// of workers that will execute tasks.
// numWorkers is the number of workers that
// tasks will be distributed to.
// start is a sync.Once that will ensure that
// stop is a sync.Once that will ensure that
// quit is a channel that will be closed when
// the pool is stopped.
type Pool struct {
	numWorkers    int
	tasks         chan Executor
	start         sync.Once
	stop          sync.Once
	taskCompleted chan bool
	quit          chan struct{}
}

func NewPool(numWorkers int, taskChannelSize int) (*Pool, error) {
	if numWorkers <= 0 {
		return nil, errors.New("number of workers must be greater than 0")
	}

	if taskChannelSize <= 0 {
		return nil, errors.New("task channel size must be greater than 0")
	}

	return &Pool{
		numWorkers:    numWorkers,
		tasks:         make(chan Executor, taskChannelSize),
		start:         sync.Once{},
		stop:          sync.Once{},
		taskCompleted: make(chan bool),
		quit:          make(chan struct{}),
	}, nil
}

func (p *Pool) Start(ctx context.Context) {
	p.start.Do(func() {
		p.startWorker(ctx)
	})
}

func (p *Pool) TaskCompleted() <-chan bool {
	return p.taskCompleted
}

func (p *Pool) IsPoolDone() bool {
	return len(p.tasks) == 0
}

func (p *Pool) Stop() {
	p.stop.Do(func() {
		close(p.quit)
	})
}

func (p *Pool) AddTask(task Executor) {
	select {
	case p.tasks <- task:
	case <-p.quit:
	}
}

func (p *Pool) AddTaskNonBlocking(task Executor) {
	go func() {
		p.tasks <- task
	}()
}

func (p *Pool) startWorker(ctx context.Context) {
	for i := 0; i < p.numWorkers; i++ {
		go func(workerNum int) {
			fmt.Printf("Worker %d started\n", workerNum)
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("Worker %d stopped\n", workerNum)
					return
				case <-p.quit:
					fmt.Printf("Worker %d stopped\n", workerNum)
					return
				case task, ok := <-p.tasks:
					if !ok {
						fmt.Printf("Worker %d stopped\n", workerNum)
						return
					}
					err := task.Execute()
					if err != nil {
						task.OnError(err)
					}
					p.taskCompleted <- true
					fmt.Printf("Worker %d completed task\n", workerNum)
				}
			}
		}(i)
	}
}
