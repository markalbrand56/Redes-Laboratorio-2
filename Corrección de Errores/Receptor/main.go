package main

import (
	"bufio"
	"fmt"
	"net"
)

func startServer(port string) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
		return
	}
	defer ln.Close()
	fmt.Println("Servidor escuchando en el puerto", port)

	for {
		fmt.Println("Esperando conexión...")
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error al aceptar la conexión:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func binaryToASCII(binary string) string {
	ascii := ""
	for i := 0; i < len(binary); i += 8 {
		// Convertir el byte actual a un entero
		byteInt := 0
		for j := 0; j < 8; j++ {
			bit := binary[i+j]
			if bit == '1' {
				byteInt += 1 << uint(7-j)
			}
		}
		// Convertir el entero a un caracter ASCII
		ascii += string(byteInt)
	}
	return ascii
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error al leer el mensaje:", err)
		return
	}
	message = message[:len(message)-1]
	fmt.Printf("-- Capa de Transmisión --\n")
	fmt.Printf("Mensaje recibido: %s\n", message)

	fmt.Printf("\n-- Capa de Enlace --\n")
	decodedMessage := HammingDecode(message)

	fmt.Printf("-- Capa de presentación --\n")
	fmt.Printf("Mensaje decodificado: %s\n", decodedMessage)

	fmt.Printf("\n-- Capa de Aplicación --\n")
	decodedMessage = binaryToASCII(decodedMessage)
	fmt.Printf("Mensaje decodificado en ASCII: %s\n", decodedMessage)
}

func main() {

	fmt.Println("Ingrese el puerto para el servidor:")
	var port string
	fmt.Scanln(&port)

	// Iniciar el servidor
	go startServer(port)

	// Mantener el programa principal en ejecución
	fmt.Println("Presione 'Enter' para terminar...")
	fmt.Scanln()

	/* fmt.Println("Ingrese un mensaje en binario:")
	var message string
	fmt.Scanln(&message)

	res := HammingDecode(message)

	fmt.Println("Resultado: ", res) */
}
