package capas

import (
	"strconv"
)

// Capa de presentación: Codificar mensaje
// Transforma caracteres a su representación en binario (ASCII)
func EncodeString(message string) string {
	binaryMessage := ""
	for _, char := range message {
		ascii := int(char)
		binary := strconv.FormatInt(int64(ascii), 2)
		binary = "00000000"[len(binary):] + binary // Rellenar con ceros a la izquierda

		binaryMessage += binary
	}

	return binaryMessage
}
