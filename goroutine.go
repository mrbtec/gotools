package gotools

import "sync"

type Goroutine struct {
	wg sync.WaitGroup
}

func New() *Goroutine {
	return &Goroutine{}
}

func (g *Goroutine) Run(f func()) {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		f()
	}()
}

func (g *Goroutine) Wait() {
	g.wg.Wait()
}
