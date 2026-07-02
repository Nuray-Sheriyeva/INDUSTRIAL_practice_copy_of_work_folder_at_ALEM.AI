package piscine

func DescendAppendRange(max, min int) []int {
	if min >= max {
		arrr := []int{}
		return arrr
	}
	var arrr []int
	for i := max; i > min; i-- {
		arrr = append(arrr, i)
	}
	return arrr
}
