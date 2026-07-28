package hamming

import (
	"errors"
	"strings"
)

func Distance(a, b string) (int, error) {
	letrasA := strings.Split(a, "")
	letrasB := strings.Split(b, "")
	ponto := 0
	if len(letrasA) != len(letrasB) {
		return 0, errors.New("Quantidade diferente")
	}
	for i := 0; i < len(letrasA); i++ {
		if letrasA[i] != letrasB[i] {
			ponto = ponto + 1
		}
	}
	return ponto, nil
}