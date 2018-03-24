package alert

import "sync"

type Valuer interface {
	Value() float64
}

type Monitor struct {
	MonitorChan chan interface{}
	done        chan int
	Wg          *sync.WaitGroup
	Dispatcher
}

func NewMonitor(wg *sync.WaitGroup) *Monitor {
	m := &Monitor{
		MonitorChan: make(chan interface{}),
		done:        make(chan int),
		Dispatcher:  newDispatcher(),
		Wg:          wg,
	}
	m.start()
	return m
}

func (m *Monitor) start() {
	go func() {
		var stop = false
		for !stop {
			select {
			case val := <-m.MonitorChan:
				m.Emit(val)
				m.Wg.Done()
			case <-m.done:
				stop = true
			}
		}
	}()
}

func (m *Monitor) Stop() {
	m.done <- 1
	m.Dispatcher.Stop()

}
