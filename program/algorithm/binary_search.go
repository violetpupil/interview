package algorithm

// 二分查找
// 时间复杂度：O(log n)
func BinarySearch(list []int, target int) int {
	i, j := 0, len(list)-1
	for i <= j {
		m := (i + j) / 2
		if list[m] < target {
			i = m + 1
		} else if list[m] > target {
			j = m - 1
		} else {
			return m
		}
	}

	return -1
}
