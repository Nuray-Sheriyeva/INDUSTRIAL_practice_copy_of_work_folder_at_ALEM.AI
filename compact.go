package piscine

func Compact(ptr *[]string) int {
	var array []string
	for _, v := range *ptr {
		if v != "" {
			array = append(array, v)
		}
	}
	*ptr = array
	return len(array)
}
