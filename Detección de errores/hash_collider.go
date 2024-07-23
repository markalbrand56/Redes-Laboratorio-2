package main

import (
	"fmt"
	"hash/crc32"
	"sync"
)

func AddLetter(c chan string, combo string, alphabet string, length int) {
	if length <= 0 {
		return
	}
	var newCombo string
	for _, ch := range alphabet {
		newCombo = combo + string(ch)
		c <- newCombo
		AddLetter(c, newCombo, alphabet, length-1)
	}
}

func worker(wChan chan string, target uint32, originalMessage string, wg *sync.WaitGroup) {
	defer wg.Done()
	for tString := range wChan {
		if crc32.ChecksumIEEE([]byte(tString)) == target && tString != originalMessage {
			fmt.Println("Collision found:", tString)
			close(wChan)
			return
		}
	}
}

func main() {
	// Calcular el checksum objetivo del mensaje "1100"
	originalMessage := "1100"
	targetChecksum := crc32.ChecksumIEEE([]byte(originalMessage))
	fmt.Printf("Target Checksum: %08x\n", targetChecksum)

	// Longitud máxima de la cadena a generar
	maxLen := 5
	// Número de workers concurrentes
	numWorkers := 8
	// Alfabeto a usar
	alphabet := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}"

	var wg sync.WaitGroup
	wChan := make(chan string, 1000)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(wChan, targetChecksum, originalMessage, &wg)
	}

	// Generar todas las combinaciones posibles
	go func() {
		AddLetter(wChan, "", alphabet, maxLen)
		close(wChan)
	}()

	wg.Wait()
}
