package darts

func Score(x, y float64) int {
	raio := x*x + y*y
	switch {
	case raio <= 1:
		return 10
	case raio <= 25:
		return 5
	case raio > 100:
		return 0

	default:
		return 1
	}
}
