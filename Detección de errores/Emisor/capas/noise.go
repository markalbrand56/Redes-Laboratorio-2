package capas

import "math/rand/v2"

// función que reciba un mensaje (string) comprendido de solo 0s y 1s y devuelva el mensaje con ruido en base a una probabilidad p

const p = 0.01 // 1/100 de probabilidad de que un bit cambie

func AddNoise(message string) string {
	noisyMessage := ""
	for _, bit := range message {
		if rand.Float64() < p {
			if bit == '0' {
				noisyMessage += "1"
			} else {
				noisyMessage += "0"
			}
		} else {
			noisyMessage += string(bit)
		}
	}
	return noisyMessage
}
