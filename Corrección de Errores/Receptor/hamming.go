package main

import (
	"fmt"
	"strconv"
)

func HammingDecode(message string) string {
	fmt.Println("Decoding message with Hamming code")

	// Pasar el mensaje original a un array de enteros
	messageBits := make([]int, len(message))
	for i, char := range message {
		messageBits[i], _ = strconv.Atoi(string(char))
	}
	fmt.Println("Message: ", messageBits)

	// Para trabajar con el mensaje, se invierte el orden de los bits
	for i, j := 0, len(messageBits)-1; i < j; i, j = i+1, j-1 {
		messageBits[i], messageBits[j] = messageBits[j], messageBits[i]
	}
	fmt.Println("Reversed message: ", messageBits)

	// Calcular cuantos bits de paridad hay en el mensaje
	parityBits := 0
	for i := 0; i < len(messageBits); i++ {
		if i&(i-1) == 0 {
			parityBits++
		}
	}
	fmt.Println("Parity bits: ", parityBits)

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
		fmt.Println("Error in position: ", errorPosition)
		messageBits[errorPosition-1] ^= 1
	} else {
		fmt.Println("No errors found")
	}

	// Se invierte nuevamente el mensaje para devolverlo en el orden original
	for i, j := 0, len(messageBits)-1; i < j; i, j = i+1, j-1 {
		messageBits[i], messageBits[j] = messageBits[j], messageBits[i]
	}
	fmt.Println("Corrected message: ", messageBits)

	// Se eliminan los bits de paridad y devolver en string el mensaje original
	decodedMessage := ""
	for i := 0; i < len(messageBits); i++ {
		if i&(i-1) == 0 {
			decodedMessage += strconv.Itoa(messageBits[i])
		}
	}

	return decodedMessage
}
