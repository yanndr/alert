package alert

import "sync"

type Subscriber interface {
	Subscribe(e string, ch chan<- interface{})
	unsubscribe(e string, ch chan<- interface{})
}

type Dispatcher interface {
	Subscriber
	Emit(interface{})
	Stop()
}

type eventDispatcher struct {
	mutex         sync.RWMutex
	eventChannels map[string]chan<- interface{}
}

func newDispatcher() Dispatcher {
	return &eventDispatcher{
		eventChannels: make(map[string]chan<- interface{}),
	}
}

func (d *eventDispatcher) Emit(data interface{}) {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	for _, ch := range d.eventChannels {
		ch <- data
	}
}

func (d *eventDispatcher) Stop() {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	for key, ch := range d.eventChannels {
		close(ch)
		d.unsubscribe(key, ch)
	}
}

func (d *eventDispatcher) Subscribe(e string, ch chan<- interface{}) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if d.eventChannels == nil {
		d.eventChannels = make(map[string]chan<- interface{})
	}

	d.eventChannels[e] = ch
}

func (d *eventDispatcher) Unsubscribe(e string, ch chan<- interface{}) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.unsubscribe(e, ch)
}

func (d *eventDispatcher) unsubscribe(e string, ch chan<- interface{}) {
	if _, ok := d.eventChannels[e]; !ok {
		return
	}
	delete(d.eventChannels, e)
}
