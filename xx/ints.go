package xx

func IntListReverse(i []int) []int {
	length := len(i)
	i2 := make([]int, length)
	for index, value := range i {
		i2[length-index-1] = value
	}
	return i2
}
