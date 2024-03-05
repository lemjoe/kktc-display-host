package internal

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/lemjoe/kktc-display-host/internal/config"
	"github.com/lemjoe/kktc-display-host/internal/service"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() error {
	confApp, err := config.InitConfig("../.env")
	log.Println(confApp.BaudRate, confApp.PortName)
	if err != nil {
		return err
	}
	conn := service.OpenConnection(confApp)
	defer service.CloseConnection(conn)
	input, err := service.ReadFromSeial(conn)
	if err != nil {
		return err
	}
	log.Println(string(input[0]))
	for {
		input, err = service.ReadFromSeial(conn)
		if err != nil {
			return err
		}
		log.Println("Input: " + string(input[0]))
		switch in := string(input[0]); in {
		case "L":
			err = service.WriteToSeial(conn, []byte("rdy"))
		case "0":
			msg := strconv.Itoa(time.Now().Year()) + fmt.Sprintf("%02d", time.Now().Month()) + fmt.Sprintf("%02d", time.Now().Day()) + fmt.Sprintf("%02d", time.Now().Hour()) + fmt.Sprintf("%02d", time.Now().Minute()) + "270Clear10"
			err = service.WriteToSeial(conn, []byte(msg))
			log.Println("Output: " + msg)
		case "1":
			err = service.WriteToSeial(conn, []byte("20000850"))
			log.Println("Output: 20000850")
		case "2":
			err = service.WriteToSeial(conn, []byte("005845670 25"))
			log.Println("Output: 584567025")
		default:
			err = conn.Flush()
		}
		if err != nil {
			return err
		}
	}
}
