package piscine

func AppendRange(min, max int) []int {
	if min > max {
		var array []int
		for i := 0; i < 1; i++ {
			array = append(array)
		}
		return array
	} else {
		var array []int
		for i := min; i < max; i++ {
			array = append(array, i)
		}
		return array
	}
}
