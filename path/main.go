package main

import (
	"fmt"
	"path"
	// "path/filepath"
)

func main() {
	pathExample()
}

func pathExample() {
	// 返回路径的最后一部分
	fmt.Println(path.Base("/var/www"))
	fmt.Println(path.Clean("/a/b/../././c"))
	fmt.Println(path.Dir("/a/b/../c/d/e"))
	fmt.Println(path.Ext("/a/b/test.txt"))
	fmt.Println(path.IsAbs("/a/b/c"))
	fmt.Println(path.Join("a/b", "c"))
	fmt.Println(path.Split("/a/b/test.txt"))
}
func filepathExample() {
	// fmt.Println(filepath.Walk(".",))
}
