package crc

const poly uint32 = 0xEDB88320 // Polinomio CRC-32

// Referencia: A simple example: CRC32 calculation – The LXP32 Processor. (s. f.). https://lxp32.github.io/docs/a-simple-example-crc32-calculation/

// Tabla CRC-32 para cálculos rápidos
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
