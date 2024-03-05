package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/lemjoe/kktc-display-host/internal/models"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func CreateDefaultConfig() error {
	log.Println("Config file is missing. Creating default config")
	input, err := os.ReadFile("../.env.default")
	if err != nil {
		return err
	}

	err = os.WriteFile("../.env", input, 0644)
	if err != nil {
		return err
	}
	return nil
}

func InitConfig(confPath string) (models.ConfigApp, error) {
	if confPath != "" {
		if fileExists(confPath) { //if from dot env
			if err := godotenv.Load(confPath); err != nil {
				return models.ConfigApp{}, fmt.Errorf("InitConfig: unable to read file '%s'", confPath)
			}

		} else {
			if err := CreateDefaultConfig(); err != nil {
				return models.ConfigApp{}, fmt.Errorf("unable to create default config. You should create it manually")
			}
			godotenv.Load(confPath)
		}

	}

	// Application config
	defaultConfApp := models.ConfigApp{
		PortName: "COM3",
		BaudRate: 9600,
	}
	PORT_NAME, exist := os.LookupEnv("PORT_NAME")
	if !exist {
		fmt.Printf("warn: %s\n", fmt.Errorf("env '%s' not found", "PORT_NAME"))
	} else {
		defaultConfApp.PortName = PORT_NAME
	}
	BAUD_RATE, exist := os.LookupEnv("BAUD_RATE")
	if !exist {
		fmt.Printf("warn: %s\n", fmt.Errorf("env '%s' not found", "BAUD_RATE"))
	} else {
		defaultConfApp.BaudRate, _ = strconv.Atoi(BAUD_RATE)
	}
	return defaultConfApp, nil
}
