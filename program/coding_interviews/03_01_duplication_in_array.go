package program

// 找出数组中重复的数字
func Duplicate(numbers []int, duplication *int) bool {
	if len(numbers) == 0 {
		return false
	}

	for _, n := range numbers {
		if n < 0 || n > len(numbers)-1 {
			return false
		}
	}

	for i := 0; i < len(numbers); i++ {
		for numbers[i] != i {
			if numbers[i] == numbers[numbers[i]] {
				*duplication = numbers[i]
				return true
			}

			numbers[i], numbers[numbers[i]] = numbers[numbers[i]], numbers[i]
		}
	}

	return false
}
