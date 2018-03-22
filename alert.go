package alert

//Alert represents an alert that is raised
type Alert interface {
	Stop()
	Channel() chan<- interface{}
}

type baseAlert struct {
	ch     chan interface{}
	done   chan int
	action func(interface{})
}

func (a *baseAlert) Channel() chan<- interface{} {
	return a.ch
}

func (a *baseAlert) Stop() {
	a.done <- 1
}

type raiseAlert struct {
	baseAlert
	threshold          float64
	minimumFluctuation float64
	previousValue      float64
	alertOn            bool
}

//NewRaiseAlert returns a alert that trigger when a threshold is reached by raising.
func NewRaiseAlert(threshold, fluctuation float64, action func(interface{})) Alert {
	a := &raiseAlert{
		baseAlert: baseAlert{
			ch:     make(chan interface{}),
			done:   make(chan int),
			action: action,
		},
		threshold:          threshold,
		minimumFluctuation: fluctuation,
	}
	a.start()
	return a

}

func (a *raiseAlert) check(v float64) bool {
	fluctuation := v - a.previousValue
	a.previousValue = v
	if v < a.threshold {
		if !a.alertOn {
			return false
		}
		if a.threshold-a.minimumFluctuation > v {
			a.alertOn = false
			return false
		}

		return false
	}
	if a.alertOn {
		return false
	}
	if fluctuation <= 0 {
		return false
	}
	a.alertOn = true
	return true
}

func (a *raiseAlert) start() {
	go func() {
		var exit bool
		for !exit {
			select {
			case val := <-a.ch:
				v, ok := val.(Valuer)

				if !ok {
					continue
				}

				if !a.check(v.Value()) {
					continue
				}
				a.alertOn = true

				if a.action != nil {
					a.action(val)
				}

			case <-a.done:
				exit = true
			}
		}
	}()
}
