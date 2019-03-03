package main

import (
	"fmt"
	"strconv"
	s "strings"
	"unicode/utf8"
)

var p = fmt.Println

func main() {
	// test()
	// stringformater()
	// rangeString()
	strConv()
}

func test() {
	p("Contains :", s.Contains("test", "es"))
	p("Count:", s.Count("test now tt", "t"))
	p("hasPrefix:", s.HasPrefix("test", "te"))
	p("HasSuffix:", s.HasSuffix("test", "st"))
	p("Index:", s.Index("test", "e"))
	p("join:", s.Join([]string{"a", "b"}, "-"))
	p("repeat:", s.Repeat("hi", 4))
	p("replace:", s.Replace("hii", "i", "ello", 1))
	p("replace:", s.Replace("hii", "i", "ello", -1))
	p("split:", s.Split("h-e-l-l-o", "-"))
	p("toLower:", s.ToLower("HI"))
	p("toUpper:", s.ToUpper("hi"))
}

func rangeString() {
	// gives only the bytes:
	var s = "嘻哈china"
	fmt.Println("bytes")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	// 一个字符串的字节数：len(str)
	fmt.Printf("\n一个字符串的字节数:%d", len(s))
	// gives the Unicode characters:
	fmt.Println("\nthe Unicode characters")
	for codepoint, runeValue := range s {
		fmt.Printf("%d %d ", codepoint, int32(runeValue))
	}
	// 如何获取一个字符串的字符数：
	fmt.Printf("\n一个字符串的字符数:%d", utf8.RuneCountInString(s))

}

// Format 系列函数把其他类型的转换为字符串
func strConv() {
	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e)
}
