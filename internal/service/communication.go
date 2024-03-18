package service

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func StartCommunication(conn serialService) error {

	testMessages := [3]string{
		strconv.Itoa(time.Now().Year()) + fmt.Sprintf("%02d", time.Now().Month()) + fmt.Sprintf("%02d", time.Now().Day()) + fmt.Sprintf("%02d", time.Now().Hour()) + fmt.Sprintf("%02d", time.Now().Minute()) + "270Clear10",
		"20000850",
		"005845670 25",
	}

	for {
		input, err := serialService.ReadFromSeial(conn)
		if err != nil {
			return err
		}
		log.Println("Input: " + string(input[0]))
		switch in := string(input[0]); in {
		case "L":
			output, err := conn.WriteToSeial([]byte("rdy"))
			if err != nil {
				return err
			}
			log.Println("Output: " + output)
		case "0":
			output, err := conn.WriteToSeial([]byte(testMessages[0]))
			if err != nil {
				return err
			}
			log.Println("Output: " + output)
		case "1":
			output, err := conn.WriteToSeial([]byte(testMessages[1]))
			if err != nil {
				return err
			}
			log.Println("Output: " + output)
		case "2":
			output, err := conn.WriteToSeial([]byte(testMessages[2]))
			if err != nil {
				return err
			}
			log.Println("Output: " + output)
		default:
			err = conn.Flush()
			if err != nil {
				return err
			}
		}
	}
}
