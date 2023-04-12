package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {

	sChan := make([]chan interface{}, len(stages))

	runner := func(stage Stage, inner <-chan interface{}, outer chan<- interface{}) {
		defer close(outer)
		inValues := stage(inner)
		for {
			select {
			case <-done:
				return
			case inValue, open := <-inValues:
				if !open {
					return
				}
				outer <- inValue
			}
		}
	}

	for k, stage := range stages {
		sChan[k] = make(chan interface{})
		go func(k int, stage Stage) {
			if k == 0 {
				runner(stage, in, sChan[k])
			} else {
				runner(stage, sChan[k-1], sChan[k])
			}
		}(k, stage)
	}

	return sChan[len(stages)-1]
}
