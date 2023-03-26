package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}
	mMx := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	nk := 0

	for _, task := range tasks {
		wg.Add(1)
		nk++
		go func(task Task) {
			defer wg.Done()
			err := task()
			if err != nil {
				mMx.Lock()
				m--
				mMx.Unlock()
			}
		}(task)

		if nk%n == 0 && nk > 0 {
			wg.Wait()
			if m <= 0 {
				return ErrErrorsLimitExceeded
			}
			nk = 0
		}
	}

	if nk > 0 {
		wg.Wait()
		if m <= 0 {
			return ErrErrorsLimitExceeded
		}
	}

	return nil
}
