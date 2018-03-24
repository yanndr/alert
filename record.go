package alert

import (
	"bufio"
	"fmt"
	"io"
)

type recorder struct {
	w io.Writer
	listener
}

func NewRecorder(w io.Writer, action handler) *recorder {
	a := &recorder{
		listener: listener{
			ch:      make(chan interface{}),
			done:    make(chan int),
			handler: action,
		},
	}
	a.start()
	return a
}

func (r *recorder) Write(bytes []byte) (int, error) {
	w := bufio.NewWriter(r.w)
	i := 0
	for _, b := range bytes {
		if err := w.WriteByte(b); err != nil {
			return 0, fmt.Errorf("error writing byte: %s", err)
		}
		i++
	}

	w.Flush()
	return i, nil
}

func (r *recorder) start() {
	go func() {
		var exit bool
		for !exit {
			select {
			case val := <-r.ch:
				v, ok := val.(float64)
				if !ok {
					valuer, ok := val.(Valuer)
					if !ok {
						continue
					}
					v = valuer.Value()
				}
				s := fmt.Sprintf("%v,", v)
				r.Write([]byte(s))

			case <-r.done:
				exit = true
			}
		}
	}()
}
