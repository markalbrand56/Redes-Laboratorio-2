package capas

import (
	"fmt"
	"math/rand/v2"
)

const p = 0.01 // 1/100 de probabilidad de que un bit cambie

// Capa de ruido: Agregar ruido al mensaje
// Cambia aleatoriamente algunos bits del mensaje, en función de la probabilidad p
func AddNoise(message string) string {
	noisyMessage := ""
	for _, bit := range message {
		if rand.Float64() < p {
			if bit == '0' {
				noisyMessage += "1"
			} else {
				noisyMessage += "0"
			}
			fmt.Println("Ruido aplicado")
		} else {
			noisyMessage += string(bit)
		}
	}
	return noisyMessage
}
