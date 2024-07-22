package main

import (
	"fmt"
)

const poly uint32 = 0xEDB88320

// Tabla CRC-32
func MakeCrcTable() [256]uint32 {
	var table [256]uint32
	for i := 0; i < 256; i++ {
		crc := uint32(i)
		for j := 0; j < 8; j++ {
			if crc&1 != 0 {
				crc = (crc >> 1) ^ poly
			} else {
				crc >>= 1
			}
		}
		table[i] = crc
	}
	return table
}

// CRC-32 de una cadena de bits
func Crc32Encode(data string) uint32 {
	table := MakeCrcTable()
	crc := uint32(0xFFFFFFFF)
	for i := 0; i < len(data); i++ {
		byteVal := data[i]
		crc = table[(crc^uint32(byteVal))&0xFF] ^ (crc >> 8)
	}
	return crc ^ 0xFFFFFFFF
}

func main() {
	var message string
	fmt.Println("Ingrese un mensaje en binario:")
	fmt.Scanln(&message)

	crc := Crc32Encode(message)
	crcChecksum := fmt.Sprintf("%032b", crc)
	encodedMessage := message + crcChecksum
	fmt.Printf("Mensaje codificado con CRC-32: %s\n", encodedMessage)
}
