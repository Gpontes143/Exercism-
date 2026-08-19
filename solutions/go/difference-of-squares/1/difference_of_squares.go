package differenceofsquares

func SquareOfSum(n int) int {
	soma := 0
	for i := 1; i <= n; i++ {
		soma = i + soma
	}

	return soma * soma
}

func SumOfSquares(n int) int {
	quadrado, soma := 0, 0
	for i := 1; i <= n; i++ {
		quadrado = i * i
		soma = quadrado + soma
	}
	return soma
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
