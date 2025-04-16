package main

import (
	m "machine"
	"time"

	"tinygo.org/x/drivers/hd44780"
)

func main() {
	m.Serial.Configure(m.UARTConfig{BaudRate: 9600})
	lcd, err := hd44780.NewGPIO4Bit([]m.Pin{m.D12, m.D11, m.D10, m.D9}, m.D8, m.D7, m.NoPin)
	if err != nil {
		println("error: create LCD", err.Error())
		return
	}
	if err := lcd.Configure(hd44780.Config{Width: 16, Height: 2}); err != nil {
		println("error: configure LCD", err.Error())
		return
	}

	for {
		println("Start")
		lcd.Write([]byte("Hello World!"))
		err = lcd.Display()
		if err != nil {
			println("error: display", err.Error())
		}
		time.Sleep(1 * time.Second)
	}
}
