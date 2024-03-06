package internal

import (
	"log"

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
	err = service.StartCommunication(conn)
	if err != nil {
		return err
	}
	return nil
}
