package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	ponto := 0
	if len(a) != len(b) {
		return 0, errors.New("Quantidade diferente")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			ponto = ponto + 1
		}
	}
	return ponto, nil
}
