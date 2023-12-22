package util

func GreatestCommonDivisor(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LowestCommonMultiple(a, b int, integers ...int) int {
	result := a * b / GreatestCommonDivisor(a, b)

	for i := 0; i < len(integers); i++ {
		result = LowestCommonMultiple(result, integers[i])
	}

	return result
}
