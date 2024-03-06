package service

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

// type communicationService struct {
// }

func StartCommunication(conn serialService) error {
	for {
		input, err := serialService.ReadFromSeial(conn)
		if err != nil {
			return err
		}
		log.Println("Input: " + string(input[0]))
		switch in := string(input[0]); in {
		case "L":
			output, err := serialService.WriteToSeial(conn, []byte("rdy"))
			if err != nil {
				return err
			}
			log.Println("Output: " + output)
		case "0":
			msg := strconv.Itoa(time.Now().Year()) + fmt.Sprintf("%02d", time.Now().Month()) + fmt.Sprintf("%02d", time.Now().Day()) + fmt.Sprintf("%02d", time.Now().Hour()) + fmt.Sprintf("%02d", time.Now().Minute()) + "270Clear10"
			output, err := serialService.WriteToSeial(conn, []byte(msg))
			if err != nil {
				return err
			}
			log.Println("Output: " + output)
		case "1":
			output, err := serialService.WriteToSeial(conn, []byte("20000850"))
			if err != nil {
				return err
			}
			log.Println("Output: " + output)
		case "2":
			output, err := serialService.WriteToSeial(conn, []byte("005845670 25"))
			if err != nil {
				return err
			}
			log.Println("Output: " + output)
		default:
			err = serialService.Flush(conn)
			if err != nil {
				return err
			}
		}
	}
}
