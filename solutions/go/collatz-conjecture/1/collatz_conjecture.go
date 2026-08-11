package collatzconjecture

import "fmt"

func CollatzConjecture(n int) (int, error) {
	etapas := 0

	if n == 0 || n < 0 {
		return 0, fmt.Errorf("Erro")
	}
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = (n * 3) + 1
		}
		etapas++
	}
	return etapas, nil
}
