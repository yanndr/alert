package alert

type Valuer interface {
	Value() float64
}

type Monitor struct {
	MonitorChan chan interface{}
	done        chan int
	eventDispatcher
}

func NewMonitor() *Monitor {
	m := &Monitor{
		MonitorChan: make(chan interface{}),
		done:        make(chan int),
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
				m.emit(val)
			case <-m.done:
				stop = true
			}

		}
	}()
}

func (m *Monitor) Stop() {
	m.done <- 1
	m.closeChannels()
}
