package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		var array []int
		return array
	} else {
		size := max - min
		array := make([]int, size)
		counter := 0
		for i := min; i < max; i++ {
			array[counter] = i
			counter++
		}
		return array
	}
}
