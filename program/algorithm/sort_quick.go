package algorithm

// 快速排序
// 通过比较基准pivot，将数列分区partition
// 时间复杂度 Ο(nlogn)
// log n 求以2为底的对数

// 实现一
// 大量的空间占用
func QuickSort(list []int) []int {
	if len(list) <= 1 {
		return list
	}

	pivot := list[0]

	var small, big []int
	for i := 1; i < len(list); i++ {
		if list[i] < pivot {
			small = append(small, list[i])
		} else {
			big = append(big, list[i])
		}
	}

	return append(append(QuickSort(small), pivot), QuickSort(big)...)
}
