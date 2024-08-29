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

	var small, equal, big []int
	for _, v := range list {
		if v < pivot {
			small = append(small, v)
		} else if v == pivot {
			equal = append(equal, v)
		} else {
			big = append(big, v)
		}
	}

	return append(append(QuickSort(small), equal...), QuickSort(big)...)
}
