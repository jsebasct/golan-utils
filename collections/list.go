package collections

func SumItems(numbers []int) int {
	res := 0
	for _, value := range numbers {
		res += value
	}
	return res
}
