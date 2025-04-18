package main

import (
	"arduino-golang-daily-quote/hd44780"
	m "machine"
)

func main() {
	m.Serial.Configure(m.UARTConfig{BaudRate: 9600})
	lcd, err := hd44780.NewGPIO8Bit([]m.Pin{m.D0, m.D1, m.D2, m.D3, m.D4, m.D5, m.D6, m.D7}, m.D13, m.D12, m.NoPin)
	if err != nil {
		println("error: create LCD", err.Error())
		return
	}
	if err := lcd.Configure(hd44780.Config{Width: 16, Height: 2}); err != nil {
		println("error: configure LCD", err.Error())
		return
	}

	lcd.Test()
}
