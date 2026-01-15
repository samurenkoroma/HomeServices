package main

import (
	"fmt"
	"github.com/simonvetter/modbus"
	"log"
	"time"
)

func main() {
	client, err := modbus.NewClient(&modbus.ClientConfiguration{
		URL:     "tcp://localhost:5020",
		Timeout: 2 * time.Second,
	})

	if err != nil {
		log.Fatal(err)
	}

	err = client.Open()
	if err != nil {
		log.Fatal(err)
	}

	var reg16 []uint16
	reg16, err = client.ReadRegisters(0, 4, modbus.INPUT_REGISTER)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("value: %v", reg16)
	}

	//client.WriteCoil(1, !reg16)

	// close the TCP connection/serial port
	client.Close()
}
