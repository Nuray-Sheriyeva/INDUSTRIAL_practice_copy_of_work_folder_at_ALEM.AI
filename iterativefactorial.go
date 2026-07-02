package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	} else {
		result := 1
		for i := 2; i < nb+1; i++ {
			result = result * i
			if result > 2432902008176640000 {
				result = 0
				break
			}
		}
		return result
	}
}
