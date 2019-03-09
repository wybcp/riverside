package byte

// 求奇偶速度较快

// AndExample and位运算
func AndExample(num int) string {
	if num&1 == 1 {
		// fmt.Printf("%d is odd\n", num)
		return "odd"
	}
	return "even"

}

// DivExamle 求余运算
func DivExamle(num int) string {
	if num%2 == 1 {
		// fmt.Printf("%d is odd\n", num)
		return "odd"
	}
	return "even"
}
