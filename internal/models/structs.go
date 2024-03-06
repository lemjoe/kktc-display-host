package models

import "github.com/tarm/serial"

type ConfigApp struct {
	PortName string
	BaudRate int
}

type SerialConnection struct {
	Serial *serial.Port
}
