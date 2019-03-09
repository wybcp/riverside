package bubble

// 冒泡算法，flag 可以在没有数据交换的时候，跳出循环

// BubbleIncSort 冒泡算法，从小到大
func BubbleIncSort(values []int) {
	for i := 0; i < len(values)-1; i++ {
		flag := false
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

// BubbleDecSort 冒泡算法，从大到小
func BubbleDecSort(values []int) {
	for i := 0; i < len(values)-1; i++ {
		flag := false
		for j := i + 1; j < len(values); j++ {
			if values[i] < values[j] {
				values[i], values[j] = values[j], values[i]
				flag = true
			}
		}
		if !flag {
			break
		}
	}

}
