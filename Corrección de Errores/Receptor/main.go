package main

import "fmt"

func main() {

	fmt.Println("Ingrese un mensaje en binario:")
	var message string
	fmt.Scanln(&message)

	res := HammingDecode(message)

	fmt.Println("Resultado: ", res)
}
