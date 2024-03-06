package service

import (
	"log"

	"github.com/lemjoe/kktc-display-host/internal/models"
	"github.com/tarm/serial"
)

type serialService struct {
	port models.SerialConnection
}

func OpenConnection(config models.ConfigApp) serialService {
	conf := &serial.Config{Name: config.PortName, Baud: config.BaudRate}
	s, err := serial.OpenPort(conf)
	if err != nil {
		log.Fatal(err)
	}
	connection := models.SerialConnection{Serial: s}
	return serialService{port: connection}
}

func CloseConnection(conn serialService) {
	err := conn.port.Serial.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func (c serialService) WriteToSeial(msg []byte) (string, error) {
	_, err := c.port.Serial.Write(msg)
	if err != nil {
		return "", err
	}
	return string(msg), nil
}

func (c serialService) ReadFromSeial() ([]byte, error) {
	buf := make([]byte, 128)
	_, err := c.port.Serial.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (c serialService) Flush() error {
	err := c.port.Serial.Flush()
	if err != nil {
		return err
	}
	return nil
}
