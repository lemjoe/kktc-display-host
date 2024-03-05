package service

import (
	"log"

	"github.com/lemjoe/kktc-display-host/internal/models"
	"github.com/tarm/serial"
)

func OpenConnection(config models.ConfigApp) *serial.Port {
	c := &serial.Config{Name: config.PortName, Baud: config.BaudRate}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	return s
}

func CloseConnection(port *serial.Port) {
	err := port.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func WriteToSeial(port *serial.Port, msg []byte) error {
	_, err := port.Write(msg)
	if err != nil {
		return err
	}
	return nil
}

func ReadFromSeial(port *serial.Port) ([]byte, error) {
	buf := make([]byte, 128)
	_, err := port.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func Flush(port *serial.Port) error {
	err := port.Flush()
	if err != nil {
		return err
	}
	return nil
}
