package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
)

func main() {
	//lengthOfLongestSubstring("wudfdif38dfo")
	// maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	// romanToInt("MCMXCIV")
	// romanToIntByte("MCMXCIV")
	//fmt.Println(ThreeSum([]int{-1, -2, -3, 4, 1, 3, 0, 3, -2, 1, -2, 2, -1, 1, -5, 4, -3}))
	//threeSumClosest([]int{-1,2,1,-4},1)
	//threeSumClosest2([]int{-1,2,1,-4},1)
	//letterCombinations("235")

	listNode1 := ListNode{1, nil}

	//fmt.Println(listNode1,listNode2)
	//listNode3:=ListNode{1,2}
	//listNode4:=ListNode{1,2}
	//listNode1=ListNode{1,&ListNode2{}}
	//listNode2:=ListNode{1,&ListNode2{}}
	//listNode3:=ListNode{1,2}
	//listNode4:=ListNode{1,2}
	removeNthFromEnd(&listNode1, 2)
}

func lengthOfLongestSubstring(s string) int {
	lastOccurred := make(map[rune]int) //记录字符最后出现的位置
	// 最大子字符串开始的索引
	maxStart := 0
	// 比较的子字符串的开始的索引
	start := 0
	//  最大子字符串的长度
	maxLength := 0

	for i, ch := range []rune(s) {
		//字符非第一次出现，开始新的子字符串
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		// 子字符串的长度对比
		if i-start+1 > maxLength {
			maxStart = start
			maxLength = i - start + 1
		}
		// 记录最新的索引
		lastOccurred[ch] = i
	}
	fmt.Println(s[maxStart : maxStart+maxLength])

	// fmt.Printf("Longest Substring Without Repeating Characters:%d", maxLength)
	return maxLength
}

func reverse(x int) int {
	reversed := 0
	for {
		r := x % 10
		reversed += r

		x /= 10
		if x == 0 {
			break
		}
		reversed *= 10
	}

	if reversed > int(math.Exp2(31.0)) || reversed < int(-math.Exp2(31.0)) {
		return 0
	}

	return reversed
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	y := x
	reversedNumber := 0
	for y != 0 {
		rem := y % 10
		reversedNumber = reversedNumber*10 + rem
		y = y / 10
	}

	if reversedNumber == x {
		return true
	}

	return false
}

func maxArea(height []int) int {
	l := 0
	lMax := 0
	r := len(height) - 1
	rMax := len(height) - 1
	mArea := (len(height) - 1) * min(height[l], height[r])

	for i := 0; i < len(height); i++ {
		if l < r {
			if height[l] < height[r] {
				l = l + 1
			} else {
				r = r - 1
			}
			if mArea < (r-l)*min(height[l], height[r]) {
				mArea = (r - l) * min(height[l], height[r])
				lMax = l
				rMax = r
			}
		} else {
			fmt.Println("进行次数：", i)
			break
		}
	}
	fmt.Println("左边的序号：", lMax, "；左边的值：", height[lMax])
	fmt.Println("右边的序号：", rMax, "；右边的值：", height[rMax])
	fmt.Println("最大面积：", mArea)

	return mArea
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func intToRoman(num int) string {
	var roman string
	var count int
	if num < 1 || num > 3999 {
		return roman
	}
	romans := [][]string{
		{"I", "V", "X"},
		{"X", "L", "C"},
		{"C", "D", "M"},
		{"M", "M", "M"},
	}

	for num > 0 {
		mod := num % 10
		num /= 10

		roman = helper(mod, romans[count]) + roman
		count++
	}

	return roman
}
func helper(x int, units []string) (roman string) {
	for x > 0 {
		switch {
		case x == 9:
			roman += units[0] + units[2]
			x -= 9
		case x >= 5:
			roman += units[1]
			x -= 5
		case x == 4:
			roman += units[0] + units[1]
			x -= 4
		default:
			roman += units[0]
			x--
		}
	}
	return roman
}
func romanToInt(s string) int {
	var num int
	if len(s) == 0 {
		return num
	}
	romans := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

	for i := 0; i < len(s); {
		if i < len(s)-1 && romans[s[i:i+1]] < romans[s[i+1:i+2]] {
			num += romans[s[i+1:i+2]] - romans[s[i:i+1]]
			i += 2
		} else {
			num += romans[s[i:i+1]]
			i++
		}
	}
	return num
}
func romanToIntByte(s string) int {
	var num int
	if len(s) == 0 {
		return num
	}
	romans := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	sByte := []byte(s)
	for i := 0; i < len(sByte); {
		if i < len(sByte)-1 && romans[sByte[i]] < romans[sByte[i+1]] {
			num += romans[sByte[i+1]] - romans[sByte[i]]
			i += 2
		} else {
			num += romans[sByte[i]]
			i++
		}
	}
	return num
}
func longestCommonPrefix(strs []string) string {
	var prefix string
	if len(strs) == 0 {
		return prefix
	}

	prefix = strs[0]
	for i := 1; i < len(strs); i++ {
		if len(prefix) == 0 {
			break
		}
		for strings.HasPrefix(strs[i], prefix) == false {
			prefix = string([]rune(prefix)[:len(prefix)-1])
		}
	}

	return prefix
}
func ThreeSum(nums []int) [][]int {
	results := [][]int{}
	n := len(nums)
	if n < 3 {
		return results
	}
	sort.Ints(nums)
	if nums[0] > 0 || nums[n-1] < 0 {
		return results
	}
	var wg sync.WaitGroup
	var wg2 sync.WaitGroup
	resultSum := make(chan []int, 10)
	for i := 0; i < n-2; i++ {
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		wg.Add(1)
		go threeSumZero(&wg, nums, i, n, resultSum)
	}
	wg2.Add(1)
	go rec(&wg2, resultSum, &results)
	// 待计数值变成零，Wait() 才会返回
	wg.Wait()
	// 关闭通道
	close(resultSum)
	wg2.Wait()
	return results
}
func rec(wg *sync.WaitGroup, resultSum chan []int, results *[][]int) {
	defer wg.Done()
	for v := range resultSum {
		*results = append(*results, v)
	}
}
func threeSumZero(wg *sync.WaitGroup, nums []int, i, n int, in chan []int) {
	defer wg.Done()
	target := -nums[i]
	left := i + 1
	right := n - 1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			in <- []int{nums[left], nums[right], nums[i]}
			left++
			right--
			// 去重
			for left < right && nums[left] == nums[left-1] {
				left++
			}
			for left < right && nums[right] == nums[right+1] {
				right--
			}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
}

func threeSum(nums []int) [][]int {
	results := [][]int{}
	n := len(nums)
	if n < 3 {
		return results
	}
	sort.Ints(nums)
	if nums[0] > 0 || nums[n-1] < 0 {
		return results
	}

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		left := i + 1
		right := n - 1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				results = append(results, []int{nums[left], nums[right], nums[i]})
				left++
				right--
				// 去重
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum > target {
				right--
			} else {
				left++
			}
		}

	}
	return results
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	length := len(nums)
	bestTarget := nums[0] + nums[1] + nums[2] // initial value
	if length == 3 {
		return bestTarget
	}
	for i := 0; i < length-2; i++ {
		newTarget := target - nums[i]
		left, right := i+1, length-1
		for left < right {
			if abs(bestTarget-target) > abs(nums[left]+nums[right]+nums[i]-target) {
				bestTarget = nums[left] + nums[right] + nums[i]
			}
			if bestTarget == target {
				return target
			}
			if nums[left]+nums[right] < newTarget {
				left++
			} else {
				right--
			}
		}
	}
	return bestTarget
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	digitWordsMap := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	for _, digit := range digits {
		words := digitWordsMap[string(digit)]
		tmp := make([]string, 0)

		for _, word := range words {
			if len(res) > 0 {
				//生成新的slice
				for _, item := range res {
					tmp = append(tmp, item+string(word))
				}
			} else {
				//第一个字母
				tmp = append(tmp, string(word))
			}
		}
		//赋值
		res = tmp
	}
	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	hash := make(map[int]*ListNode)

	for i := 0; head != nil; i++ {
		hash[i] = head
		head = head.Next
	}

	switch {
	case len(hash)-n > 0:
		head = hash[0]
		hash[len(hash)-n-1].Next = hash[len(hash)-n].Next
	case len(hash) > 1:
		head = hash[1]
	default:
		head = nil
	}

	return head
}
