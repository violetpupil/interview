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

// 实现二
// 原地(in-place)交换元素
// left 左索引 right 右索引
func QuickSortInPlace(list []int, left, right int) {
	if left < right {
		key := list[(left+right)/2]
		i := left
		j := right

		for {
			for list[i] < key {
				i++
			}
			for list[j] > key {
				j--
			}
			if i >= j {
				break
			}
			list[i], list[j] = list[j], list[i]
		}

		QuickSortInPlace(list, left, i-1)
		QuickSortInPlace(list, j+1, right)
	}
}
