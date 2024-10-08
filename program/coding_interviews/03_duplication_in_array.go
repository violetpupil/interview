package program

// 找出数组中重复的数字
// 在一个长度为n的数组里的所有数字都在0~n-1的范围内。
// 数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。
// 请找出数组中任意一个重复的数字。例如，如果输入长度为7的数组{2,3,1,0,2,5,3}，那么对应的输出是重复的数字2或者3。
func FindDuplicate(list []int) int {
	if len(list) == 0 {
		return -1
	}

	// 元素不合法
	for _, v := range list {
		if v < 0 || v >= len(list) {
			return -2
		}
	}

	for i := 0; i < len(list); i++ {
		// 让每个值都在对应的索引上
		for list[i] != i {
			if list[i] == list[list[i]] {
				return list[i]
			} else {
				list[i], list[list[i]] = list[list[i]], list[i]
			}
		}
	}

	return -1
}

// 不修改数组找出重复的数字
// 在一个长度为n+1的数组里的所有数字都在1~n的范围内，所以数组中至少有一个数字是重复的。
// 请找出数组中任意一个重复的数字，但不能修改输入的数组。
// 例如，如果输入长度为8的数组{2,3,5,4,3,2,6,7}，那么对应的输出是重复的数字2或者3。

// 空间效率优先
// 时间复杂度 Ο(nlogn)
// 空间复杂度 O(1)
func FindDuplicateSpace(list []int) int {
	// 数组不合法
	if len(list) < 2 {
		return -1
	}
	// 元素不合法
	for _, v := range list {
		if v < 1 || v >= len(list) {
			return -1
		}
	}

	// 二分法
	start, end := 1, len(list)-1
	for start < end {
		middle := (start + end) / 2
		// 计算start,middle间的元素个数
		// 如果元素个数大于数字个数，则重复的在左半边
		// 否则在右半边
		count := FindDuplicateCount(list, start, middle)
		if count > middle-start+1 {
			end = middle
		} else {
			start = middle + 1
		}
	}

	// 此时start=end
	// 上面的合法性判断，保证了至少一个数字重复
	// 所以最后留下的就是重复的
	return start
}

func FindDuplicateCount(list []int, min, max int) int {
	var c int
	for _, v := range list {
		if v >= min && v <= max {
			c++
		}
	}
	return c
}
