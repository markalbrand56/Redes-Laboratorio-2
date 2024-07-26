package main

import (
	"emisor/capas"
	"fmt"
)

func main() {
	var message string
	fmt.Println("Ingrese un mensaje en binario:")
	fmt.Scanln(&message)

	encodedMessage := capas.EncodeMessage(message)

	fmt.Printf("Mensaje codificado con CRC-32: %s\n", encodedMessage)

	server := capas.NewServer("localhost", 8080)

	server.Send(encodedMessage)
}
