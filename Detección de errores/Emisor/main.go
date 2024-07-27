package main

import (
	"emisor/capas"
	"fmt"
)

func main() {
	var message string
	fmt.Println("Ingrese un mensaje en binario:")
	fmt.Scanln(&message)

	encodedMessage := capas.EncodeString(message)        // Codificar el mensaje a su representación en (ascii binario)
	hashedMessage := capas.EncodeMessage(encodedMessage) // Codificar el mensaje con CRC-32
	hashedMessage = capas.AddNoise(hashedMessage)        // Agregar ruido al mensaje

	server := capas.NewServer("localhost", 8080)

	server.Send(hashedMessage)
}
