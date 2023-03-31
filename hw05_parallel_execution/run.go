package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if n < 1 || m <= 0 {
		return ErrErrorsLimitExceeded
	}
	mA := int64(m)
	wg := &sync.WaitGroup{}
	nk := 0

	for _, task := range tasks {
		wg.Add(1)
		nk++
		go func(task Task) {
			defer wg.Done()
			err := task()
			if err != nil {
				atomic.AddInt64(&mA, -1)
			}
		}(task)

		if nk%n == 0 && nk > 0 {
			wg.Wait()
			if mA <= 0 {
				return ErrErrorsLimitExceeded
			}
			nk = 0
		}
	}

	if nk > 0 {
		wg.Wait()
		if mA <= 0 {
			return ErrErrorsLimitExceeded
		}
	}

	return nil
}
