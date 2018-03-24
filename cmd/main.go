package main

import (
	"fmt"
	"sync"

	"github.com/yanndr/alert"
	"github.com/yanndr/temperature"
)

var wg sync.WaitGroup
var num = 0

func main() {

	m := alert.NewMonitor(&wg)

	a := alert.NewRaiseAlert(35, 1, raiseHandler)
	m.Subscribe("Raised to 35 alert", a.Channel())

	tt := temperature.NewWithHandler(0, temperature.Celsius, func(t temperature.Temperature) {
		wg.Add(1)
		fmt.Println("temp received: ", t)
		m.MonitorChan <- t
	})
	tt.SetTemperature(temperature.New(32, temperature.Celsius))

	tt.SetTemperature(temperature.New(35, temperature.Celsius))
	//time.Sleep(time.Second / 2)
	tt.SetTemperature(temperature.New(37, temperature.Celsius))
	//time.Sleep(time.Second / 2)
	tt.SetTemperature(temperature.New(35, temperature.Celsius))
	//	time.Sleep(time.Second / 2)
	tt.SetTemperature(temperature.New(34.2, temperature.Celsius))
	//	time.Sleep(time.Second / 2)
	tt.SetTemperature(temperature.New(35, temperature.Celsius))
	//	time.Sleep(time.Second / 2)
	tt.SetTemperature(temperature.New(33, temperature.Celsius))
	//	time.Sleep(time.Second / 2)
	tt.SetTemperature(temperature.New(35, temperature.Celsius))
	//time.Sleep(time.Second / 2)
	tt.SetTemperature(temperature.New(35, temperature.Celsius))

	defer m.Stop()
	fmt.Println("end ")
	wg.Wait()
}

func raiseHandler(t interface{}) {
	temp, ok := t.(temperature.Temperature)
	if !ok {
		fmt.Println("Not a temperature")
		return
	}
	fmt.Printf("------ Temperature reached %s ------- \n", temp)

}
