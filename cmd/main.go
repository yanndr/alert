package main

import (
	"fmt"
	"time"

	"github.com/yanndr/alert"
	"github.com/yanndr/temperature"
)

func main() {

	m := alert.NewMonitor()
	defer m.Stop()

	a := alert.NewRaiseAlert(35, 1, func(t interface{}) {
		temp, ok := t.(temperature.Temperature)
		if !ok {
			fmt.Println("Not a temperature")
			return
		}
		fmt.Println("Temperature reached", temp)
	})
	m.Subscribe("My channel", a.Channel())

	tt := temperature.NewMonitoredTemperature(temperature.Celsius, m.MonitorChan)
	tt.SetTemperature(temperature.NewCelsius(32))

	tt.SetTemperature(temperature.NewCelsius(35))
	time.Sleep(time.Second / 2)
	tt.SetTemperature(temperature.NewCelsius(37))
	time.Sleep(time.Second / 2)
	tt.SetTemperature(temperature.NewCelsius(35))
	time.Sleep(time.Second / 2)
	tt.SetTemperature(temperature.NewCelsius(34.2))
	time.Sleep(time.Second / 2)
	tt.SetTemperature(temperature.NewCelsius(35))
	time.Sleep(time.Second / 2)
	tt.SetTemperature(temperature.NewCelsius(33))
	time.Sleep(time.Second / 2)
	tt.SetTemperature(temperature.NewCelsius(35))
	time.Sleep(time.Second / 2)

}
