package main

import (
	"fmt"
	"strconv"
)

func HammingDecode(message string) string {
	// Pasar el mensaje original a un array de enteros
	messageBits := make([]int, len(message))
	for i, char := range message {
		messageBits[i], _ = strconv.Atoi(string(char))
	}

	// Para trabajar con el mensaje, se invierte el orden de los bits
	for i, j := 0, len(messageBits)-1; i < j; i, j = i+1, j-1 {
		messageBits[i], messageBits[j] = messageBits[j], messageBits[i]
	}
	// fmt.Println("Reversed message: ", messageBits)

	// Calcular cuantos bits de paridad hay en el mensaje
	parityBits := 0
	for i := 0; i < len(messageBits); i++ {
		if i&(i-1) == 0 {
			parityBits++
		}
	}
	fmt.Printf("\tMessage '%s' has %d parity bits\n", message, parityBits)

	// Verificar si hay errores en el mensaje en base a los bits de paridad
	errorPosition := 0
	for i := 0; i < parityBits; i++ {
		position := 1 << uint(i) // 2^i
		parity := 0

		// Calcular el bit de paridad para la posición actual
		for j := position - 1; j < len(messageBits); j += 2 * position {
			for k := j; k < j+position && k < len(messageBits); k++ {
				parity ^= messageBits[k]
			}
		}

		// Si el bit de paridad no coincide con el calculado, hay un error
		if parity != 0 {
			errorPosition += position
		}
	}

	if errorPosition != 0 {
		fmt.Println("\tError in position: ", errorPosition)

		if errorPosition-1 < len(messageBits) {
			messageBits[errorPosition-1] ^= 1
		} else {
			fmt.Println("\tError in parity bits, can't correct")
		}
	} else {
		fmt.Println("\tNo errors found")
	}

	// Eliminar los bits de paridad del mensaje, para regresar el mensaje original
	decodedMessage := make([]int, 0)
	for i := 0; i < len(messageBits); i++ {
		if i&(i+1) != 0 {
			decodedMessage = append(decodedMessage, messageBits[i])
		}
	}

	// Regresar el mensaje a su orden original
	for i, j := 0, len(decodedMessage)-1; i < j; i, j = i+1, j-1 {
		decodedMessage[i], decodedMessage[j] = decodedMessage[j], decodedMessage[i]
	}

	// Convertir el mensaje a string
	decoded := ""
	for _, bit := range decodedMessage {
		decoded += strconv.Itoa(bit)
	}

	return decoded
}
