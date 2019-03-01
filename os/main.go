package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// f := createFile("test.txt")
	// defer closeFile(f)
	// writeFile(f)
	// fmt.Println(f.Name())
	// env()
	// execExample()
	// readInput()
	// bufioRead()
	// readFile("src/riverside/os/test.txt")
	copyFile("src/riverside/os/test_copy.txt", "src/riverside/os/test.txt")
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}
func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}
func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
func env() {
	// 为了设置一个key/value对，使用`os.Setenv`
	// 为了获取一个key的value，使用`os.Getenv`
	// 如果所提供的key在环境变量中没有对应的value，
	// 那么返回空字符串
	os.Setenv("foo", "1")
	fmt.Println("foo:", os.Getenv("foo"))
	fmt.Println("foo 1:", os.Getenv("foo1"))
	for i, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(e)
		fmt.Println(pair)
		fmt.Println(i)
	}
}

func execExample() {
	// `exec.Command` 函数创建了一个代表外部进程的对象
	dateCmd := exec.Command("date")
	// `Output`是另一个运行命令时用来处理信息的函数，这个
	// 函数等待命令结束，然后收集命令输出。
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dateOut))
	// stdin输入数据的命令，
	// 数据输入传给外部进程的stdin，然后从它输出到stdout
	// 的运行结果收集信息
	grepCmd := exec.Command("grep", "hello")
	// 显式地获取input/output管道，启动进程，
	// 向进程写入数据，然后读取输出结果，最后等待进程结束
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()
	fmt.Println(string(grepBytes))
}

func readInput() {
	var firstName, lastName, s string
	var i int
	var f float32
	var input = "56.12 / 5212 / Go"
	var format = "%f / %d / %s"
	fmt.Println("please enter your name:")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("hi %s %s!\n", firstName, lastName)
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)
}

// 使用 bufio 包提供的缓冲读取（buffered reader）来读取数据
func bufioRead() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please input some thing:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("wrong")
		return
	}
	fmt.Printf("the input was: %s\n", input)
}

func readFile(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("wrong")
		return
	}
	defer f.Close()
	inputReader := bufio.NewReader(f)
	for {
		// s, err := inputReader.ReadString('\n')
		s, _, err := inputReader.ReadLine()
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("wrong")
			return
		}
		// fmt.Println(s)
		fmt.Println(string(s))
	}
	// 整个文件的内容读到一个字符串里：可以使用 io/ioutil 包里的 ioutil.ReadFile() 方法
}

func copyFile(dstFile, srcFile string) (written int64, err error) {
	src, err := os.Open(srcFile)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.Create(dstFile)
	if err != nil {
		fmt.Println("create file wrong")
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
