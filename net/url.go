package main

import (
	"fmt"
	"net"
	"net/url"
	"strings"
)

func main() {
	urlExample()
	parseIP()
}
func urlExample() {
	// 我们将解析这个URL，它包含了模式，验证信息，
	// 主机，端口，路径，查询参数和查询片段
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	fmt.Println(u.Host)
	h := strings.Split(u.Host, ":")
	fmt.Println(h[0])
	fmt.Println(h[1])

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// 为了得到`k=v`格式的查询参数，使用RawQuery。你可以将
	// 查询参数解析到一个map里面。这个map为字符串作为key，
	// 字符串切片作为value。
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}

func parseIP() {
	name := "192.0.2.1"
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}
}
