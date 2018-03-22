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
		fmt.Printf("------ Temperature reached %s ------- \n", temp)
	})
	m.Subscribe("Raised to 35 alert", a.Channel())

	tt := temperature.NewTemperatureWithHandler(0, temperature.Celsius, func(t temperature.Temperature) {
		fmt.Println("new temperature received", t)
		m.MonitorChan <- t
	})
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
