package capas

import (
	"emisor/crc"
	"fmt"
)

func EncodeMessage(message string) string {
	checksum := crc.Crc32Encode(message)
	return EncodeString(message) + fmt.Sprintf("%032b", checksum)
}
