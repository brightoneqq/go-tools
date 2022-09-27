package exec

import "sync"

type Executor struct {
	c  chan struct{}
	wg *sync.WaitGroup
}

func NewExecutor(threshold int) *Executor {
	return &Executor{
		c:  make(chan struct{}, threshold),
		wg: new(sync.WaitGroup),
	}
}

func (s *Executor) Exec(action func()) {
	s.wg.Add(1)
	go func() {
		defer s.done()
		s.c <- struct{}{}
		action()
	}()
}

func (s *Executor) Await() {
	s.wg.Wait()
}

func (s *Executor) done() {
	<-s.c
	s.wg.Done()
}
