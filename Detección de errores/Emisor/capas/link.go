package capas

import (
	"emisor/crc"
	"fmt"
)

// Capa de enlace: Codificar mensaje
// Codifica el mensaje con el algoritmo CRC-32
func EncodeMessage(message string) string {
	checksum := crc.Crc32Encode(message)

	return message + fmt.Sprintf("%032b", checksum)
}
