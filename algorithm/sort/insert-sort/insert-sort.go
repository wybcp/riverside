package insert

// 插入排序算法

// 	numbers := []int{6, 2, 7, 3, 8, 5}

// InsertIncSort 插入排序，从小到大
func InsertIncSort(s []int) {
	for j := 1; j < len(s); j++ {
		i := j - 1
		for i >= 0 && s[i] > s[i+1] {
			s[i+1], s[i] = s[i], s[i+1]
			i--
		}
	}
}

// InsertDecSort 插入排序，从大到小
func InsertDecSort(s []int) {
	for j := 1; j < len(s); j++ {
		i := j - 1
		for i >= 0 && s[i] < s[i+1] {
			s[i+1], s[i] = s[i], s[i+1]
			i--
		}
	}
}
